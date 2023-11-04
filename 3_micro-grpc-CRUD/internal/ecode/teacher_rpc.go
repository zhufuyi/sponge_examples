package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teacher business-level rpc error codes.
// the _teacherNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_teacherNO       = 32
	_teacherName     = "teacher"
	_teacherBaseCode = errcode.RCode(_teacherNO)

	StatusCreateTeacher         = errcode.NewRPCStatus(_teacherBaseCode+1, "failed to create "+_teacherName)
	StatusDeleteByIDTeacher     = errcode.NewRPCStatus(_teacherBaseCode+2, "failed to delete "+_teacherName)
	StatusDeleteByIDsTeacher    = errcode.NewRPCStatus(_teacherBaseCode+3, "failed to delete by batch ids "+_teacherName)
	StatusUpdateByIDTeacher     = errcode.NewRPCStatus(_teacherBaseCode+4, "failed to update "+_teacherName)
	StatusGetByIDTeacher        = errcode.NewRPCStatus(_teacherBaseCode+5, "failed to get "+_teacherName+" details")
	StatusGetByConditionTeacher = errcode.NewRPCStatus(_teacherBaseCode+6, "failed to get "+_teacherName+" by conditions")
	StatusListByIDsTeacher      = errcode.NewRPCStatus(_teacherBaseCode+7, "failed to list by batch ids "+_teacherName)
	StatusListTeacher           = errcode.NewRPCStatus(_teacherBaseCode+8, "failed to list of "+_teacherName)
	// error codes are globally unique, adding 1 to the previous error code
)
