package routers

import (
	"user/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	apiV1RouterFns = append(apiV1RouterFns, func(group *gin.RouterGroup) {
		teacherRouter(group, handler.NewTeacherHandler())
	})
}

func teacherRouter(group *gin.RouterGroup, h handler.TeacherHandler) {
	group.POST("/teacher", h.Create)
	group.DELETE("/teacher/:id", h.DeleteByID)
	group.POST("/teacher/delete/ids", h.DeleteByIDs)
	group.PUT("/teacher/:id", h.UpdateByID)
	group.GET("/teacher/:id", h.GetByID)
	group.POST("/teacher/list/ids", h.ListByIDs)
	group.POST("/teacher/list", h.List)
}
