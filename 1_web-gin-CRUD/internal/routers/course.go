package routers

import (
	"user/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		courseRouter(group, handler.NewCourseHandler())
	})
}

func courseRouter(group *gin.RouterGroup, h handler.CourseHandler) {
	//group.Use(middleware.Auth()) // all of the following routes use jwt authentication
	group.POST("/course", h.Create)
	group.DELETE("/course/:id", h.DeleteByID)
	group.POST("/course/delete/ids", h.DeleteByIDs)
	group.PUT("/course/:id", h.UpdateByID)
	group.GET("/course/:id", h.GetByID)
	group.POST("/course/condition", h.GetByCondition)
	group.POST("/course/list/ids", h.ListByIDs)
	group.POST("/course/list", h.List)
}
