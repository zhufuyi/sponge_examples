// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// postService rpc service level error code
var (
	_postServiceNO       = 9 // number range 1~100, if there is the same number, trigger panic.
	_postServiceName     = "postService"
	_postServiceBaseCode = errcode.RCode(_postServiceNO)

	StatusCreatePostService         = errcode.NewRPCStatus(_postServiceBaseCode+1, "failed to Create "+_postServiceName)
	StatusUpdateContentPostService  = errcode.NewRPCStatus(_postServiceBaseCode+2, "failed to UpdateByID "+_postServiceName)
	StatusDeletePostService         = errcode.NewRPCStatus(_postServiceBaseCode+3, "failed to Delete "+_postServiceName)
	StatusGetByIDPostService        = errcode.NewRPCStatus(_postServiceBaseCode+4, "failed to GetByID "+_postServiceName)
	StatusListByUserIDPostService   = errcode.NewRPCStatus(_postServiceBaseCode+5, "failed to ListByUserID "+_postServiceName)
	StatusListLatestPostService     = errcode.NewRPCStatus(_postServiceBaseCode+6, "failed to ListLatest "+_postServiceName)
	StatusListHotPostService        = errcode.NewRPCStatus(_postServiceBaseCode+7, "failed to ListHot "+_postServiceName)
	StatusIncrViewCountPostService  = errcode.NewRPCStatus(_postServiceBaseCode+8, "failed to IncrViewCount "+_postServiceName)
	StatusIncrShareCountPostService = errcode.NewRPCStatus(_postServiceBaseCode+9, "failed to IncrShareCount "+_postServiceName)
	// add +1 to the previous error code
)
