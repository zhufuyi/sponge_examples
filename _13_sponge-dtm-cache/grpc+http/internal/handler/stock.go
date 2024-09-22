package handler

import (
	"context"

	stockV1 "stock/api/stock/v1"
	"stock/internal/service"
)

var _ stockV1.StockLogicer = (*stockHandler)(nil)

type stockHandler struct {
	server stockV1.StockServer
}

// NewStockHandler create a handler
func NewStockHandler() stockV1.StockLogicer {
	return &stockHandler{
		server: service.NewStockServer(),
	}
}

// Create a record
func (h *stockHandler) Create(ctx context.Context, req *stockV1.CreateStockRequest) (*stockV1.CreateStockReply, error) {
	return h.server.Create(ctx, req)
}

// DeleteByID delete a record by id
func (h *stockHandler) DeleteByID(ctx context.Context, req *stockV1.DeleteStockByIDRequest) (*stockV1.DeleteStockByIDReply, error) {
	return h.server.DeleteByID(ctx, req)
}

// UpdateByID update a record by id
func (h *stockHandler) UpdateByID(ctx context.Context, req *stockV1.UpdateStockByIDRequest) (*stockV1.UpdateStockByIDReply, error) {
	return h.server.UpdateByID(ctx, req)
}

// GetByID get a record by id
func (h *stockHandler) GetByID(ctx context.Context, req *stockV1.GetStockByIDRequest) (*stockV1.GetStockByIDReply, error) {
	return h.server.GetByID(ctx, req)
}

// List of records by query parameters
func (h *stockHandler) List(ctx context.Context, req *stockV1.ListStockRequest) (*stockV1.ListStockReply, error) {
	return h.server.List(ctx, req)
}
