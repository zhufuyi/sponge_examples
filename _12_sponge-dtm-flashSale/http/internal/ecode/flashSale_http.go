// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// flashSale business-level http error codes.
// the flashSaleNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	flashSaleNO       = 1
	flashSaleName     = "flashSale"
	flashSaleBaseCode = errcode.HCode(flashSaleNO)

	ErrSetProductStockFlashSale       = errcode.NewError(flashSaleBaseCode+1, "failed to SetProductStock "+flashSaleName)
	ErrFlashSaleFlashSale             = errcode.NewError(flashSaleBaseCode+2, "failed to FlashSale")
	ErrRedisQueryPreparedFlashSale    = errcode.NewError(flashSaleBaseCode+3, "failed to RedisQueryPrepared "+flashSaleName)
	ErrSendSubmitOrderNotifyFlashSale = errcode.NewError(flashSaleBaseCode+4, "failed to SendSubmitOrderNotify "+flashSaleName)

	// error codes are globally unique, adding 1 to the previous error code
)
