// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"errors"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/kafka"
	"github.com/zhufuyi/sponge/pkg/krand"
	"github.com/zhufuyi/sponge/pkg/logger"

	flashSaleV1 "eshop/api/flashSale/v1"
	"eshop/flashSale/internal/config"
	"eshop/flashSale/internal/data"
	"eshop/flashSale/internal/ecode"
	"eshop/flashSale/internal/notify/sender"
	"eshop/flashSale/internal/rpcclient"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		flashSaleV1.RegisterFlashSaleServer(server, NewFlashSaleServer())
	})
}

var _ flashSaleV1.FlashSaleServer = (*flashSale)(nil)

type flashSale struct {
	flashSaleV1.UnimplementedFlashSaleServer

	rdb          *redis.Client
	syncProducer *kafka.SyncProducer
}

// NewFlashSaleServer create a server
func NewFlashSaleServer() flashSaleV1.FlashSaleServer {
	return &flashSale{
		rdb:          data.GetRedisCli(),
		syncProducer: sender.GetSyncProducer(),
	}
}

// FlashSale 秒杀抢购
func (s *flashSale) FlashSale(ctx context.Context, req *flashSaleV1.FlashSaleRequest) (*flashSaleV1.FlashSaleReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	gid := newGid()
	createOrderReq := &flashSaleV1.SendSubmitOrderNotifyRequest{
		UserID:       req.UserID,
		ProductID:    req.ProductID,
		ProductCount: 1,
		Amount:       req.Amount,
		CouponID:     0,
	}
	headers := map[string]string{interceptor.ContextRequestIDKey: interceptor.ServerCtxRequestID(ctx)}

	// 创建二阶段消息事务，1.检查库存是充足， 2.扣减库存，创建订单通知
	msg := dtmgrpc.NewMsgGrpc(rpcclient.GetDtmEndpoint(), gid, dtmgrpc.WithBranchHeaders(headers))
	msg.Add(rpcclient.GetFlashSaleEndpoint()+flashSaleV1.FlashSale_SendSubmitOrderNotify_FullMethodName, createOrderReq)
	err = msg.DoAndSubmit(rpcclient.GetFlashSaleEndpoint()+flashSaleV1.FlashSale_RedisQueryPrepared_FullMethodName, func(bb *dtmcli.BranchBarrier) error {
		return bb.RedisCheckAdjustAmount(s.rdb, data.GetStockCacheKey(req.ProductID), -1, 86400)
	})
	if err != nil {
		if errors.Is(err, dtmimp.ErrFailure) {
			return nil, ecode.StatusAborted.ToRPCErr()
		}
		logger.Warn("DoAndSubmit error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	logger.Info("flashSale success", logger.String("gid", gid), logger.String("trans type", "msg"), interceptor.ServerCtxRequestIDField(ctx))

	return &flashSaleV1.FlashSaleReply{}, nil
}

// RedisQueryPrepared 反查redis数据
func (s *flashSale) RedisQueryPrepared(ctx context.Context, req *flashSaleV1.RedisQueryPreparedRequest) (*flashSaleV1.RedisQueryPreparedReply, error) {
	ctx = interceptor.WrapServerCtx(ctx)

	bb, err := dtmgrpc.BarrierFromGrpc(ctx)
	if err != nil {
		logger.Warn("BarrierFromQuery error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}
	err = bb.RedisQueryPrepared(s.rdb, 7*86400)
	if err != nil {
		logger.Warn("RedisQueryPrepared error", logger.Err(err), interceptor.ServerCtxRequestIDField(ctx))
		return nil, err
	}

	return &flashSaleV1.RedisQueryPreparedReply{}, nil
}

// SendSubmitOrderNotify 发送提交订单通知
func (s *flashSale) SendSubmitOrderNotify(ctx context.Context, req *flashSaleV1.SendSubmitOrderNotifyRequest) (*flashSaleV1.SendSubmitOrderNotifyReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.ToRPCErr()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	// 发送订单通知，这里没有子事务屏障，消费端需要做幂等处理
	partition, offset, err := s.syncProducer.SendData(config.Get().Kafka.OrderTopic, req)
	if err != nil {
		logger.Warn("syncProducer.SendData error", logger.Err(err), logger.String("topic", config.Get().Kafka.OrderTopic), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}
	logger.Info("SendSubmitOrderNotify success", logger.String("topic", config.Get().Kafka.OrderTopic), logger.Int("partition", int(partition)), logger.Int64("offset", offset), interceptor.ServerCtxRequestIDField(ctx))

	return &flashSaleV1.SendSubmitOrderNotifyReply{}, nil
}

func newGid() string {
	return krand.NewSeriesID()
}
