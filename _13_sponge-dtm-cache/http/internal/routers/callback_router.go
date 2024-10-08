// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/zhufuyi/sponge/pkg/logger"
	//"github.com/zhufuyi/sponge/pkg/middleware"

	stockV1 "stock/api/stock/v1"
	"stock/internal/handler"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		callbackMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			callbackRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewCallbackHandler())
		})
}

func callbackRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService stockV1.CallbackLogicer) {
	stockV1.RegisterCallbackRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		stockV1.WithCallbackLogger(logger.Get()),
		stockV1.WithCallbackHTTPResponse(),
		stockV1.WithCallbackErrorToHTTPCode(
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
func callbackMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/callback/:id  will take effect
	// c.setGroupPath("/api/v1/callback", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("GET", "/api/v1/stock/queryPrepared", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/stock/deleteCache", middleware.Auth())
}
