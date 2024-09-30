// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/rockscache"
	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	stockV1 "stock/api/stock/v1"
	"stock/internal/config"
	"stock/internal/dao"
	"stock/internal/ecode"
	"stock/internal/model"
)

var _ stockV1.StrongLogicer = (*strongHandler)(nil)

type strongHandler struct {
	db                *sql.DB
	strongCacheClient *rockscache.Client
}

// NewStrongHandler create a handler
func NewStrongHandler() stockV1.StrongLogicer {
	return &strongHandler{
		db:                model.GetSDB(),
		strongCacheClient: model.GetStrongRockscacheClient(),
	}
}

// Update 更新数据，DB和缓存强一致性
func (h *strongHandler) Update(ctx context.Context, req *stockV1.UpdateStrongRequest) (*stockV1.UpdateStrongRequestReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}

	gid := newGid()
	callbackStockAddr := getCallbackStockAddr()
	deleteCacheURL := callbackStockAddr + "/api/v1/stock/deleteCache"
	queryPreparedURL := callbackStockAddr + "/api/v1/stock/queryPrepared"
	deleteCacheBody := &stockV1.DeleteCacheRequest{
		Key: getStockCacheKey(req.Id),
	}
	stock := &model.Stock{
		ID:    req.Id,
		Stock: uint(req.Stock),
	}

	// 创建二阶段消息事务
	msg := dtmcli.NewMsg(config.Get().Dtm.Server, gid)
	msg.Add(deleteCacheURL, deleteCacheBody)
	msg.WaitResult = false // when return success, the global transaction has finished
	err = msg.DoAndSubmit(queryPreparedURL, func(bb *dtmcli.BranchBarrier) error {
		return bb.CallWithDB(h.db, func(tx *sql.Tx) error {
			return dao.UpdateStockInTx(tx, stock)
		})
	})
	if err != nil {
		logger.Warn("msg.DoAndSubmit error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, adaptErr(err)
	}

	logger.Info("更新数据，DB和缓存强一致性", logger.Err(err), logger.Any("dtm gid", gid))

	return &stockV1.UpdateStrongRequestReply{}, nil
}

// Query  查询
func (h *strongHandler) Query(ctx context.Context, req *stockV1.QueryStrongRequest) (*stockV1.QueryStrongReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	key := getStockCacheKey(req.Id)
	query := func() (string, error) {
		return dao.GetStockByID(h.db, req.Id)
	}

	value, err := h.strongCacheClient.Fetch(key, 300*time.Second, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ecode.NotFound.Err()
		}
		logger.Warn("fetch cache error", logger.Err(err), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InternalServerError.Err()
	}

	return &stockV1.QueryStrongReply{
		Stock: utils.StrToUint32(value),
	}, nil
}