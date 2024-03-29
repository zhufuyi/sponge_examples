// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// likeService business-level http error codes.
// the likeServiceNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	likeServiceNO       = 3
	likeServiceName     = "likeService"
	likeServiceBaseCode = errcode.HCode(likeServiceNO)

	ErrCreateLikeService      = errcode.NewError(likeServiceBaseCode+1, "failed to Create "+likeServiceName)
	ErrDeleteLikeService      = errcode.NewError(likeServiceBaseCode+2, "failed to Delete "+likeServiceName)
	ErrListPostLikeService    = errcode.NewError(likeServiceBaseCode+3, "failed to ListPost "+likeServiceName)
	ErrListCommentLikeService = errcode.NewError(likeServiceBaseCode+4, "failed to ListComment "+likeServiceName)
	// error codes are globally unique, adding 1 to the previous error code
)
