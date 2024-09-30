// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// product business-level rpc error codes.
// the _productNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_productNO       = 79
	_productName     = "product"
	_productBaseCode = errcode.RCode(_productNO)

	StatusGetByIDProduct = errcode.NewRPCStatus(_productBaseCode+1, "failed to GetByID "+_productName)

	// error codes are globally unique, adding 1 to the previous error code
)