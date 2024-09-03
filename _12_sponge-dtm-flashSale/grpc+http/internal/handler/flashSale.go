// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"

	"google.golang.org/grpc/codes"

	flashSaleV1 "flashSale/api/flashSale/v1"
	"flashSale/internal/ecode"
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

// SetProductStock 设置库存数量
func (h *flashSaleHandler) SetProductStock(ctx context.Context, req *flashSaleV1.SetProductStockRequest) (*flashSaleV1.SetProductStockReply, error) {
	return h.server.SetProductStock(ctx, req)
}

// FlashSale 秒杀抢购
func (h *flashSaleHandler) FlashSale(ctx context.Context, req *flashSaleV1.FlashSaleRequest) (*flashSaleV1.FlashSaleReply, error) {
	reply, err := h.server.FlashSale(ctx, req)
	if err != nil {
		switch ecode.GetStatusCode(err) {
		case codes.Aborted, ecode.StatusAborted.Code():
			return nil, ecode.StatusConflict.ErrToHTTP("已售罄，欢迎下次再来抢购") // 返回409错误码
		case codes.Internal, ecode.StatusInternalServerError.Code():
			return nil, ecode.StatusInternalServerError.ErrToHTTP() // 返回500错误码
		}
	}
	return reply, err
}
