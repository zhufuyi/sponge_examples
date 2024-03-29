// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// commentService business-level http error codes.
// the commentServiceNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	commentServiceNO       = 2
	commentServiceName     = "commentService"
	commentServiceBaseCode = errcode.HCode(commentServiceNO)

	ErrCreateCommentService     = errcode.NewError(commentServiceBaseCode+1, "failed to Create "+commentServiceName)
	ErrDeleteByIDCommentService = errcode.NewError(commentServiceBaseCode+2, "failed to DeleteByID "+commentServiceName)
	ErrUpdateByIDCommentService = errcode.NewError(commentServiceBaseCode+3, "failed to UpdateByID "+commentServiceName)
	ErrReplyCommentService      = errcode.NewError(commentServiceBaseCode+4, "failed to Reply "+commentServiceName)
	ErrGetByIDCommentService    = errcode.NewError(commentServiceBaseCode+5, "failed to GetByID "+commentServiceName)
	ErrListByIDsCommentService  = errcode.NewError(commentServiceBaseCode+6, "failed to ListByIDs "+commentServiceName)
	ErrListLatestCommentService = errcode.NewError(commentServiceBaseCode+7, "failed to ListLatest "+commentServiceName)
	ErrListHotCommentService    = errcode.NewError(commentServiceBaseCode+8, "failed to ListHot "+commentServiceName)
	ErrListReplyCommentService  = errcode.NewError(commentServiceBaseCode+9, "failed to ListReply "+commentServiceName)
	ErrReCommentService         = errcode.NewError(commentServiceBaseCode+10, "不支持多次嵌套已回复的评论")
	// error codes are globally unique, adding 1 to the previous error code
)
