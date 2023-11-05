// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	communityV1 "community/api/community/v1"
	"community/internal/handler"

	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/gin-gonic/gin"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		collectServiceMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			collectServiceRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewCollectServiceHandler())
		})
}

func collectServiceRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService communityV1.CollectServiceLogicer) {
	communityV1.RegisterCollectServiceRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		communityV1.WithCollectServiceHTTPResponse(),
		communityV1.WithCollectServiceLogger(logger.Get()),
		communityV1.WithCollectServiceErrorToHTTPCode(
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
func collectServiceMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/collectService/:id  will take effect
	// c.setGroupPath("/api/v1/collect", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/collect", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/collect/delete", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/collect/list", middleware.Auth())
}