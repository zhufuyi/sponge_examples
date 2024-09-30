// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// user business-level rpc error codes.
// the _userNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_userNO       = 1
	_userName     = "user"
	_userBaseCode = errcode.RCode(_userNO)

	StatusRegisterUser       = errcode.NewRPCStatus(_userBaseCode+1, "failed to Register "+_userName)
	StatusLoginUser          = errcode.NewRPCStatus(_userBaseCode+2, "failed to Login "+_userName)
	StatusLogoutUser         = errcode.NewRPCStatus(_userBaseCode+3, "failed to Logout "+_userName)
	StatusChangePasswordUser = errcode.NewRPCStatus(_userBaseCode+4, "failed to ChangePassword "+_userName)
	// error codes are globally unique, adding 1 to the previous error code
)