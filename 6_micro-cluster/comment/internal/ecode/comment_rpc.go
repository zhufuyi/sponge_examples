// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// comment rpc service level error code
var (
	_commentNO       = 78 // number range 1~100, if there is the same number, trigger panic.
	_commentName     = "comment"
	_commentBaseCode = errcode.HCode(_commentNO)

	StatusListByProductIDComment = errcode.NewError(_commentBaseCode+1, "failed to ListByProductID "+_commentName)
	// add +1 to the previous error code
)
