// Code generated by https://github.com/zhufuyi/sponge, DO NOT EDIT.

package v1

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/zhufuyi/sponge/pkg/errcode"
	"github.com/zhufuyi/sponge/pkg/gin/middleware"
)

type UserLogicer interface {
	Create(ctx context.Context, req *CreateUserRequest) (*CreateUserReply, error)
	DeleteByID(ctx context.Context, req *DeleteUserByIDRequest) (*DeleteUserByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdateUserByIDRequest) (*UpdateUserByIDReply, error)
	GetByID(ctx context.Context, req *GetUserByIDRequest) (*GetUserByIDReply, error)
	List(ctx context.Context, req *ListUserRequest) (*ListUserReply, error)
}

type UserOption func(*userOptions)

type userOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *userOptions) apply(opts ...UserOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithUserHTTPResponse() UserOption {
	return func(o *userOptions) {
		o.isFromRPC = false
	}
}

func WithUserRPCResponse() UserOption {
	return func(o *userOptions) {
		o.isFromRPC = true
	}
}

func WithUserResponser(responser errcode.Responser) UserOption {
	return func(o *userOptions) {
		o.responser = responser
	}
}

func WithUserLogger(zapLog *zap.Logger) UserOption {
	return func(o *userOptions) {
		o.zapLog = zapLog
	}
}

func WithUserErrorToHTTPCode(e ...*errcode.Error) UserOption {
	return func(o *userOptions) {
		o.httpErrors = e
	}
}

func WithUserRPCStatusToHTTPCode(s ...*errcode.RPCStatus) UserOption {
	return func(o *userOptions) {
		o.rpcStatus = s
	}
}

func WithUserWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) UserOption {
	return func(o *userOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterUserRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic UserLogicer,
	opts ...UserOption) {

	o := &userOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &userRouter{
		iRouter:               iRouter,
		groupPathMiddlewares:  groupPathMiddlewares,
		singlePathMiddlewares: singlePathMiddlewares,
		iLogic:                iLogic,
		iResponse:             o.responser,
		zapLog:                o.zapLog,
		wrapCtxFn:             o.wrapCtxFn,
	}
	r.register()
}

type userRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                UserLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *userRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/user", r.withMiddleware("POST", "/api/v1/user", r.Create_0)...)
	r.iRouter.Handle("DELETE", "/api/v1/user/:id", r.withMiddleware("DELETE", "/api/v1/user/:id", r.DeleteByID_0)...)
	r.iRouter.Handle("PUT", "/api/v1/user/:id", r.withMiddleware("PUT", "/api/v1/user/:id", r.UpdateByID_0)...)
	r.iRouter.Handle("GET", "/api/v1/user/:id", r.withMiddleware("GET", "/api/v1/user/:id", r.GetByID_0)...)
	r.iRouter.Handle("POST", "/api/v1/user/list", r.withMiddleware("POST", "/api/v1/user/list", r.List_0)...)

}

func (r *userRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
	handlerFns := []gin.HandlerFunc{}

	// determine if a route group is hit or miss, left prefix rule
	for groupPath, fns := range r.groupPathMiddlewares {
		if groupPath == "" || groupPath == "/" {
			handlerFns = append(handlerFns, fns...)
			continue
		}
		size := len(groupPath)
		if len(path) < size {
			continue
		}
		if groupPath == path[:size] {
			handlerFns = append(handlerFns, fns...)
		}
	}

	// determine if a single route has been hit
	key := strings.ToUpper(method) + "->" + path
	if fns, ok := r.singlePathMiddlewares[key]; ok {
		handlerFns = append(handlerFns, fns...)
	}

	return append(handlerFns, fn)
}

func (r *userRouter) Create_0(c *gin.Context) {
	req := &CreateUserRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.Create(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userRouter) DeleteByID_0(c *gin.Context) {
	req := &DeleteUserByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindQuery(req); err != nil {
		r.zapLog.Warn("ShouldBindQuery error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.DeleteByID(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userRouter) UpdateByID_0(c *gin.Context) {
	req := &UpdateUserByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.UpdateByID(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userRouter) GetByID_0(c *gin.Context) {
	req := &GetUserByIDRequest{}
	var err error

	if err = c.ShouldBindUri(req); err != nil {
		r.zapLog.Warn("ShouldBindUri error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	if err = c.ShouldBindQuery(req); err != nil {
		r.zapLog.Warn("ShouldBindQuery error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.GetByID(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *userRouter) List_0(c *gin.Context) {
	req := &ListUserRequest{}
	var err error

	if err = c.ShouldBindJSON(req); err != nil {
		r.zapLog.Warn("ShouldBindJSON error", zap.Error(err), middleware.GCtxRequestIDField(c))
		r.iResponse.ParamError(c, err)
		return
	}

	var ctx context.Context
	if r.wrapCtxFn != nil {
		ctx = r.wrapCtxFn(c)
	} else {
		ctx = middleware.WrapCtx(c)
	}

	out, err := r.iLogic.List(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
