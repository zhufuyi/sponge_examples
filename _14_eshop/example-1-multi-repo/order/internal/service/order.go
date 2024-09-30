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
	"github.com/zhufuyi/sponge/pkg/kafka"
	"github.com/zhufuyi/sponge/pkg/krand"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	orderV1 "order/api/order/v1"
	"order/internal/cache"
	"order/internal/config"
	"order/internal/dao"
	"order/internal/ecode"
	"order/internal/model"
	"order/internal/notify/sender"
	"order/internal/rpcclient"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		orderV1.RegisterOrderServer(server, NewOrderServer()) // register service to the rpc service
	})
}

var _ orderV1.OrderServer = (*order)(nil)
var _ time.Time

type order struct {
	orderV1.UnimplementedOrderServer

	iDao         dao.OrderDao
	syncProducer *kafka.SyncProducer
}

// NewOrderServer create a new service
func NewOrderServer() orderV1.OrderServer {
	return &order{
		iDao: dao.NewOrderDao(
			model.GetDB(),
			cache.NewOrderCache(model.GetCacheType()),
		),
		syncProducer: sender.GetSyncProducer(),
	}
}

// SendSubmitOrderNotify 发送提交订单通知
func (s *order) SendSubmitOrderNotify(ctx context.Context, req *orderV1.SendSubmitOrderNotifyRequest) (*orderV1.SendSubmitOrderNotifyReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	partition, offset, err := s.syncProducer.SendData(config.Get().Kafka.OrderTopic, req)
	if err != nil {
		logger.Error("syncProducer.SendData error", logger.Err(err), logger.String("topic", config.Get().Kafka.OrderTopic), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}
	logger.Info("SendSubmitOrderNotify success", logger.String("topic", config.Get().Kafka.OrderTopic), logger.Int("partition", int(partition)), logger.Int64("offset", offset), interceptor.ServerCtxRequestIDField(ctx))

	return &orderV1.SendSubmitOrderNotifyReply{}, nil
}

// Submit 提交订单
func (s *order) Submit(ctx context.Context, req *orderV1.SubmitOrderRequest) (*orderV1.SubmitOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	orderID := newGid()
	gid := orderID // 订单id也是dtm的gid
	stockReq := &orderV1.StockDeductRequest{
		ProductID:    req.ProductID,
		ProductCount: req.ProductCount,
	}
	orderReq := &orderV1.CreateOrderRequest{
		OrderID:      orderID,
		UserID:       req.UserID,
		ProductID:    req.ProductID,
		ProductCount: req.ProductCount,
		Amount:       req.Amount,
		CouponID:     req.CouponID,
	}
	payReq := &orderV1.CreatePayRequest{
		UserID:  req.UserID,
		OrderID: orderID,
		Amount:  req.Amount,
	}
	headers := map[string]string{interceptor.ContextRequestIDKey: interceptor.ServerCtxRequestID(ctx)}

	// 使用saga模式分布式事务
	saga := dtmgrpc.NewSagaGrpc(rpcclient.GetDtmEndpoint(), gid, dtmgrpc.WithBranchHeaders(headers))
	saga.Add(rpcclient.GetOrderEndpoint()+orderV1.Order_Create_FullMethodName, rpcclient.GetOrderEndpoint()+orderV1.Order_CreateRevert_FullMethodName, orderReq)
	saga.Add(rpcclient.GetStockEndpoint()+"/api.stock.v1.stock/StockDeduct", rpcclient.GetStockEndpoint()+"/api.stock.v1.stock/StockDeductRevert", stockReq)
	if req.CouponID > 0 {
		couponReq := &orderV1.CouponUseRequest{CouponID: req.CouponID}
		saga.Add(rpcclient.GetCouponEndpoint()+"/api.coupon.v1.coupon/CouponUse", rpcclient.GetCouponEndpoint()+"/api.coupon.v1.coupon/CouponUseRevert", couponReq)
	}
	saga.Add(rpcclient.GetPayEndpoint()+"/api.pay.v1.pay/Create", rpcclient.GetPayEndpoint()+"/api.pay.v1.pay/CreateRevert", payReq)
	err = saga.Submit()
	if err != nil {
		logger.Error("submit error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &orderV1.SubmitOrderReply{
		OrderID: orderID,
	}, nil
}

// Create 创建订单
func (s *order) Create(ctx context.Context, req *orderV1.CreateOrderRequest) (*orderV1.CreateOrderReply, error) {
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
		return nil, ecode.StatusCreateOrder.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.CreateOrder(tx, req.OrderID, req.UserID, req.ProductID, req.Amount, req.ProductCount, req.CouponID)
	})
	if err != nil {
		logger.Error("CreateOrder error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &orderV1.CreateOrderReply{}, nil
}

// CreateRevert 取消创建订单
func (s *order) CreateRevert(ctx context.Context, req *orderV1.CreateOrderRevertRequest) (*orderV1.CreateOrderRevertReply, error) {
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
		return nil, ecode.StatusCreateOrder.Err()
	}

	tx := model.GetDB().Begin().Statement.ConnPool.(*sql.Tx)
	err = barrier.Call(tx, func(tx *sql.Tx) error {
		return s.iDao.CreateOrderRevert(tx, req.OrderID)
	})
	if err != nil {
		logger.Error("CreateOrderRevert error", logger.Err(err), logger.Any("orderID", req.OrderID), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &orderV1.CreateOrderRevertReply{}, nil
}

// DeleteByID delete a record by id
func (s *order) DeleteByID(ctx context.Context, req *orderV1.DeleteOrderByIDRequest) (*orderV1.DeleteOrderByIDReply, error) {
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

	return &orderV1.DeleteOrderByIDReply{}, nil
}

// UpdateByID update a record by id
func (s *order) UpdateByID(ctx context.Context, req *orderV1.UpdateOrderByIDRequest) (*orderV1.UpdateOrderByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.OrderRecord{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDOrder.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	record.ID = req.Id

	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("order", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &orderV1.UpdateOrderByIDReply{}, nil
}

// GetByID get a record by id
func (s *order) GetByID(ctx context.Context, req *orderV1.GetOrderByIDRequest) (*orderV1.GetOrderByIDReply, error) {
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

	data, err := convertOrder(record)
	if err != nil {
		logger.Warn("convertOrder error", logger.Err(err), logger.Any("order", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDOrder.Err()
	}

	return &orderV1.GetOrderByIDReply{Order: data}, nil
}

// List of records by query parameters
func (s *order) List(ctx context.Context, req *orderV1.ListOrderRequest) (*orderV1.ListOrderReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListOrder.Err()
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

	orders := []*orderV1.Order{}
	for _, record := range records {
		data, err := convertOrder(record)
		if err != nil {
			logger.Warn("convertOrder error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		orders = append(orders, data)
	}

	return &orderV1.ListOrderReply{
		Total:  total,
		Orders: orders,
	}, nil
}

func convertOrder(record *model.OrderRecord) (*orderV1.Order, error) {
	value := &orderV1.Order{}
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

func newGid() string {
	return krand.NewSeriesID()
}
