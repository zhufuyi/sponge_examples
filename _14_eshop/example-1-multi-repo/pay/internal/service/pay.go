package service

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	payV1 "pay/api/pay/v1"
	"pay/internal/cache"
	"pay/internal/dao"
	"pay/internal/ecode"
	"pay/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		payV1.RegisterPayServer(server, NewPayServer()) // register service to the rpc service
	})
}

var _ payV1.PayServer = (*pay)(nil)
var _ time.Time

type pay struct {
	payV1.UnimplementedPayServer

	iDao dao.PayDao
}

// NewPayServer create a new service
func NewPayServer() payV1.PayServer {
	return &pay{
		iDao: dao.NewPayDao(
			model.GetDB(),
			cache.NewPayCache(model.GetCacheType()),
		),
	}
}

// Create 创建支付订单
func (s *pay) Create(ctx context.Context, req *payV1.CreatePayRequest) (*payV1.CreatePayReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	// 使用子事务屏障
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		logger.Error("BarrierFromGrpc error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusCreatePay.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.CreatePay(tx, req.UserID, req.OrderID, req.Amount)
	})
	if err != nil {
		logger.Error("CreatePay error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &payV1.CreatePayReply{}, nil
}

// CreateRevert 取消支付订单
func (s *pay) CreateRevert(ctx context.Context, req *payV1.CreatePayRevertRequest) (*payV1.CreatePayRevertReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	// 使用子事务屏障
	barrier, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		logger.Error("BarrierFromGrpc error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusCreateRevertPay.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.CreatePayRevert(tx, req.OrderID)
	})
	if err != nil {
		logger.Error("CreateOrderRevert error", logger.Err(err), logger.Any("orderID", req.OrderID), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &payV1.CreatePayRevertReply{}, nil
}

// DeleteByID delete a record by id
func (s *pay) DeleteByID(ctx context.Context, req *payV1.DeletePayByIDRequest) (*payV1.DeletePayByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	err = s.iDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &payV1.DeletePayByIDReply{}, nil
}

// UpdateByID update a record by id
func (s *pay) UpdateByID(ctx context.Context, req *payV1.UpdatePayByIDRequest) (*payV1.UpdatePayByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.Pay{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDPay.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	record.ID = req.Id

	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("pay", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &payV1.UpdatePayByIDReply{}, nil
}

// GetByID get a record by id
func (s *pay) GetByID(ctx context.Context, req *payV1.GetPayByIDRequest) (*payV1.GetPayByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record, err := s.iDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertPay(record)
	if err != nil {
		logger.Warn("convertPay error", logger.Err(err), logger.Any("pay", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDPay.Err()
	}

	return &payV1.GetPayByIDReply{Pay: data}, nil
}

// List of records by query parameters
func (s *pay) List(ctx context.Context, req *payV1.ListPayRequest) (*payV1.ListPayReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListPay.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	pays := []*payV1.Pay{}
	for _, record := range records {
		data, err := convertPay(record)
		if err != nil {
			logger.Warn("convertPay error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		pays = append(pays, data)
	}

	return &payV1.ListPayReply{
		Total: total,
		Pays:  pays,
	}, nil
}

func convertPay(record *model.Pay) (*payV1.Pay, error) {
	value := &payV1.Pay{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID
	value.CreatedAt = utils.FormatDateTimeRFC3339(*record.CreatedAt)
	value.UpdatedAt = utils.FormatDateTimeRFC3339(*record.UpdatedAt)

	return value, nil
}
