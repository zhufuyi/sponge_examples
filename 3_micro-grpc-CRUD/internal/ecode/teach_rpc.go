// nolint

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teach rpc service level error code
// each resource name corresponds to a unique number (rpc type), the number range is 1~100, if there is the same number, trigger panic
var (
	_teachNO       = 3
	_teachName     = "teach"
	_teachBaseCode = errcode.RCode(_teachNO)

	StatusCreateTeach      = errcode.NewRPCStatus(_teachBaseCode+1, "failed to create "+_teachName)
	StatusDeleteTeach      = errcode.NewRPCStatus(_teachBaseCode+2, "failed to delete "+_teachName)
	StatusDeleteByIDsTeach = errcode.NewRPCStatus(_teachBaseCode+3, "failed to delete by batch ids "+_teachName)
	StatusUpdateTeach      = errcode.NewRPCStatus(_teachBaseCode+4, "failed to update "+_teachName)
	StatusGetTeach         = errcode.NewRPCStatus(_teachBaseCode+5, "failed to get "+_teachName+" details")
	StatusListByIDsTeach   = errcode.NewRPCStatus(_teachBaseCode+6, "failed to list by batch ids "+_teachName)
	StatusListTeach        = errcode.NewRPCStatus(_teachBaseCode+7, "failed to list of "+_teachName)
	// add +1 to the previous error code
)
