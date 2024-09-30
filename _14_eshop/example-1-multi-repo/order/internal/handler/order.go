package handler

import (
	"context"

	orderV1 "order/api/order/v1"
	"order/internal/service"
)

var _ orderV1.OrderLogicer = (*orderHandler)(nil)

type orderHandler struct {
	server orderV1.OrderServer
}

// NewOrderHandler create a handler
func NewOrderHandler() orderV1.OrderLogicer {
	return &orderHandler{
		server: service.NewOrderServer(),
	}
}

// Submit 提交订单(分布式事务)
func (h *orderHandler) Submit(ctx context.Context, req *orderV1.SubmitOrderRequest) (*orderV1.SubmitOrderReply, error) {
	return h.server.Submit(ctx, req)
}

// SendSubmitOrderNotify 发送提交订单通知
func (h *orderHandler) SendSubmitOrderNotify(ctx context.Context, req *orderV1.SendSubmitOrderNotifyRequest) (*orderV1.SendSubmitOrderNotifyReply, error) {
	return h.server.SendSubmitOrderNotify(ctx, req)
}

// Create 创建订单
func (h *orderHandler) Create(ctx context.Context, req *orderV1.CreateOrderRequest) (*orderV1.CreateOrderReply, error) {
	return h.server.Create(ctx, req)
}

// CreateRevert 取消创建订单
func (h *orderHandler) CreateRevert(ctx context.Context, req *orderV1.CreateOrderRevertRequest) (*orderV1.CreateOrderRevertReply, error) {
	return h.server.CreateRevert(ctx, req)
}

// DeleteByID delete a record by id
func (h *orderHandler) DeleteByID(ctx context.Context, req *orderV1.DeleteOrderByIDRequest) (*orderV1.DeleteOrderByIDReply, error) {
	return h.server.DeleteByID(ctx, req)
}

// UpdateByID update a record by id
func (h *orderHandler) UpdateByID(ctx context.Context, req *orderV1.UpdateOrderByIDRequest) (*orderV1.UpdateOrderByIDReply, error) {
	return h.server.UpdateByID(ctx, req)
}

// GetByID get a record by id
func (h *orderHandler) GetByID(ctx context.Context, req *orderV1.GetOrderByIDRequest) (*orderV1.GetOrderByIDReply, error) {
	return h.server.GetByID(ctx, req)
}

// List of records by query parameters
func (h *orderHandler) List(ctx context.Context, req *orderV1.ListOrderRequest) (*orderV1.ListOrderReply, error) {
	return h.server.List(ctx, req)
}