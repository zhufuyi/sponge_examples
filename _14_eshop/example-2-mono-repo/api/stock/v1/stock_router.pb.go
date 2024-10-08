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

type StockLogicer interface {
	StockDeduct(ctx context.Context, req *StockDeductRequest) (*StockDeductReply, error)
	StockDeductRevert(ctx context.Context, req *StockDeductRevertRequest) (*StockDeductRevertReply, error)
	SetFlashSaleStock(ctx context.Context, req *SetFlashSaleStockRequest) (*SetFlashSaleStockReply, error)
	Create(ctx context.Context, req *CreateStockRequest) (*CreateStockReply, error)
	DeleteByID(ctx context.Context, req *DeleteStockByIDRequest) (*DeleteStockByIDReply, error)
	UpdateByID(ctx context.Context, req *UpdateStockByIDRequest) (*UpdateStockByIDReply, error)
	GetByID(ctx context.Context, req *GetStockByIDRequest) (*GetStockByIDReply, error)
	List(ctx context.Context, req *ListStockRequest) (*ListStockReply, error)
}

type StockOption func(*stockOptions)

type stockOptions struct {
	isFromRPC bool
	responser errcode.Responser
	zapLog    *zap.Logger
	httpErrors []*errcode.Error
	rpcStatus  []*errcode.RPCStatus
	wrapCtxFn  func(c *gin.Context) context.Context
}

func (o *stockOptions) apply(opts ...StockOption) {
	for _, opt := range opts {
		opt(o)
	}
}

func WithStockHTTPResponse() StockOption {
	return func(o *stockOptions) {
		o.isFromRPC = false
	}
}

func WithStockRPCResponse() StockOption {
	return func(o *stockOptions) {
		o.isFromRPC = true
	}
}

func WithStockResponser(responser errcode.Responser) StockOption {
	return func(o *stockOptions) {
		o.responser = responser
	}
}

func WithStockLogger(zapLog *zap.Logger) StockOption {
	return func(o *stockOptions) {
		o.zapLog = zapLog
	}
}

func WithStockErrorToHTTPCode(e ...*errcode.Error) StockOption {
	return func(o *stockOptions) {
		o.httpErrors = e
	}
}

func WithStockRPCStatusToHTTPCode(s ...*errcode.RPCStatus) StockOption {
	return func(o *stockOptions) {
		o.rpcStatus = s
	}
}

func WithStockWrapCtx(wrapCtxFn func(c *gin.Context) context.Context) StockOption {
	return func(o *stockOptions) {
		o.wrapCtxFn = wrapCtxFn
	}
}

func RegisterStockRouter(
	iRouter gin.IRouter,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iLogic StockLogicer,
	opts ...StockOption) {

	o := &stockOptions{}
	o.apply(opts...)

	if o.responser == nil {
		o.responser = errcode.NewResponser(o.isFromRPC, o.httpErrors, o.rpcStatus)
	}
	if o.zapLog == nil {
		o.zapLog,_ = zap.NewProduction()
	}

	r := &stockRouter {
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

type stockRouter struct {
	iRouter               gin.IRouter
	groupPathMiddlewares  map[string][]gin.HandlerFunc
	singlePathMiddlewares map[string][]gin.HandlerFunc
	iLogic                StockLogicer
	iResponse             errcode.Responser
	zapLog                *zap.Logger
	wrapCtxFn             func(c *gin.Context) context.Context
}

func (r *stockRouter) register() {
	r.iRouter.Handle("POST", "/api/v1/stock/deduct", r.withMiddleware("POST", "/api/v1/stock/deduct", r.StockDeduct_0)...)
	r.iRouter.Handle("POST", "/api/v1/stock/deductRevert", r.withMiddleware("POST", "/api/v1/stock/deductRevert", r.StockDeductRevert_0)...)
	r.iRouter.Handle("POST", "/api/v1/stock/setFlashSale", r.withMiddleware("POST", "/api/v1/stock/setFlashSale", r.SetFlashSaleStock_0)...)
	r.iRouter.Handle("POST", "/api/v1/stock", r.withMiddleware("POST", "/api/v1/stock", r.Create_0)...)
	r.iRouter.Handle("DELETE", "/api/v1/stock/:id", r.withMiddleware("DELETE", "/api/v1/stock/:id", r.DeleteByID_0)...)
	r.iRouter.Handle("PUT", "/api/v1/stock/:id", r.withMiddleware("PUT", "/api/v1/stock/:id", r.UpdateByID_0)...)
	r.iRouter.Handle("GET", "/api/v1/stock/:id", r.withMiddleware("GET", "/api/v1/stock/:id", r.GetByID_0)...)
	r.iRouter.Handle("POST", "/api/v1/stock/list", r.withMiddleware("POST", "/api/v1/stock/list", r.List_0)...)

}

func (r *stockRouter) withMiddleware(method string, path string, fn gin.HandlerFunc) []gin.HandlerFunc {
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


func (r *stockRouter) StockDeduct_0 (c *gin.Context) {
	req := &StockDeductRequest{}
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


	out, err := r.iLogic.StockDeduct(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *stockRouter) StockDeductRevert_0 (c *gin.Context) {
	req := &StockDeductRevertRequest{}
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


	out, err := r.iLogic.StockDeductRevert(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *stockRouter) SetFlashSaleStock_0 (c *gin.Context) {
	req := &SetFlashSaleStockRequest{}
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


	out, err := r.iLogic.SetFlashSaleStock(ctx, req)
	if err != nil {
		if errors.Is(err, errcode.SkipResponse) {
			return
		}
		r.iResponse.Error(c, err)
		return
	}

	r.iResponse.Success(c, out)
}

func (r *stockRouter) Create_0 (c *gin.Context) {
	req := &CreateStockRequest{}
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

func (r *stockRouter) DeleteByID_0 (c *gin.Context) {
	req := &DeleteStockByIDRequest{}
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

func (r *stockRouter) UpdateByID_0 (c *gin.Context) {
	req := &UpdateStockByIDRequest{}
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

func (r *stockRouter) GetByID_0 (c *gin.Context) {
	req := &GetStockByIDRequest{}
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

func (r *stockRouter) List_0 (c *gin.Context) {
	req := &ListStockRequest{}
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

