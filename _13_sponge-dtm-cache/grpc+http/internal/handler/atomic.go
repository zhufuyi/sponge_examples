// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"

	stockV1 "stock/api/stock/v1"
	"stock/internal/service"
)

var _ stockV1.AtomicLogicer = (*atomicHandler)(nil)

type atomicHandler struct {
	server stockV1.AtomicServer
}

// NewAtomicHandler create a handler
func NewAtomicHandler() stockV1.AtomicLogicer {
	return &atomicHandler{
		server: service.NewAtomicServer(),
	}
}

// Update 更新数据，保证DB与缓存操作的原子性。如果更新完DB后，redis发生进程crash，重新启动redis服务后，也能够保证缓存能够更新。
func (h *atomicHandler) Update(ctx context.Context, req *stockV1.UpdateAtomicRequest) (*stockV1.UpdateAtomicRequestReply, error) {

	return h.server.Update(ctx, req)
}

// Query  查询
func (h *atomicHandler) Query(ctx context.Context, req *stockV1.QueryAtomicRequest) (*stockV1.QueryAtomicReply, error) {

	return h.server.Query(ctx, req)
}
