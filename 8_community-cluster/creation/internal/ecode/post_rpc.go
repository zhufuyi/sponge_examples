// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// postService business-level rpc error codes.
// the _postServiceNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_postServiceNO       = 43
	_postServiceName     = "postService"
	_postServiceBaseCode = errcode.RCode(_postServiceNO)

	StatusCreatePostService         = errcode.NewRPCStatus(_postServiceBaseCode+1, "failed to Create "+_postServiceName)
	StatusUpdateContentPostService  = errcode.NewRPCStatus(_postServiceBaseCode+2, "failed to UpdateContent "+_postServiceName)
	StatusDeletePostService         = errcode.NewRPCStatus(_postServiceBaseCode+3, "failed to Delete "+_postServiceName)
	StatusGetByIDPostService        = errcode.NewRPCStatus(_postServiceBaseCode+4, "failed to GetByID "+_postServiceName)
	StatusListByIDsPostService      = errcode.NewRPCStatus(_postServiceBaseCode+5, "failed to ListByIDs "+_postServiceName)
	StatusListByUserIDPostService   = errcode.NewRPCStatus(_postServiceBaseCode+6, "failed to ListByUserID "+_postServiceName)
	StatusListLatestPostService     = errcode.NewRPCStatus(_postServiceBaseCode+7, "failed to ListLatest "+_postServiceName)
	StatusListHotPostService        = errcode.NewRPCStatus(_postServiceBaseCode+8, "failed to ListHot "+_postServiceName)
	StatusIncrViewCountPostService  = errcode.NewRPCStatus(_postServiceBaseCode+9, "failed to IncrViewCount "+_postServiceName)
	StatusIncrShareCountPostService = errcode.NewRPCStatus(_postServiceBaseCode+10, "failed to IncrShareCount "+_postServiceName)
	StatusPostTypePostService       = errcode.NewRPCStatus(_postServiceBaseCode+11, "发布类型参数错误")
	StatusPostType2PostService      = errcode.NewRPCStatus(_postServiceBaseCode+12, "发布类型中图片和视频只能选其中一个")
	StatusVideoParamPostService     = errcode.NewRPCStatus(_postServiceBaseCode+13, "发布视频参数错误")
	// error codes are globally unique, adding 1 to the previous error code
)
