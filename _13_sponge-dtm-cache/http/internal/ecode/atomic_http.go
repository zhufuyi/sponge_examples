// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// atomic business-level http error codes.
// the atomicNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	atomicNO       = 80
	atomicName     = "atomic"
	atomicBaseCode = errcode.HCode(atomicNO)

	ErrUpdateAtomic = errcode.NewError(atomicBaseCode+1, "failed to Update "+atomicName)
	ErrQueryAtomic  = errcode.NewError(atomicBaseCode+2, "failed to Query "+atomicName)

	// error codes are globally unique, adding 1 to the previous error code
)