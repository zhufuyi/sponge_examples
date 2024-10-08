// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"

	couponV1 "coupon/api/coupon/v1"
	"coupon/internal/handler"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		couponMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			couponRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewCouponHandler())
		})
}

func couponRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService couponV1.CouponLogicer) {
	ctxFn := func(c *gin.Context) context.Context {
		md := metadata.New(map[string]string{
			middleware.ContextRequestIDKey: middleware.GCtxRequestID(c),
		})
		return metadata.NewIncomingContext(c.Request.Context(), md)
	}
	couponV1.RegisterCouponRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		couponV1.WithCouponLogger(logger.Get()),
		couponV1.WithCouponRPCResponse(),
		couponV1.WithCouponWrapCtx(ctxFn),
		couponV1.WithCouponErrorToHTTPCode(
		// Set some error codes to standard http return codes,
		// by default there is already ecode.InternalServerError and ecode.ServiceUnavailable
		// example:
		// 	ecode.Forbidden, ecode.LimitExceed,
		),
	)
}

// you can set the middleware of a route group, or set the middleware of a single route,
// or you can mix them, pay attention to the duplication of middleware when mixing them,
// it is recommended to set the middleware of a single route in preference
func couponMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/coupon/:id  will take effect
	// c.setGroupPath("/api/v1/coupon", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/coupon/use", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/coupon/useRevert", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/coupon", middleware.Auth())
	//c.setSinglePath("DELETE", "/api/v1/coupon/:id", middleware.Auth())
	//c.setSinglePath("PUT", "/api/v1/coupon/:id", middleware.Auth())
	//c.setSinglePath("GET", "/api/v1/coupon/:id", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/coupon/list", middleware.Auth())
}
