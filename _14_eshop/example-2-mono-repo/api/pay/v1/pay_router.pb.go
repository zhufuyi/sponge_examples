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

type PayLogicer interface {
	Create(ctx context.Context, req *CreatePayRequest) (*CreatePayReply, error)
	CreateRevert(ctx context.Context, req *CreatePayRevertRequest) (*CreatePayRevertReply, error)
	DeleteByID(ctx context.Context, req *DeletePayByIDRequest) (*DeletePayByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdatePayByIDRequest) (*UpdatePayByIDReply, error)
	GetByID(ctx context.Context, req *GetPayByIDRequest) (*GetPayByIDReply, error)
	List(ctx context.Context, req *ListPayRequest) (*ListPayReply, error)
}

type PayOption func(*payOptions)

type payOptions struct {
	isFromRPC bool
	responser errcode.Responser
	zapLog    *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *payOptions) apply(opts ...PayOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithPayHTTPResponse() PayOption {
	return func(o *payOptions) {
		o.isFromRPC = false
	}
}

func WithPayRPCResponse() PayOption {
	return func(o *payOptions) {
		o.isFromRPC = true
	}
}

func WithPayResponser(responser errcode.Responser) PayOption {
	return func(o *payOptions) {
		o.responser = responser
	}
}

func WithPayLogger(zapLog *zap.Logger) PayOption {
	return func(o *payOptions) {
		o.zapLog = zapLog
	}
}

func WithPayErrorToHTTPCode(e ...*errcode.Error) PayOption {
	return func(o *payOptions) {
		o.httpErrors = e
	}
}

func WithPayRPCStatusToHTTPCode(s ...*errcode.RPCStatus) PayOption {
	return func(o *payOptions) {
		o.rpcStatus = s
	}
}

func WithPayWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) PayOption {
	return func(o *payOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterPayRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic PayLogicer,
	opts ...PayOption) {

	o := &payOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog,_ = zap.NewProduction()
	}

	r := &payRouter {
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

type payRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                PayLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *payRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/pay/create", r.withMiddleware("POST", "/api/v1/pay/create", r.Create_0)...)
	r.iRouter.Handle("POST", "/api/v1/pay/createRevert", r.withMiddleware("POST", "/api/v1/pay/createRevert", r.CreateRevert_0)...)
	r.iRouter.Handle("DELETE", "/api/v1/pay/:id", r.withMiddleware("DELETE", "/api/v1/pay/:id", r.DeleteByID_0)...)
	r.iRouter.Handle("PUT", "/api/v1/pay/:id", r.withMiddleware("PUT", "/api/v1/pay/:id", r.UpdateByID_0)...)
	r.iRouter.Handle("GET", "/api/v1/pay/:id", r.withMiddleware("GET", "/api/v1/pay/:id", r.GetByID_0)...)
	r.iRouter.Handle("POST", "/api/v1/pay/list", r.withMiddleware("POST", "/api/v1/pay/list", r.List_0)...)

}

func (r *payRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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


func (r *payRouter) Create_0 (c *gin.Context) {
	req := &CreatePayRequest{}
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

func (r *payRouter) CreateRevert_0 (c *gin.Context) {
	req := &CreatePayRevertRequest{}
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

func (r *payRouter) DeleteByID_0 (c *gin.Context) {
	req := &DeletePayByIDRequest{}
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

func (r *payRouter) UpdateByID_0 (c *gin.Context) {
	req := &UpdatePayByIDRequest{}
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

func (r *payRouter) GetByID_0 (c *gin.Context) {
	req := &GetPayByIDRequest{}
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

func (r *payRouter) List_0 (c *gin.Context) {
	req := &ListPayRequest{}
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
