// nolint

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// course rpc service level error code
// each resource name corresponds to a unique number (rpc type), the number range is 1~100, if there is the same number, trigger panic
var (
	_courseNO       = 2
	_courseName     = "course"
	_courseBaseCode = errcode.RCode(_courseNO)

	StatusCreateCourse      = errcode.NewRPCStatus(_courseBaseCode+1, "failed to create "+_courseName)
	StatusDeleteCourse      = errcode.NewRPCStatus(_courseBaseCode+2, "failed to delete "+_courseName)
	StatusDeleteByIDsCourse = errcode.NewRPCStatus(_courseBaseCode+3, "failed to delete by batch ids "+_courseName)
	StatusUpdateCourse      = errcode.NewRPCStatus(_courseBaseCode+4, "failed to update "+_courseName)
	StatusGetCourse         = errcode.NewRPCStatus(_courseBaseCode+5, "failed to get "+_courseName+" details")
	StatusListByIDsCourse   = errcode.NewRPCStatus(_courseBaseCode+6, "failed to list by batch ids "+_courseName)
	StatusListCourse        = errcode.NewRPCStatus(_courseBaseCode+7, "failed to list of "+_courseName)
	// add +1 to the previous error code
)
