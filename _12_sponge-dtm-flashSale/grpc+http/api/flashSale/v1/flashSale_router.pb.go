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

type FlashSaleLogicer interface {
	SetProductStock(ctx context.Context, req *SetProductStockRequest) (*SetProductStockReply, error)
	FlashSale(ctx context.Context, req *FlashSaleRequest) (*FlashSaleReply, error)
}

type FlashSaleOption func(*flashSaleOptions)

type flashSaleOptions struct {
	isFromRPC  bool
	responser  errcode.Responser
	zapLog     *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *flashSaleOptions) apply(opts ...FlashSaleOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithFlashSaleHTTPResponse() FlashSaleOption {
	return func(o *flashSaleOptions) {
		o.isFromRPC = false
	}
}

func WithFlashSaleRPCResponse() FlashSaleOption {
	return func(o *flashSaleOptions) {
		o.isFromRPC = true
	}
}

func WithFlashSaleResponser(responser errcode.Responser) FlashSaleOption {
	return func(o *flashSaleOptions) {
		o.responser = responser
	}
}

func WithFlashSaleLogger(zapLog *zap.Logger) FlashSaleOption {
	return func(o *flashSaleOptions) {
		o.zapLog = zapLog
	}
}

func WithFlashSaleErrorToHTTPCode(e ...*errcode.Error) FlashSaleOption {
	return func(o *flashSaleOptions) {
		o.httpErrors = e
	}
}

func WithFlashSaleRPCStatusToHTTPCode(s ...*errcode.RPCStatus) FlashSaleOption {
	return func(o *flashSaleOptions) {
		o.rpcStatus = s
	}
}

func WithFlashSaleWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) FlashSaleOption {
	return func(o *flashSaleOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterFlashSaleRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic FlashSaleLogicer,
	opts ...FlashSaleOption) {

	o := &flashSaleOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog, _ = zap.NewProduction()
	}

	r := &flashSaleRouter{
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

type flashSaleRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                FlashSaleLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *flashSaleRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/flashSale/setProductStock", r.withMiddleware("POST", "/api/v1/flashSale/setProductStock", r.SetProductStock_0)...)
	r.iRouter.Handle("POST", "/api/v1/flashSale", r.withMiddleware("POST", "/api/v1/flashSale", r.FlashSale_0)...)

}

func (r *flashSaleRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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

func (r *flashSaleRouter) SetProductStock_0(c *gin.Context) {
	req := &SetProductStockRequest{}
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

	out, err := r.iLogic.SetProductStock(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *flashSaleRouter) FlashSale_0(c *gin.Context) {
	req := &FlashSaleRequest{}
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

	out, err := r.iLogic.FlashSale(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}
