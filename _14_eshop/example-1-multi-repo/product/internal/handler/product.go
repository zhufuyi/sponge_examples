package handler

import (
	"context"

	productV1 "product/api/product/v1"
	"product/internal/service"
)

var _ productV1.ProductLogicer = (*productHandler)(nil)

type productHandler struct {
	server productV1.ProductServer
}

// NewProductHandler create a handler
func NewProductHandler() productV1.ProductLogicer {
	return &productHandler{
		server: service.NewProductServer(),
	}
}

// Create a record
func (h *productHandler) Create(ctx context.Context, req *productV1.CreateProductRequest) (*productV1.CreateProductReply, error) {
	return h.server.Create(ctx, req)
}

// DeleteByID delete a record by id
func (h *productHandler) DeleteByID(ctx context.Context, req *productV1.DeleteProductByIDRequest) (*productV1.DeleteProductByIDReply, error) {
	return h.server.DeleteByID(ctx, req)
}

// UpdateByID update a record by id
func (h *productHandler) UpdateByID(ctx context.Context, req *productV1.UpdateProductByIDRequest) (*productV1.UpdateProductByIDReply, error) {
	return h.server.UpdateByID(ctx, req)
}

// GetByID get a record by id
func (h *productHandler) GetByID(ctx context.Context, req *productV1.GetProductByIDRequest) (*productV1.GetProductByIDReply, error) {
	return h.server.GetByID(ctx, req)
}

// List of records by query parameters
func (h *productHandler) List(ctx context.Context, req *productV1.ListProductRequest) (*productV1.ListProductReply, error) {

	return h.server.List(ctx, req)
}
