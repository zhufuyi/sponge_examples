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

type OrderLogicer interface {
	Submit(ctx context.Context, req *SubmitOrderRequest) (*SubmitOrderReply, error)
	SendSubmitOrderNotify(ctx context.Context, req *SendSubmitOrderNotifyRequest) (*SendSubmitOrderNotifyReply, error)
	Create(ctx context.Context, req *CreateOrderRequest) (*CreateOrderReply, error)
	CreateRevert(ctx context.Context, req *CreateOrderRevertRequest) (*CreateOrderRevertReply, error)
	DeleteByID(ctx context.Context, req *DeleteOrderByIDRequest) (*DeleteOrderByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdateOrderByIDRequest) (*UpdateOrderByIDReply, error)
	GetByID(ctx context.Context, req *GetOrderByIDRequest) (*GetOrderByIDReply, error)
	List(ctx context.Context, req *ListOrderRequest) (*ListOrderReply, error)
}

type OrderOption func(*orderOptions)

type orderOptions struct {
	isFromRPC bool
	responser errcode.Responser
	zapLog    *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *orderOptions) apply(opts ...OrderOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithOrderHTTPResponse() OrderOption {
	return func(o *orderOptions) {
		o.isFromRPC = false
	}
}

func WithOrderRPCResponse() OrderOption {
	return func(o *orderOptions) {
		o.isFromRPC = true
	}
}

func WithOrderResponser(responser errcode.Responser) OrderOption {
	return func(o *orderOptions) {
		o.responser = responser
	}
}

func WithOrderLogger(zapLog *zap.Logger) OrderOption {
	return func(o *orderOptions) {
		o.zapLog = zapLog
	}
}

func WithOrderErrorToHTTPCode(e ...*errcode.Error) OrderOption {
	return func(o *orderOptions) {
		o.httpErrors = e
	}
}

func WithOrderRPCStatusToHTTPCode(s ...*errcode.RPCStatus) OrderOption {
	return func(o *orderOptions) {
		o.rpcStatus = s
	}
}

func WithOrderWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) OrderOption {
	return func(o *orderOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterOrderRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic OrderLogicer,
	opts ...OrderOption) {

	o := &orderOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog,_ = zap.NewProduction()
	}

	r := &orderRouter {
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

type orderRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                OrderLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *orderRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/order/submit", r.withMiddleware("POST", "/api/v1/order/submit", r.Submit_0)...)
	r.iRouter.Handle("POST", "/api/v1/order/sendSubmitNotify", r.withMiddleware("POST", "/api/v1/order/sendSubmitNotify", r.SendSubmitOrderNotify_0)...)
	r.iRouter.Handle("POST", "/api/v1/order/create", r.withMiddleware("POST", "/api/v1/order/create", r.Create_0)...)
	r.iRouter.Handle("POST", "/api/v1/order/createRevert", r.withMiddleware("POST", "/api/v1/order/createRevert", r.CreateRevert_0)...)
	r.iRouter.Handle("DELETE", "/api/v1/order/:id", r.withMiddleware("DELETE", "/api/v1/order/:id", r.DeleteByID_0)...)
	r.iRouter.Handle("PUT", "/api/v1/order/:id", r.withMiddleware("PUT", "/api/v1/order/:id", r.UpdateByID_0)...)
	r.iRouter.Handle("GET", "/api/v1/order/:id", r.withMiddleware("GET", "/api/v1/order/:id", r.GetByID_0)...)
	r.iRouter.Handle("POST", "/api/v1/order/list", r.withMiddleware("POST", "/api/v1/order/list", r.List_0)...)

}

func (r *orderRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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


func (r *orderRouter) Submit_0 (c *gin.Context) {
	req := &SubmitOrderRequest{}
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


	out, err := r.iLogic.Submit(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *orderRouter) SendSubmitOrderNotify_0 (c *gin.Context) {
	req := &SendSubmitOrderNotifyRequest{}
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


	out, err := r.iLogic.SendSubmitOrderNotify(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *orderRouter) Create_0 (c *gin.Context) {
	req := &CreateOrderRequest{}
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

func (r *orderRouter) CreateRevert_0 (c *gin.Context) {
	req := &CreateOrderRevertRequest{}
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


	out, err := r.iLogic.CreateRevert(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *orderRouter) DeleteByID_0 (c *gin.Context) {
	req := &DeleteOrderByIDRequest{}
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

func (r *orderRouter) UpdateByID_0 (c *gin.Context) {
	req := &UpdateOrderByIDRequest{}
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

func (r *orderRouter) GetByID_0 (c *gin.Context) {
	req := &GetOrderByIDRequest{}
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

func (r *orderRouter) List_0 (c *gin.Context) {
	req := &ListOrderRequest{}
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

