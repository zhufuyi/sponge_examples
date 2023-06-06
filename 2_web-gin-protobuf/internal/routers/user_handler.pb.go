// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	userV1 "user/api/user/v1"
	"user/internal/handler"

	"github.com/zhufuyi/sponge/pkg/logger"
	//"github.com/zhufuyi/sponge/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		userMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			userRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewUserHandler())
		})
}

func userRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService userV1.UserLogicer) {
	userV1.RegisterUserRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		userV1.WithUserHTTPResponse(),
		userV1.WithUserLogger(logger.Get()),
		userV1.WithUserErrorToHTTPCode(
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
func userMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/user/:id  will take effect
	// c.setGroupPath("/api/v1/user", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/auth/register", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/auth/login", middleware.Auth())
	//c.setSinglePath("POST", "/api/v1/auth/logout", middleware.Auth())
}
