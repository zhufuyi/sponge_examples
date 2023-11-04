package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// course business-level http error codes.
// the courseNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	courseNO       = 86
	courseName     = "course"
	courseBaseCode = errcode.HCode(courseNO)

	ErrCreateCourse         = errcode.NewError(courseBaseCode+1, "failed to create "+courseName)
	ErrDeleteByIDCourse     = errcode.NewError(courseBaseCode+2, "failed to delete "+courseName)
	ErrDeleteByIDsCourse    = errcode.NewError(courseBaseCode+3, "failed to delete by batch ids "+courseName)
	ErrUpdateByIDCourse     = errcode.NewError(courseBaseCode+4, "failed to update "+courseName)
	ErrGetByIDCourse        = errcode.NewError(courseBaseCode+5, "failed to get "+courseName+" details")
	ErrGetByConditionCourse = errcode.NewError(courseBaseCode+6, "failed to get "+courseName+" details by conditions")
	ErrListByIDsCourse      = errcode.NewError(courseBaseCode+7, "failed to list by batch ids "+courseName)
	ErrListCourse           = errcode.NewError(courseBaseCode+8, "failed to list of "+courseName)
	// error codes are globally unique, adding 1 to the previous error code
)
