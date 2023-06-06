// nolint

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teacher http service level error code
// each resource name corresponds to a unique number (http type), the number range is 1~100, if there is the same number, trigger panic
var (
	teacherNO       = 1
	teacherName     = "teacher"
	teacherBaseCode = errcode.HCode(teacherNO)

	ErrCreateTeacher      = errcode.NewError(teacherBaseCode+1, "failed to create "+teacherName)
	ErrDeleteTeacher      = errcode.NewError(teacherBaseCode+2, "failed to delete "+teacherName)
	ErrDeleteByIDsTeacher = errcode.NewError(teacherBaseCode+3, "failed to delete by batch ids "+teacherName)
	ErrUpdateTeacher      = errcode.NewError(teacherBaseCode+4, "failed to update "+teacherName)
	ErrGetTeacher         = errcode.NewError(teacherBaseCode+5, "failed to get "+teacherName+" details")
	ErrListByIDsTeacher   = errcode.NewError(teacherBaseCode+6, "failed to list by batch ids "+teacherName)
	ErrListTeacher        = errcode.NewError(teacherBaseCode+7, "failed to list of "+teacherName)
	// for each error code added, add +1 to the previous error code
)
