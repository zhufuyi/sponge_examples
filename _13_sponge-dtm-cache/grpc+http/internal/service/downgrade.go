// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/dtm-labs/rockscache"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"
	"stock/internal/dao"
	"stock/internal/ecode"
	"stock/internal/model"
	"stock/internal/rpcclient"
	"time"

	"google.golang.org/grpc"

	//"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	//"github.com/zhufuyi/sponge/pkg/logger"

	stockV1 "stock/api/stock/v1"
	//"stock/internal/cache"
	//"stock/internal/dao"
	//"stock/internal/ecode"
	//"stock/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		stockV1.RegisterDowngradeServer(server, NewDowngradeServer())
	})
}

var _ stockV1.DowngradeServer = (*downgrade)(nil)

type downgrade struct {
	stockV1.UnimplementedDowngradeServer

	db                *sql.DB
	strongCacheClient *rockscache.Client
}

// NewDowngradeServer create a server
func NewDowngradeServer() stockV1.DowngradeServer {
	return &downgrade{
		db:                model.GetSDB(),
		strongCacheClient: model.GetStrongRockscacheClient(),
	}
}

// Update 更新数据，升降级中的DB和缓存强一致性
func (s *downgrade) Update(ctx context.Context, req *stockV1.UpdateDowngradeRequest) (*stockV1.UpdateDowngradeRequestReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	gid := newGid()
	downgradeBranchURL := rpcclient.GetStockEndpoint() + stockV1.Downgrade_DowngradeBranch_FullMethodName
	downgradeBranchBody := &stockV1.DowngradeBranchRequest{
		Gid:   gid,
		Key:   getStockCacheKey(req.Id),
		Id:    req.Id,
		Stock: req.Stock,
	}
	headers := map[string]string{interceptor.ContextRequestIDKey: interceptor.ServerCtxRequestID(ctx)}

	saga := dtmgrpc.NewSagaGrpc(rpcclient.GetDtmEndpoint(), gid, dtmgrpc.WithBranchHeaders(headers))
	saga.Add(downgradeBranchURL, "", downgradeBranchBody)
	saga.RetryInterval = 3
	err = saga.Submit()
	if err != nil {
		logger.Warn("saga.Submit error", logger.Err(err), logger.String("dtm gid", gid), interceptor.ServerCtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}

	logger.Info("更新数据，升降级中的DB和缓存强一致性", logger.Any("dtm gid", gid))

	return &stockV1.UpdateDowngradeRequestReply{}, nil
}

// Query  查询
func (s *downgrade) Query(ctx context.Context, req *stockV1.QueryDowngradeRequest) (*stockV1.QueryDowngradeReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	key := getStockCacheKey(req.Id)
	queryFn := func() (string, error) {
		return dao.GetStockByID(s.db, req.Id)
	}

	value, err := s.strongCacheClient.Fetch(key, 300*time.Second, queryFn)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Warn("strongCacheClient.Fetch error", logger.Err(err), logger.String("key", key), interceptor.ServerCtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}

	return &stockV1.QueryDowngradeReply{
		Stock: utils.StrToUint32(value),
	}, nil
}

// DowngradeBranch  升降级中的强一致性分支，注：这是dtm回调，只返回nil、codes.Internal、codes.Aborted错误码
func (s *downgrade) DowngradeBranch(ctx context.Context, req *stockV1.DowngradeBranchRequest) (*stockV1.DowngradeBranchReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}
	update := &model.Stock{
		ID:    req.Id,
		Stock: uint(req.Stock),
	}

	ctx, _ = context.WithTimeout(ctx, 15*time.Second)

	err = s.strongCacheClient.LockForUpdate(ctx, req.Key, req.Gid)
	if err != nil {
		logger.Warn("s.strongCacheClient.LockForUpdate", logger.Err(err), logger.String("key", req.Key), logger.String("gid", req.Gid), interceptor.ServerCtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}

	bb, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		return nil, adaptErr(err)
	}
	err = bb.CallWithDB(s.db, func(tx *sql.Tx) error {
		// if business failed, user should return error dtmcli.ErrFailure
		// other error will be retried
		return dao.UpdateStockInTx(tx, update)
	})
	if err != nil {
		logger.Warn("dao.UpdateStockInTx", logger.Err(err), logger.String("key", req.Key), logger.String("gid", req.Gid), interceptor.ServerCtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}

	err = s.strongCacheClient.UnlockForUpdate(ctx, req.Key, req.Gid)
	if err != nil {
		logger.Warn("s.strongCacheClient.UnlockForUpdate", logger.Err(err), logger.String("key", req.Key), logger.String("gid", req.Gid), interceptor.ServerCtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}

	return &stockV1.DowngradeBranchReply{}, nil
}
