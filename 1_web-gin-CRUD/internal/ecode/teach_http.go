package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// teach business-level http error codes.
// the teachNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	teachNO       = 36
	teachName     = "teach"
	teachBaseCode = errcode.HCode(teachNO)

	ErrCreateTeach         = errcode.NewError(teachBaseCode+1, "failed to create "+teachName)
	ErrDeleteByIDTeach     = errcode.NewError(teachBaseCode+2, "failed to delete "+teachName)
	ErrDeleteByIDsTeach    = errcode.NewError(teachBaseCode+3, "failed to delete by batch ids "+teachName)
	ErrUpdateByIDTeach     = errcode.NewError(teachBaseCode+4, "failed to update "+teachName)
	ErrGetByIDTeach        = errcode.NewError(teachBaseCode+5, "failed to get "+teachName+" details")
	ErrGetByConditionTeach = errcode.NewError(teachBaseCode+6, "failed to get "+teachName+" details by conditions")
	ErrListByIDsTeach      = errcode.NewError(teachBaseCode+7, "failed to list by batch ids "+teachName)
	ErrListTeach           = errcode.NewError(teachBaseCode+8, "failed to list of "+teachName)
	// error codes are globally unique, adding 1 to the previous error code
)
