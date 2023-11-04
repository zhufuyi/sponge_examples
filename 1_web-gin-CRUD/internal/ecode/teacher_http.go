package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teacher business-level http error codes.
// the teacherNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	teacherNO       = 41
	teacherName     = "teacher"
	teacherBaseCode = errcode.HCode(teacherNO)

	ErrCreateTeacher         = errcode.NewError(teacherBaseCode+1, "failed to create "+teacherName)
	ErrDeleteByIDTeacher     = errcode.NewError(teacherBaseCode+2, "failed to delete "+teacherName)
	ErrDeleteByIDsTeacher    = errcode.NewError(teacherBaseCode+3, "failed to delete by batch ids "+teacherName)
	ErrUpdateByIDTeacher     = errcode.NewError(teacherBaseCode+4, "failed to update "+teacherName)
	ErrGetByIDTeacher        = errcode.NewError(teacherBaseCode+5, "failed to get "+teacherName+" details")
	ErrGetByConditionTeacher = errcode.NewError(teacherBaseCode+6, "failed to get "+teacherName+" details by conditions")
	ErrListByIDsTeacher      = errcode.NewError(teacherBaseCode+7, "failed to list by batch ids "+teacherName)
	ErrListTeacher           = errcode.NewError(teacherBaseCode+8, "failed to list of "+teacherName)
	// error codes are globally unique, adding 1 to the previous error code
)
