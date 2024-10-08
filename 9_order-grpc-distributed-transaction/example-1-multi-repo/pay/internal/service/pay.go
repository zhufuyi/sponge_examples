// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"database/sql"

	payV1 "pay/api/pay/v1"
	"pay/internal/dao"
	"pay/internal/ecode"
	"pay/internal/model"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/dtm-labs/client/dtmgrpc"
	"google.golang.org/grpc"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		payV1.RegisterPayServer(server, NewPayServer())
	})
}

var _ payV1.PayServer = (*pay)(nil)

type pay struct {
	payV1.UnimplementedPayServer
}

// NewPayServer create a server
func NewPayServer() payV1.PayServer {
	return &pay{}
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
		return dao.CreatePay(tx, req.UserID, req.OrderID, req.Amount)
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
		return dao.CreatePayRevert(tx, req.OrderID)
	})
	if err != nil {
		logger.Error("CreateOrderRevert error", logger.Err(err), logger.Any("orderID", req.OrderID), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &payV1.CreatePayRevertReply{}, nil
}
