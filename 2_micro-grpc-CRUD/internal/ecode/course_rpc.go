package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// course business-level rpc error codes.
// the _courseNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_courseNO       = 90
	_courseName     = "course"
	_courseBaseCode = errcode.RCode(_courseNO)

	StatusCreateCourse         = errcode.NewRPCStatus(_courseBaseCode+1, "failed to create "+_courseName)
	StatusDeleteByIDCourse     = errcode.NewRPCStatus(_courseBaseCode+2, "failed to delete "+_courseName)
	StatusDeleteByIDsCourse    = errcode.NewRPCStatus(_courseBaseCode+3, "failed to delete by batch ids "+_courseName)
	StatusUpdateByIDCourse     = errcode.NewRPCStatus(_courseBaseCode+4, "failed to update "+_courseName)
	StatusGetByIDCourse        = errcode.NewRPCStatus(_courseBaseCode+5, "failed to get "+_courseName+" details")
	StatusGetByConditionCourse = errcode.NewRPCStatus(_courseBaseCode+6, "failed to get "+_courseName+" by conditions")
	StatusListByIDsCourse      = errcode.NewRPCStatus(_courseBaseCode+7, "failed to list by batch ids "+_courseName)
	StatusListCourse           = errcode.NewRPCStatus(_courseBaseCode+8, "failed to list of "+_courseName)
	// error codes are globally unique, adding 1 to the previous error code
)
