package handler

import (
	"context"

	couponV1 "coupon/api/coupon/v1"
	"coupon/internal/service"
)

var _ couponV1.CouponLogicer = (*couponHandler)(nil)

type couponHandler struct {
	server couponV1.CouponServer
}

// NewCouponHandler create a handler
func NewCouponHandler() couponV1.CouponLogicer {
	return &couponHandler{
		server: service.NewCouponServer(),
	}
}

// CouponUse 使用优惠券
func (h *couponHandler) CouponUse(ctx context.Context, req *couponV1.CouponUseRequest) (*couponV1.CouponUseReply, error) {
	return h.server.CouponUse(ctx, req)
}

// CouponUseRevert 补偿优惠券
func (h *couponHandler) CouponUseRevert(ctx context.Context, req *couponV1.CouponUseRevertRequest) (*couponV1.CouponUseRevertReply, error) {
	return h.server.CouponUseRevert(ctx, req)
}

// Create a record
func (h *couponHandler) Create(ctx context.Context, req *couponV1.CreateCouponRequest) (*couponV1.CreateCouponReply, error) {
	return h.server.Create(ctx, req)
}

// DeleteByID delete a record by id
func (h *couponHandler) DeleteByID(ctx context.Context, req *couponV1.DeleteCouponByIDRequest) (*couponV1.DeleteCouponByIDReply, error) {
	return h.server.DeleteByID(ctx, req)
}

// UpdateByID update a record by id
func (h *couponHandler) UpdateByID(ctx context.Context, req *couponV1.UpdateCouponByIDRequest) (*couponV1.UpdateCouponByIDReply, error) {
	return h.server.UpdateByID(ctx, req)
}

// GetByID get a record by id
func (h *couponHandler) GetByID(ctx context.Context, req *couponV1.GetCouponByIDRequest) (*couponV1.GetCouponByIDReply, error) {
	return h.server.GetByID(ctx, req)
}

// List of records by query parameters
func (h *couponHandler) List(ctx context.Context, req *couponV1.ListCouponRequest) (*couponV1.ListCouponReply, error) {
	return h.server.List(ctx, req)
}
