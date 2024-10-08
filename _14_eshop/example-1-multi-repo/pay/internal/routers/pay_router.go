// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"

	payV1 "pay/api/pay/v1"
	"pay/internal/handler"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		payMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			payRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewPayHandler())
		})
}

func payRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService payV1.PayLogicer) {
	ctxFn := func(c *gin.Context) context.Context {
		md := metadata.New(map[string]string{
			middleware.ContextRequestIDKey: middleware.GCtxRequestID(c),
		})
		return metadata.NewIncomingContext(c.Request.Context(), md)
	}
	payV1.RegisterPayRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		payV1.WithPayLogger(logger.Get()),
		payV1.WithPayRPCResponse(),
		payV1.WithPayWrapCtx(ctxFn),
		payV1.WithPayErrorToHTTPCode(
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
func payMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/pay/:id  will take effect
	// c.setGroupPath("/api/v1/pay", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/pay/create", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/pay/createRevert", middleware.Auth())
	//c.setSinglePath("DELETE", "/api/v1/pay/:id", middleware.Auth())
	//c.setSinglePath("PUT", "/api/v1/pay/:id", middleware.Auth())
	//c.setSinglePath("GET", "/api/v1/pay/:id", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/pay/list", middleware.Auth())
}
