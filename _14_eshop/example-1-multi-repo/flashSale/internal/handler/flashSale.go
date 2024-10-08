// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"
	"errors"
	"flashSale/internal/ecode"
	"github.com/zhufuyi/sponge/pkg/gin/middleware"

	flashSaleV1 "flashSale/api/flashSale/v1"
	"flashSale/internal/service"
)

var _ flashSaleV1.FlashSaleLogicer = (*flashSaleHandler)(nil)

type flashSaleHandler struct {
	server flashSaleV1.FlashSaleServer
}

// NewFlashSaleHandler create a handler
func NewFlashSaleHandler() flashSaleV1.FlashSaleLogicer {
	return &flashSaleHandler{
		server: service.NewFlashSaleServer(),
	}
}

// FlashSale 秒杀抢购
func (h *flashSaleHandler) FlashSale(ctx context.Context, req *flashSaleV1.FlashSaleRequest) (*flashSaleV1.FlashSaleReply, error) {
	reply, err := h.server.FlashSale(ctx, req)
	if err != nil {
		if errors.Is(err, ecode.StatusAborted.ToRPCErr()) {
			//return nil, ecode.Conflict.WithOutMsg("销售完毕，请下次再来~").ErrToHTTP() // 返回409错误码
			return nil, ecode.StatusForbidden.ErrToHTTP() // 返回409错误码
		} else if errors.Is(err, ecode.StatusInternalServerError.ToRPCErr()) {
			return nil, ecode.StatusInternalServerError.ErrToHTTP() // 返回500错误码
		}
	}
	return reply, err
}

// RedisQueryPrepared 反查redis数据
func (h *flashSaleHandler) RedisQueryPrepared(ctx context.Context, req *flashSaleV1.RedisQueryPreparedRequest) (*flashSaleV1.RedisQueryPreparedReply, error) {
	_, ctx = middleware.AdaptCtx(ctx)
	return h.server.RedisQueryPrepared(ctx, req)
}

// SendSubmitOrderNotify 发送提交订单通知
func (h *flashSaleHandler) SendSubmitOrderNotify(ctx context.Context, req *flashSaleV1.SendSubmitOrderNotifyRequest) (*flashSaleV1.SendSubmitOrderNotifyReply, error) {
	return h.server.SendSubmitOrderNotify(ctx, req)
}
