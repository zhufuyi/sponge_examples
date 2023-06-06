// nolint

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teach http service level error code
// each resource name corresponds to a unique number (http type), the number range is 1~100, if there is the same number, trigger panic
var (
	teachNO       = 3
	teachName     = "teach"
	teachBaseCode = errcode.HCode(teachNO)

	ErrCreateTeach      = errcode.NewError(teachBaseCode+1, "failed to create "+teachName)
	ErrDeleteTeach      = errcode.NewError(teachBaseCode+2, "failed to delete "+teachName)
	ErrDeleteByIDsTeach = errcode.NewError(teachBaseCode+3, "failed to delete by batch ids "+teachName)
	ErrUpdateTeach      = errcode.NewError(teachBaseCode+4, "failed to update "+teachName)
	ErrGetTeach         = errcode.NewError(teachBaseCode+5, "failed to get "+teachName+" details")
	ErrListByIDsTeach   = errcode.NewError(teachBaseCode+6, "failed to list by batch ids "+teachName)
	ErrListTeach        = errcode.NewError(teachBaseCode+7, "failed to list of "+teachName)
	// for each error code added, add +1 to the previous error code
)
