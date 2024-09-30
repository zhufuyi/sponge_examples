// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"

	stockV1 "stock/api/stock/v1"
	"stock/internal/handler"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		atomicMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			atomicRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewAtomicHandler())
		})
}

func atomicRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService stockV1.AtomicLogicer) {
	ctxFn := func(c *gin.Context) context.Context {
		md := metadata.New(map[string]string{
			middleware.ContextRequestIDKey: middleware.GCtxRequestID(c),
		})
		return metadata.NewIncomingContext(c.Request.Context(), md)
	}
	stockV1.RegisterAtomicRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		stockV1.WithAtomicLogger(logger.Get()),
		stockV1.WithAtomicRPCResponse(),
		stockV1.WithAtomicWrapCtx(ctxFn),
		stockV1.WithAtomicErrorToHTTPCode(
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
func atomicMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/atomic/:id  will take effect
	// c.setGroupPath("/api/v1/atomic", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("PUT", "/api/v1/stock/:id/atomic", middleware.Auth())
	//c.setSinglePath("GET", "/api/v1/stock/:id/atomic", middleware.Auth())
}