// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// collectService rpc service level error code
var (
	_collectServiceNO       = 6 // number range 1~100, if there is the same number, trigger panic.
	_collectServiceName     = "collectService"
	_collectServiceBaseCode = errcode.RCode(_collectServiceNO)

	StatusCreateCollectService = errcode.NewRPCStatus(_collectServiceBaseCode+1, "can not collecting your own posts ")
	StatusDeleteCollectService = errcode.NewRPCStatus(_collectServiceBaseCode+2, "failed to Delete "+_collectServiceName)
	StatusListCollectService   = errcode.NewRPCStatus(_collectServiceBaseCode+3, "failed to List "+_collectServiceName)
	// add +1 to the previous error code
)