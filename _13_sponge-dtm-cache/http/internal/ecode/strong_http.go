// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// strong business-level http error codes.
// the strongNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	strongNO       = 64
	strongName     = "strong"
	strongBaseCode = errcode.HCode(strongNO)

	ErrUpdateStrong = errcode.NewError(strongBaseCode+1, "failed to Update "+strongName)
	ErrQueryStrong  = errcode.NewError(strongBaseCode+2, "failed to Query "+strongName)

	// error codes are globally unique, adding 1 to the previous error code
)