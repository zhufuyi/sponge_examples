package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// stock business-level http error codes.
// the stockNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	stockNO       = 62
	stockName     = "stock"
	stockBaseCode = errcode.HCode(stockNO)

	ErrCreateStock     = errcode.NewError(stockBaseCode+1, "failed to create "+stockName)
	ErrDeleteByIDStock = errcode.NewError(stockBaseCode+2, "failed to delete "+stockName)
	ErrUpdateByIDStock = errcode.NewError(stockBaseCode+3, "failed to update "+stockName)
	ErrGetByIDStock    = errcode.NewError(stockBaseCode+4, "failed to get "+stockName+" details")
	ErrListStock       = errcode.NewError(stockBaseCode+5, "failed to list of "+stockName)

	// error codes are globally unique, adding 1 to the previous error code
)
