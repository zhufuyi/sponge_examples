// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// coupon business-level rpc error codes.
// the _couponNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_couponNO       = 52
	_couponName     = "coupon"
	_couponBaseCode = errcode.RCode(_couponNO)

	StatusCouponUseCoupon       = errcode.NewRPCStatus(_couponBaseCode+1, "failed to CouponUse "+_couponName)
	StatusCouponUseRevertCoupon = errcode.NewRPCStatus(_couponBaseCode+2, "failed to CouponUseRevert "+_couponName)
	// error codes are globally unique, adding 1 to the previous error code
)
