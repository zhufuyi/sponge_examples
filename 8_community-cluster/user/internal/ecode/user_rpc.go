package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// userService rpc service level error code
var (
	_userServiceNO       = 1 // number range 1~100, if there is the same number, trigger panic.
	_userServiceName     = "userService"
	_userServiceBaseCode = errcode.RCode(_userServiceNO)

	StatusRegisterUserService       = errcode.NewRPCStatus(_userServiceBaseCode+1, "failed to Register "+_userServiceName)
	StatusLoginUserService          = errcode.NewRPCStatus(_userServiceBaseCode+2, "failed to Login "+_userServiceName)
	StatusLogoutUserService         = errcode.NewRPCStatus(_userServiceBaseCode+3, "failed to Logout "+_userServiceName)
	StatusCreateUserService         = errcode.NewRPCStatus(_userServiceBaseCode+4, "failed to Create "+_userServiceName)
	StatusDeleteByIDUserService     = errcode.NewRPCStatus(_userServiceBaseCode+5, "failed to DeleteByID "+_userServiceName)
	StatusUpdateByIDUserService     = errcode.NewRPCStatus(_userServiceBaseCode+6, "failed to UpdateByID "+_userServiceName)
	StatusGetByIDUserService        = errcode.NewRPCStatus(_userServiceBaseCode+7, "failed to GetByID "+_userServiceName)
	StatusListUserService           = errcode.NewRPCStatus(_userServiceBaseCode+8, "failed to List "+_userServiceName)
	StatusUpdatePasswordUserService = errcode.NewRPCStatus(_userServiceBaseCode+9, "failed to UpdatePassword "+_userServiceName)
	StatusVerigyCodeUserService     = errcode.NewRPCStatus(_userServiceBaseCode+10, "验证码错误")
	StatusPasswordUserService       = errcode.NewRPCStatus(_userServiceBaseCode+11, "账号或密码错误")
	StatusTokenUserService          = errcode.NewRPCStatus(_userServiceBaseCode+12, "token验证失败")
	StatusNoAccountUserService      = errcode.NewRPCStatus(_userServiceBaseCode+13, "未注册账号")
	StatusAlreadyLoginUserService   = errcode.NewRPCStatus(_userServiceBaseCode+14, "请勿重复登录")
	StatusNotLoginUserService       = errcode.NewRPCStatus(_userServiceBaseCode+15, "未登录，禁止操作")
	// add +1 to the previous error code
)
