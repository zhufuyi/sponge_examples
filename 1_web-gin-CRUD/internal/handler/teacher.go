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

var _ TeacherHandler = (*teacherHandler)(nil)

// TeacherHandler defining the handler interface
type TeacherHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	DeleteByIDs(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	GetByCondition(c *gin.Context)
	ListByIDs(c *gin.Context)
	List(c *gin.Context)
}

type teacherHandler struct {
	iDao dao.TeacherDao
}

// NewTeacherHandler creating the handler interface
func NewTeacherHandler() TeacherHandler {
	return &teacherHandler{
		iDao: dao.NewTeacherDao(
			model.GetDB(),
			cache.NewTeacherCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create teacher
// @Description submit information to create teacher
// @Tags teacher
// @accept json
// @Produce json
// @Param data body types.CreateTeacherRequest true "teacher information"
// @Success 200 {object} types.CreateTeacherRespond{}
// @Router /api/v1/teacher [post]
func (h *teacherHandler) Create(c *gin.Context) {
	form := &types.CreateTeacherRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	teacher := &model.Teacher{}
	err = copier.Copy(teacher, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateTeacher)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, teacher)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": teacher.ID})
}

// DeleteByID delete a record by id
// @Summary delete teacher
// @Description delete teacher by id
// @Tags teacher
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteTeacherByIDRespond{}
// @Router /api/v1/teacher/{id} [delete]
func (h *teacherHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getTeacherIDFromPath(c)
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
// @Summary delete teachers
// @Description delete teachers by batch id
// @Tags teacher
// @Param data body types.DeleteTeachersByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteTeachersByIDsRespond{}
// @Router /api/v1/teacher/delete/ids [post]
func (h *teacherHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteTeachersByIDsRequest{}
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
// @Summary update teacher
// @Description update teacher information by id
// @Tags teacher
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateTeacherByIDRequest true "teacher information"
// @Success 200 {object} types.UpdateTeacherByIDRespond{}
// @Router /api/v1/teacher/{id} [put]
func (h *teacherHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getTeacherIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateTeacherByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	teacher := &model.Teacher{}
	err = copier.Copy(teacher, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDTeacher)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, teacher)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get teacher detail
// @Description get teacher detail by id
// @Tags teacher
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetTeacherByIDRespond{}
// @Router /api/v1/teacher/{id} [get]
func (h *teacherHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getTeacherIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	teacher, err := h.iDao.GetByID(ctx, id)
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

	data := &types.TeacherObjDetail{}
	err = copier.Copy(data, teacher)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDTeacher)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"teacher": data})
}

// GetByCondition get a record by condition
// @Summary get teacher by condition
// @Description get teacher by condition
// @Tags teacher
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetTeacherByConditionRespond{}
// @Router /api/v1/teacher/condition [post]
func (h *teacherHandler) GetByCondition(c *gin.Context) {
	form := &types.GetTeacherByConditionRequest{}
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
	teacher, err := h.iDao.GetByCondition(ctx, &form.Conditions)
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

	data := &types.TeacherObjDetail{}
	err = copier.Copy(data, teacher)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDTeacher)
		return
	}
	data.ID = utils.Uint64ToStr(teacher.ID)

	response.Success(c, gin.H{"teacher": data})
}

// ListByIDs list of records by batch id
// @Summary list of teachers by batch id
// @Description list of teachers by batch id
// @Tags teacher
// @Param data body types.ListTeachersByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListTeachersByIDsRespond{}
// @Router /api/v1/teacher/list/ids [post]
func (h *teacherHandler) ListByIDs(c *gin.Context) {
	form := &types.ListTeachersByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams.WithOutMsg("参数错误"), "详细错误信息")
		response.Output(c, ecode.Unauthorized.WithOutMsg("错误简单描述").ToHTTPCode(), "详细错误信息")
		return
	}

	ctx := middleware.WrapCtx(c)
	teacherMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	teachers := []*types.TeacherObjDetail{}
	for _, id := range form.IDs {
		if v, ok := teacherMap[id]; ok {
			record, err := convertTeacher(v)
			if err != nil {
				response.Error(c, ecode.ErrListTeacher)
				return
			}
			teachers = append(teachers, record)
		}
	}

	response.Success(c, gin.H{
		"teachers": teachers,
	})
}

// List of records by query parameters
// @Summary list of teachers by query parameters
// @Description list of teachers by paging and conditions
// @Tags teacher
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListTeachersRespond{}
// @Router /api/v1/teacher/list [post]
func (h *teacherHandler) List(c *gin.Context) {
	form := &types.ListTeachersRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	teachers, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertTeachers(teachers)
	if err != nil {
		response.Error(c, ecode.ErrListTeacher)
		return
	}

	response.Success(c, gin.H{
		"teachers": data,
		"total":    total,
	})
}

func getTeacherIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertTeacher(teacher *model.Teacher) (*types.TeacherObjDetail, error) {
	data := &types.TeacherObjDetail{}
	err := copier.Copy(data, teacher)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(teacher.ID)
	return data, nil
}

func convertTeachers(fromValues []*model.Teacher) ([]*types.TeacherObjDetail, error) {
	toValues := []*types.TeacherObjDetail{}
	for _, v := range fromValues {
		data, err := convertTeacher(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
