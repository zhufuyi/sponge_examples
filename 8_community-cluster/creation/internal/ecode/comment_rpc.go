// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// commentService rpc service level error code
var (
	_commentServiceNO       = 4 // number range 1~100, if there is the same number, trigger panic.
	_commentServiceName     = "commentService"
	_commentServiceBaseCode = errcode.RCode(_commentServiceNO)

	StatusCreateCommentService     = errcode.NewRPCStatus(_commentServiceBaseCode+1, "failed to Create "+_commentServiceName)
	StatusDeleteCommentService     = errcode.NewRPCStatus(_commentServiceBaseCode+2, "failed to Delete "+_commentServiceName)
	StatusUpdateByIDCommentService = errcode.NewRPCStatus(_commentServiceBaseCode+3, "failed to UpdateByID "+_commentServiceName)
	StatusReplyCommentService      = errcode.NewRPCStatus(_commentServiceBaseCode+4, "failed to Reply "+_commentServiceName)
	StatusGetByIDCommentService    = errcode.NewRPCStatus(_commentServiceBaseCode+5, "failed to GetByID "+_commentServiceName)
	StatusListByIDsCommentService  = errcode.NewRPCStatus(_commentServiceBaseCode+6, "failed to ListByIDs "+_commentServiceName)
	StatusListHotCommentService    = errcode.NewRPCStatus(_commentServiceBaseCode+7, "failed to ListHot "+_commentServiceName)
	StatusListLatestCommentService = errcode.NewRPCStatus(_commentServiceBaseCode+8, "failed to ListLatest "+_commentServiceName)
	StatusListReplyCommentService  = errcode.NewRPCStatus(_commentServiceBaseCode+9, "failed to ListReply "+_commentServiceName)
	StatusReCommentService         = errcode.NewRPCStatus(_commentServiceBaseCode+10, "不支持多次嵌套已回复的评论")
	// add +1 to the previous error code
)