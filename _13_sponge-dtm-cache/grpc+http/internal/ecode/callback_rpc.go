// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// callback business-level rpc error codes.
// the _callbackNO value range is 1~100, if the same error code is used, it will cause panic.
var (
	_callbackNO       = 44
	_callbackName     = "callback"
	_callbackBaseCode = errcode.RCode(_callbackNO)

	StatusQueryPreparedCallback = errcode.NewRPCStatus(_callbackBaseCode+1, "failed to QueryPrepared "+_callbackName)
	StatusDeleteCacheCallback   = errcode.NewRPCStatus(_callbackBaseCode+2, "failed to DeleteCache "+_callbackName)

	// error codes are globally unique, adding 1 to the previous error code
)