// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// collectService business-level rpc error codes.
// the _collectServiceNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_collectServiceNO       = 8
	_collectServiceName     = "collectService"
	_collectServiceBaseCode = errcode.RCode(_collectServiceNO)

	StatusCreateCollectService = errcode.NewRPCStatus(_collectServiceBaseCode+1, "failed to Create "+_collectServiceName)
	StatusDeleteCollectService = errcode.NewRPCStatus(_collectServiceBaseCode+2, "failed to Delete "+_collectServiceName)
	StatusListCollectService   = errcode.NewRPCStatus(_collectServiceBaseCode+3, "failed to List "+_collectServiceName)
	// error codes are globally unique, adding 1 to the previous error code
)
