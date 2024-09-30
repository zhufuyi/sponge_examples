// Code generated by https://github.com/zhufuyi/sponge

package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/zhufuyi/sponge/pkg/logger"
	//"github.com/zhufuyi/sponge/pkg/middleware"

	flashSaleV1 "flashSale/api/flashSale/v1"
	"flashSale/internal/handler"
)

func init() {
	allMiddlewareFns = append(allMiddlewareFns, func(c *middlewareConfig) {
		flashSaleMiddlewares(c)
	})

	allRouteFns = append(allRouteFns,
		func(r *gin.Engine, groupPathMiddlewares map[string][]gin.HandlerFunc, singlePathMiddlewares map[string][]gin.HandlerFunc) {
			flashSaleRouter(r, groupPathMiddlewares, singlePathMiddlewares, handler.NewFlashSaleHandler())
		})
}

func flashSaleRouter(
	r *gin.Engine,
	groupPathMiddlewares map[string][]gin.HandlerFunc,
	singlePathMiddlewares map[string][]gin.HandlerFunc,
	iService flashSaleV1.FlashSaleLogicer) {
	flashSaleV1.RegisterFlashSaleRouter(
		r,
		groupPathMiddlewares,
		singlePathMiddlewares,
		iService,
		flashSaleV1.WithFlashSaleLogger(logger.Get()),
		flashSaleV1.WithFlashSaleHTTPResponse(),
		flashSaleV1.WithFlashSaleErrorToHTTPCode(
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
func flashSaleMiddlewares(c *middlewareConfig) {
	// set up group route middleware, group path is left prefix rules,
	// if the left prefix is hit, the middleware will take effect, e.g. group route is /api/v1, route /api/v1/flashSale/:id  will take effect
	// c.setGroupPath("/api/v1/flashSale", middleware.Auth())

	// set up single route middleware, just uncomment the code and fill in the middlewares, nothing else needs to be changed
	//c.setSinglePath("POST", "/api/v1/flashSale/setProductStock", middleware.Auth())    // SetProductStock 设置库存数量
	//c.setSinglePath("POST", "/api/v1/flashSale", middleware.Auth())    // FlashSale 秒杀抢购
	//c.setSinglePath("GET", "/api/v1/flashSale/redisQueryPrepared", middleware.Auth())    // RedisQueryPrepared 反查redis数据
	//c.setSinglePath("POST", "/api/v1/flashSale/sendSubmitOrderNotify", middleware.Auth())    // SendSubmitOrderNotify 发送提交订单通知到到消息队列
}