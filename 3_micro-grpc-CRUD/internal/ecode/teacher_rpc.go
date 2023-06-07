// nolint

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teacher rpc service level error code
// each resource name corresponds to a unique number (rpc type), the number range is 1~100, if there is the same number, trigger panic
var (
	_teacherNO       = 1
	_teacherName     = "teacher"
	_teacherBaseCode = errcode.RCode(_teacherNO)

	StatusCreateTeacher      = errcode.NewRPCStatus(_teacherBaseCode+1, "failed to create "+_teacherName)
	StatusDeleteTeacher      = errcode.NewRPCStatus(_teacherBaseCode+2, "failed to delete "+_teacherName)
	StatusDeleteByIDsTeacher = errcode.NewRPCStatus(_teacherBaseCode+3, "failed to delete by batch ids "+_teacherName)
	StatusUpdateTeacher      = errcode.NewRPCStatus(_teacherBaseCode+4, "failed to update "+_teacherName)
	StatusGetTeacher         = errcode.NewRPCStatus(_teacherBaseCode+5, "failed to get "+_teacherName+" details")
	StatusListByIDsTeacher   = errcode.NewRPCStatus(_teacherBaseCode+6, "failed to list by batch ids "+_teacherName)
	StatusListTeacher        = errcode.NewRPCStatus(_teacherBaseCode+7, "failed to list of "+_teacherName)
	// add +1 to the previous error code
)
