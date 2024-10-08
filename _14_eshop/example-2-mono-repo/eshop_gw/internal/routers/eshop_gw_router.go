// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"

	eshop_gwV1 "eshop/api/eshop_gw/v1"
	"eshop/eshop_gw/internal/service"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		eshopMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			eshopRouter(r, groupPathMiddlewares, singlePathMiddlewares, service.NewEshopClient())
		})
}

func eshopRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService eshop_gwV1.EshopLogicer) {
	ctxFn := func(c *gin.Context) context.Context {
		md := metadata.New(map[string]string{
			// set metadata to be passed from http to rpc
			middleware.ContextRequestIDKey: middleware.GCtxRequestID(c), // request_id
			//middleware.HeaderAuthorizationKey: c.GetHeader(middleware.HeaderAuthorizationKey),  // authorization
		})
		return metadata.NewOutgoingContext(c.Request.Context(), md)
	}

	eshop_gwV1.RegisterEshopRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		eshop_gwV1.WithEshopLogger(logger.Get()),
		eshop_gwV1.WithEshopRPCResponse(),
		eshop_gwV1.WithEshopWrapCtx(ctxFn),
		eshop_gwV1.WithEshopRPCStatusToHTTPCode(
		// Set some error codes to standard http return codes,
		// by default there is already ecode.StatusInternalServerError and ecode.StatusServiceUnavailable
		// example:
		// 	ecode.StatusUnimplemented, ecode.StatusAborted,
		),
	)
}

// you can set the middleware of a route group, or set the middleware of a single route,
// or you can mix them, pay attention to the duplication of middleware when mixing them,
// it is recommended to set the middleware of a single route in preference
func eshopMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/eshop/:id  will take effect
	// c.setGroupPath("/api/v1/eshop", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/user", middleware.Auth())    // CreateUser create user
	//c.setSinglePath("GET", "/api/v1/user/list", middleware.Auth())    // ListUser list of user by query parameters
	//c.setSinglePath("POST", "/api/v1/product", middleware.Auth())    // CreateProduct create product
	//c.setSinglePath("GET", "/api/v1/product/list", middleware.Auth())    // ListProduct list of product by query parameters
	//c.setSinglePath("POST", "/api/v1/order/submit", middleware.Auth())    // SubmitOrder 提交订单
	//c.setSinglePath("POST", "/api/v1/order/sendSubmitNotify", middleware.Auth())    // SendSubmitOrderNotify 发送提交订单通知
	//c.setSinglePath("GET", "/api/v1/order/list", middleware.Auth())    // ListOrder list of order by query parameters
	//c.setSinglePath("POST", "/api/v1/stock", middleware.Auth())    // CreateStock create stock
	//c.setSinglePath("GET", "/api/v1/stock/list", middleware.Auth())    // ListStock list of stock by query parameters
	//c.setSinglePath("POST", "/api/v1/stock/setFlashSale", middleware.Auth())    // SetFlashSaleStock 设置秒杀产品的库存，直接更新DB和缓存，这里没有使用dtm+rockscache缓存一致性方案，主要原因是与flashSale服务使用dtm+rockscache操作redis的key相同，会产生冲突。
	//c.setSinglePath("POST", "/api/v1/flashSale", middleware.Auth())    // FlashSale 秒杀抢购
	//c.setSinglePath("POST", "/api/v1/coupon", middleware.Auth())    // CreateCoupon create coupon
	//c.setSinglePath("GET", "/api/v1/coupon/list", middleware.Auth())    // ListCoupon list of coupon by query parameters
}
