// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"

	stockV1 "stock/api/stock/v1"
	"stock/internal/service"
)

var _ stockV1.FinalLogicer = (*finalHandler)(nil)

type finalHandler struct {
	server stockV1.FinalServer
}

// NewFinalHandler create a handler
func NewFinalHandler() stockV1.FinalLogicer {
	return &finalHandler{
		server: service.NewFinalServer(),
	}
}

// Update 更新数据，DB和缓存最终一致性
func (h *finalHandler) Update(ctx context.Context, req *stockV1.UpdateFinalRequest) (*stockV1.UpdateFinalRequestReply, error) {

	return h.server.Update(ctx, req)
}

// Query  查询
func (h *finalHandler) Query(ctx context.Context, req *stockV1.QueryFinalRequest) (*stockV1.QueryFinalReply, error) {

	return h.server.Query(ctx, req)
}
