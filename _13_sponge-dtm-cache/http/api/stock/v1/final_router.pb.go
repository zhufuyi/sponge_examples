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

type FinalLogicer interface {
	Update(ctx context.Context, req *UpdateFinalRequest) (*UpdateFinalRequestReply, error)
	Query(ctx context.Context, req *QueryFinalRequest) (*QueryFinalReply, error)
}

type FinalOption func(*finalOptions)

type finalOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *finalOptions) apply(opts ...FinalOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithFinalHTTPResponse() FinalOption {
	return func(o *finalOptions) {
		o.isFromRPC = false
	}
}

func WithFinalRPCResponse() FinalOption {
	return func(o *finalOptions) {
		o.isFromRPC = true
	}
}

func WithFinalResponser(responser errcode.Responser) FinalOption {
	return func(o *finalOptions) {
		o.responser = responser
	}
}

func WithFinalLogger(zapLog *zap.Logger) FinalOption {
	return func(o *finalOptions) {
		o.zapLog = zapLog
	}
}

func WithFinalErrorToHTTPCode(e ...*errcode.Error) FinalOption {
	return func(o *finalOptions) {
		o.httpErrors = e
	}
}

func WithFinalRPCStatusToHTTPCode(s ...*errcode.RPCStatus) FinalOption {
	return func(o *finalOptions) {
		o.rpcStatus = s
	}
}

func WithFinalWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) FinalOption {
	return func(o *finalOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterFinalRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic FinalLogicer,
	opts ...FinalOption) {

	o := &finalOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &finalRouter{
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

type finalRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                FinalLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *finalRouter) register() {
	r.iRouter.Handle("PUT", "/api/v1/stock/:id/final", r.withMiddleware("PUT", "/api/v1/stock/:id/final", r.Update_4)...)
	r.iRouter.Handle("GET", "/api/v1/stock/:id/final", r.withMiddleware("GET", "/api/v1/stock/:id/final", r.Query_4)...)

}

func (r *finalRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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

func (r *finalRouter) Update_4(c *gin.Context) {
	req := &UpdateFinalRequest{}
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

	out, err := r.iLogic.Update(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *finalRouter) Query_4(c *gin.Context) {
	req := &QueryFinalRequest{}
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

	out, err := r.iLogic.Query(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
