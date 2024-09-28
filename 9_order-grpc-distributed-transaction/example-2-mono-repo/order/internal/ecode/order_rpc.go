// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// order business-level rpc error codes.
// the _orderNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_orderNO       = 97
	_orderName     = "order"
	_orderBaseCode = errcode.RCode(_orderNO)

	StatusSubmitOrder       = errcode.NewRPCStatus(_orderBaseCode+1, "failed to Submit "+_orderName)
	StatusCreateOrder       = errcode.NewRPCStatus(_orderBaseCode+2, "failed to Create "+_orderName)
	StatusCreateRevertOrder = errcode.NewRPCStatus(_orderBaseCode+3, "failed to CreateRevert "+_orderName)

	// error codes are globally unique, adding 1 to the previous error code
)
