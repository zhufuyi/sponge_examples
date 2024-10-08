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

type EShopGwLogicer interface {
	GetDetailsByProductID(ctx context.Context, req *GetDetailsByProductIDRequest) (*GetDetailsByProductIDReply, error)
}

type EShopGwOption func(*eShopGwOptions)

type eShopGwOptions struct {
	isFromRPC bool
	responser errcode.Responser
	zapLog    *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *eShopGwOptions) apply(opts ...EShopGwOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithEShopGwHTTPResponse() EShopGwOption {
	return func(o *eShopGwOptions) {
		o.isFromRPC = false
	}
}

func WithEShopGwRPCResponse() EShopGwOption {
	return func(o *eShopGwOptions) {
		o.isFromRPC = true
	}
}

func WithEShopGwResponser(responser errcode.Responser) EShopGwOption {
	return func(o *eShopGwOptions) {
		o.responser = responser
	}
}

func WithEShopGwLogger(zapLog *zap.Logger) EShopGwOption {
	return func(o *eShopGwOptions) {
		o.zapLog = zapLog
	}
}

func WithEShopGwErrorToHTTPCode(e ...*errcode.Error) EShopGwOption {
	return func(o *eShopGwOptions) {
		o.httpErrors = e
	}
}

func WithEShopGwRPCStatusToHTTPCode(s ...*errcode.RPCStatus) EShopGwOption {
	return func(o *eShopGwOptions) {
		o.rpcStatus = s
	}
}

func WithEShopGwWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) EShopGwOption {
	return func(o *eShopGwOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterEShopGwRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic EShopGwLogicer,
	opts ...EShopGwOption) {

	o := &eShopGwOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog,_ = zap.NewProduction()
	}

	r := &eShopGwRouter {
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

type eShopGwRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                EShopGwLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *eShopGwRouter) register() {
	r.iRouter.Handle("GET", "/api/v1/detail", r.withMiddleware("GET", "/api/v1/detail", r.GetDetailsByProductID_0)...)

}

func (r *eShopGwRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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


func (r *eShopGwRouter) GetDetailsByProductID_0 (c *gin.Context) {
	req := &GetDetailsByProductIDRequest{}
	var err error



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


	out, err := r.iLogic.GetDetailsByProductID(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

