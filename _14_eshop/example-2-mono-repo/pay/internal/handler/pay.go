// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"

	payV1 "eshop/api/pay/v1"
	"eshop/pay/internal/service"
)

var _ payV1.PayLogicer = (*payHandler)(nil)

type payHandler struct {
	server payV1.PayServer
}

// NewPayHandler create a handler
func NewPayHandler() payV1.PayLogicer {
	return &payHandler{
		server: service.NewPayServer(),
	}
}

// Create 创建支付订单
func (h *payHandler) Create(ctx context.Context, req *payV1.CreatePayRequest) (*payV1.CreatePayReply, error) {

	return h.server.Create(ctx, req)
}

// CreateRevert 取消支付订单
func (h *payHandler) CreateRevert(ctx context.Context, req *payV1.CreatePayRevertRequest) (*payV1.CreatePayRevertReply, error) {

	return h.server.CreateRevert(ctx, req)
}

// DeleteByID delete pay by id
func (h *payHandler) DeleteByID(ctx context.Context, req *payV1.DeletePayByIDRequest) (*payV1.DeletePayByIDReply, error) {

	return h.server.DeleteByID(ctx, req)
}

// UpdateByID update pay by id
func (h *payHandler) UpdateByID(ctx context.Context, req *payV1.UpdatePayByIDRequest) (*payV1.UpdatePayByIDReply, error) {

	return h.server.UpdateByID(ctx, req)
}

// GetByID get pay by id
func (h *payHandler) GetByID(ctx context.Context, req *payV1.GetPayByIDRequest) (*payV1.GetPayByIDReply, error) {

	return h.server.GetByID(ctx, req)
}

// List of pay by query parameters
func (h *payHandler) List(ctx context.Context, req *payV1.ListPayRequest) (*payV1.ListPayReply, error) {

	return h.server.List(ctx, req)
}
