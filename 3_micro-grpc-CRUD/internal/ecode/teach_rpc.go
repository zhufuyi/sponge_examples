package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teach business-level rpc error codes.
// the _teachNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_teachNO       = 99
	_teachName     = "teach"
	_teachBaseCode = errcode.RCode(_teachNO)

	StatusCreateTeach         = errcode.NewRPCStatus(_teachBaseCode+1, "failed to create "+_teachName)
	StatusDeleteByIDTeach     = errcode.NewRPCStatus(_teachBaseCode+2, "failed to delete "+_teachName)
	StatusDeleteByIDsTeach    = errcode.NewRPCStatus(_teachBaseCode+3, "failed to delete by batch ids "+_teachName)
	StatusUpdateByIDTeach     = errcode.NewRPCStatus(_teachBaseCode+4, "failed to update "+_teachName)
	StatusGetByIDTeach        = errcode.NewRPCStatus(_teachBaseCode+5, "failed to get "+_teachName+" details")
	StatusGetByConditionTeach = errcode.NewRPCStatus(_teachBaseCode+6, "failed to get "+_teachName+" by conditions")
	StatusListByIDsTeach      = errcode.NewRPCStatus(_teachBaseCode+7, "failed to list by batch ids "+_teachName)
	StatusListTeach           = errcode.NewRPCStatus(_teachBaseCode+8, "failed to list of "+_teachName)
	// error codes are globally unique, adding 1 to the previous error code
)
