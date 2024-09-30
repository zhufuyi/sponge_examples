// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// order business-level rpc error codes.
// the orderNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_orderNO       = 60
	_orderName     = "order"
	_orderBaseCode = errcode.RCode(_orderNO)

	StatusSubmitOrder = errcode.NewRPCStatus(_orderBaseCode+1, "failed to Submit "+_orderName)

	// error codes are globally unique, adding 1 to the previous error code
)