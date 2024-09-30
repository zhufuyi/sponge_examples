// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// order business-level rpc error codes.
// the _orderNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_orderNO       = 3
	_orderName     = "order"
	_orderBaseCode = errcode.RCode(_orderNO)

	StatusSubmitOrder                = errcode.NewRPCStatus(_orderBaseCode+1, "failed to Submit "+_orderName)
	StatusCreateOrder                = errcode.NewRPCStatus(_orderBaseCode+2, "failed to Create "+_orderName)
	StatusCreateRevertOrder          = errcode.NewRPCStatus(_orderBaseCode+3, "failed to CreateRevert "+_orderName)
	StatusDeleteByIDOrder            = errcode.NewRPCStatus(_orderBaseCode+4, "failed to DeleteByID "+_orderName)
	StatusUpdateByIDOrder            = errcode.NewRPCStatus(_orderBaseCode+5, "failed to UpdateByID "+_orderName)
	StatusGetByIDOrder               = errcode.NewRPCStatus(_orderBaseCode+6, "failed to GetByID "+_orderName)
	StatusListOrder                  = errcode.NewRPCStatus(_orderBaseCode+7, "failed to List "+_orderName)
	StatusSendSubmitOrderNotifyOrder = errcode.NewRPCStatus(_orderBaseCode+8, "failed to SendSubmitOrderNotify "+_orderName)

	// error codes are globally unique, adding 1 to the previous error code
)
