// nolint

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// course http service level error code
// each resource name corresponds to a unique number (http type), the number range is 1~100, if there is the same number, trigger panic
var (
	courseNO       = 2
	courseName     = "course"
	courseBaseCode = errcode.HCode(courseNO)

	ErrCreateCourse      = errcode.NewError(courseBaseCode+1, "failed to create "+courseName)
	ErrDeleteCourse      = errcode.NewError(courseBaseCode+2, "failed to delete "+courseName)
	ErrDeleteByIDsCourse = errcode.NewError(courseBaseCode+3, "failed to delete by batch ids "+courseName)
	ErrUpdateCourse      = errcode.NewError(courseBaseCode+4, "failed to update "+courseName)
	ErrGetCourse         = errcode.NewError(courseBaseCode+5, "failed to get "+courseName+" details")
	ErrListByIDsCourse   = errcode.NewError(courseBaseCode+6, "failed to list by batch ids "+courseName)
	ErrListCourse        = errcode.NewError(courseBaseCode+7, "failed to list of "+courseName)
	// for each error code added, add +1 to the previous error code
)
