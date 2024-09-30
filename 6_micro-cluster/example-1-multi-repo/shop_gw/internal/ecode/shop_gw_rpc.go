// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// shopGw business-level rpc error codes.
// the shopGwNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_shopGwNO       = 48
	_shopGwName     = "shopGw"
	_shopGwBaseCode = errcode.RCode(_shopGwNO)

	StatusGetDetailsByProductIDShopGw   = errcode.NewRPCStatus(_shopGwBaseCode+1, "failed to GetDetailsByProductID "+_shopGwName)
	// error codes are globally unique, adding 1 to the previous error code
)
