package handler

import (
	"errors"

	"user/internal/cache"
	"user/internal/dao"
	"user/internal/ecode"
	"user/internal/model"
	"user/internal/types"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/gin/response"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/mysql/query"
	"github.com/zhufuyi/sponge/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

var _ CourseHandler = (*courseHandler)(nil)

// CourseHandler defining the handler interface
type CourseHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	DeleteByIDs(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	ListByIDs(c *gin.Context)
	List(c *gin.Context)
}

type courseHandler struct {
	iDao dao.CourseDao
}

// NewCourseHandler creating the handler interface
func NewCourseHandler() CourseHandler {
	return &courseHandler{
		iDao: dao.NewCourseDao(
			model.GetDB(),
			cache.NewCourseCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create course
// @Description submit information to create course
// @Tags course
// @accept json
// @Produce json
// @Param data body types.CreateCourseRequest true "course information"
// @Success 200 {object} types.Result{}
// @Router /api/v1/course [post]
func (h *courseHandler) Create(c *gin.Context) {
	form := &types.CreateCourseRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	course := &model.Course{}
	err = copier.Copy(course, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateCourse)
		return
	}

	err = h.iDao.Create(c.Request.Context(), course)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": course.ID})
}

// DeleteByID delete a record by id
// @Summary delete course
// @Description delete course by id
// @Tags course
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.Result{}
// @Router /api/v1/course/{id} [delete]
func (h *courseHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getCourseIDFromPath(c)
	if isAbort {
		return
	}

	err := h.iDao.DeleteByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// DeleteByIDs delete records by batch id
// @Summary delete courses
// @Description delete courses by batch id
// @Tags course
// @Param data body types.DeleteCoursesByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.Result{}
// @Router /api/v1/course/delete/ids [post]
func (h *courseHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteCoursesByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	err = h.iDao.DeleteByIDs(c.Request.Context(), form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// UpdateByID update information by id
// @Summary update course
// @Description update course information by id
// @Tags course
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateCourseByIDRequest true "course information"
// @Success 200 {object} types.Result{}
// @Router /api/v1/course/{id} [put]
func (h *courseHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getCourseIDFromPath(c)
	if isAbort {
		return
	}

	form := &types.UpdateCourseByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	course := &model.Course{}
	err = copier.Copy(course, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateCourse)
		return
	}

	err = h.iDao.UpdateByID(c.Request.Context(), course)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get course detail
// @Description get course detail by id
// @Tags course
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.Result{}
// @Router /api/v1/course/{id} [get]
func (h *courseHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getCourseIDFromPath(c)
	if isAbort {
		return
	}

	course, err := h.iDao.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			logger.Warn("GetByID not found", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByID error", logger.Err(err), logger.Any("id", id), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.GetCourseByIDRespond{}
	err = copier.Copy(data, course)
	if err != nil {
		response.Error(c, ecode.ErrGetCourse)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"course": data})
}

// ListByIDs list of records by batch id
// @Summary list of courses by batch id
// @Description list of courses by batch id
// @Tags course
// @Param data body types.GetCoursesByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.Result{}
// @Router /api/v1/course/list/ids [post]
func (h *courseHandler) ListByIDs(c *gin.Context) {
	form := &types.GetCoursesByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	courseMap, err := h.iDao.GetByIDs(c.Request.Context(), form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())

		return
	}

	courses := []*types.GetCourseByIDRespond{}
	for _, id := range form.IDs {
		if v, ok := courseMap[id]; ok {
			record, err := convertCourse(v)
			if err != nil {
				response.Error(c, ecode.ErrListCourse)
				return
			}
			courses = append(courses, record)
		}
	}

	response.Success(c, gin.H{
		"courses": courses,
	})
}

// List of records by query parameters
// @Summary list of courses by query parameters
// @Description list of courses by paging and conditions
// @Tags course
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.Result{}
// @Router /api/v1/course/list [post]
func (h *courseHandler) List(c *gin.Context) {
	form := &types.GetCoursesRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	courses, total, err := h.iDao.GetByColumns(c.Request.Context(), &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertCourses(courses)
	if err != nil {
		response.Error(c, ecode.ErrListCourse)
		return
	}

	response.Success(c, gin.H{
		"courses": data,
		"total":   total,
	})
}

func getCourseIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return "", 0, true
	}

	return idStr, id, false
}

func convertCourse(course *model.Course) (*types.GetCourseByIDRespond, error) {
	data := &types.GetCourseByIDRespond{}
	err := copier.Copy(data, course)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(course.ID)
	return data, nil
}

func convertCourses(fromValues []*model.Course) ([]*types.GetCourseByIDRespond, error) {
	toValues := []*types.GetCourseByIDRespond{}
	for _, v := range fromValues {
		data, err := convertCourse(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
