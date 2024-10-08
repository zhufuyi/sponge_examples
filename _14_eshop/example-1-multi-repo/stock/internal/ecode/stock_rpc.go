// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// stock business-level rpc error codes.
// the _stockNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_stockNO       = 79
	_stockName     = "stock"
	_stockBaseCode = errcode.RCode(_stockNO)

	StatusStockDeductStock       = errcode.NewRPCStatus(_stockBaseCode+1, "failed to StockDeduct "+_stockName)
	StatusStockDeductRevertStock = errcode.NewRPCStatus(_stockBaseCode+2, "failed to StockDeductRevert "+_stockName)
	StatusCreateStock            = errcode.NewRPCStatus(_stockBaseCode+3, "failed to Create "+_stockName)
	StatusDeleteByIDStock        = errcode.NewRPCStatus(_stockBaseCode+4, "failed to DeleteByID "+_stockName)
	StatusUpdateByIDStock        = errcode.NewRPCStatus(_stockBaseCode+5, "failed to UpdateByID "+_stockName)
	StatusGetByIDStock           = errcode.NewRPCStatus(_stockBaseCode+6, "failed to GetByID "+_stockName)
	StatusListStock              = errcode.NewRPCStatus(_stockBaseCode+7, "failed to List "+_stockName)
	StatusSetFlashSaleStockStock = errcode.NewRPCStatus(_stockBaseCode+8, "failed to SetFlashSaleStock "+_stockName)

	// error codes are globally unique, adding 1 to the previous error code
)
