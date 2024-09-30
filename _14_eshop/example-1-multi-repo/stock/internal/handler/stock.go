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

// StockDeduct 扣减库存
func (h *stockHandler) StockDeduct(ctx context.Context, req *stockV1.StockDeductRequest) (*stockV1.StockDeductReply, error) {
	return h.server.StockDeduct(ctx, req)
}

// StockDeductRevert 补偿库存
func (h *stockHandler) StockDeductRevert(ctx context.Context, req *stockV1.StockDeductRevertRequest) (*stockV1.StockDeductRevertReply, error) {
	return h.server.StockDeductRevert(ctx, req)
}

// SetFlashSaleStock 设置秒杀产品的库存
func (h *stockHandler) SetFlashSaleStock(ctx context.Context, req *stockV1.SetFlashSaleStockRequest) (*stockV1.SetFlashSaleStockReply, error) {
	return h.server.SetFlashSaleStock(ctx, req)
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
