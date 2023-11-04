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
	GetByCondition(c *gin.Context)
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
// @Success 200 {object} types.CreateCourseRespond{}
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

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, course)
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
// @Success 200 {object} types.DeleteCourseByIDRespond{}
// @Router /api/v1/course/{id} [delete]
func (h *courseHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getCourseIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err := h.iDao.DeleteByID(ctx, id)
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
// @Success 200 {object} types.DeleteCoursesByIDsRespond{}
// @Router /api/v1/course/delete/ids [post]
func (h *courseHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteCoursesByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.DeleteByIDs(ctx, form.IDs)
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
// @Success 200 {object} types.UpdateCourseByIDRespond{}
// @Router /api/v1/course/{id} [put]
func (h *courseHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getCourseIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
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
		response.Error(c, ecode.ErrUpdateByIDCourse)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, course)
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
// @Success 200 {object} types.GetCourseByIDRespond{}
// @Router /api/v1/course/{id} [get]
func (h *courseHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getCourseIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	course, err := h.iDao.GetByID(ctx, id)
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

	data := &types.CourseObjDetail{}
	err = copier.Copy(data, course)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDCourse)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"course": data})
}

// GetByCondition get a record by condition
// @Summary get course by condition
// @Description get course by condition
// @Tags course
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetCourseByConditionRespond{}
// @Router /api/v1/course/condition [post]
func (h *courseHandler) GetByCondition(c *gin.Context) {
	form := &types.GetCourseByConditionRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	err = form.Conditions.CheckValid()
	if err != nil {
		logger.Warn("Parameters error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	course, err := h.iDao.GetByCondition(ctx, &form.Conditions)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			logger.Warn("GetByCondition not found", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
			response.Error(c, ecode.NotFound)
		} else {
			logger.Error("GetByCondition error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
			response.Output(c, ecode.InternalServerError.ToHTTPCode())
		}
		return
	}

	data := &types.CourseObjDetail{}
	err = copier.Copy(data, course)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDCourse)
		return
	}
	data.ID = utils.Uint64ToStr(course.ID)

	response.Success(c, gin.H{"course": data})
}

// ListByIDs list of records by batch id
// @Summary list of courses by batch id
// @Description list of courses by batch id
// @Tags course
// @Param data body types.ListCoursesByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListCoursesByIDsRespond{}
// @Router /api/v1/course/list/ids [post]
func (h *courseHandler) ListByIDs(c *gin.Context) {
	form := &types.ListCoursesByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams.WithOutMsg("参数错误"), "详细错误信息")
		response.Output(c, ecode.Unauthorized.WithOutMsg("错误简单描述").ToHTTPCode(), "详细错误信息")
		return
	}

	ctx := middleware.WrapCtx(c)
	courseMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	courses := []*types.CourseObjDetail{}
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
// @Success 200 {object} types.ListCoursesRespond{}
// @Router /api/v1/course/list [post]
func (h *courseHandler) List(c *gin.Context) {
	form := &types.ListCoursesRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	courses, total, err := h.iDao.GetByColumns(ctx, &form.Params)
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
		return "", 0, true
	}

	return idStr, id, false
}

func convertCourse(course *model.Course) (*types.CourseObjDetail, error) {
	data := &types.CourseObjDetail{}
	err := copier.Copy(data, course)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(course.ID)
	return data, nil
}

func convertCourses(fromValues []*model.Course) ([]*types.CourseObjDetail, error) {
	toValues := []*types.CourseObjDetail{}
	for _, v := range fromValues {
		data, err := convertCourse(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
