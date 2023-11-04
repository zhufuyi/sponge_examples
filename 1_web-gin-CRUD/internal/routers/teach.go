package routers

import (
	"user/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		teachRouter(group, handler.NewTeachHandler())
	})
}

func teachRouter(group *gin.RouterGroup, h handler.TeachHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	group.POST("/teach", h.Create)
	group.DELETE("/teach/:id", h.DeleteByID)
	group.POST("/teach/delete/ids", h.DeleteByIDs)
	group.PUT("/teach/:id", h.UpdateByID)
	group.GET("/teach/:id", h.GetByID)
	group.POST("/teach/condition", h.GetByCondition)
	group.POST("/teach/list/ids", h.ListByIDs)
	group.POST("/teach/list", h.List)
}
