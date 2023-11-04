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

var _ TeachHandler = (*teachHandler)(nil)

// TeachHandler defining the handler interface
type TeachHandler interface {
	Create(c *gin.Context)
	DeleteByID(c *gin.Context)
	DeleteByIDs(c *gin.Context)
	UpdateByID(c *gin.Context)
	GetByID(c *gin.Context)
	GetByCondition(c *gin.Context)
	ListByIDs(c *gin.Context)
	List(c *gin.Context)
}

type teachHandler struct {
	iDao dao.TeachDao
}

// NewTeachHandler creating the handler interface
func NewTeachHandler() TeachHandler {
	return &teachHandler{
		iDao: dao.NewTeachDao(
			model.GetDB(),
			cache.NewTeachCache(model.GetCacheType()),
		),
	}
}

// Create a record
// @Summary create teach
// @Description submit information to create teach
// @Tags teach
// @accept json
// @Produce json
// @Param data body types.CreateTeachRequest true "teach information"
// @Success 200 {object} types.CreateTeachRespond{}
// @Router /api/v1/teach [post]
func (h *teachHandler) Create(c *gin.Context) {
	form := &types.CreateTeachRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	teach := &model.Teach{}
	err = copier.Copy(teach, form)
	if err != nil {
		response.Error(c, ecode.ErrCreateTeach)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.Create(ctx, teach)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c, gin.H{"id": teach.ID})
}

// DeleteByID delete a record by id
// @Summary delete teach
// @Description delete teach by id
// @Tags teach
// @accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} types.DeleteTeachByIDRespond{}
// @Router /api/v1/teach/{id} [delete]
func (h *teachHandler) DeleteByID(c *gin.Context) {
	_, id, isAbort := getTeachIDFromPath(c)
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
// @Summary delete teachs
// @Description delete teachs by batch id
// @Tags teach
// @Param data body types.DeleteTeachsByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.DeleteTeachsByIDsRespond{}
// @Router /api/v1/teach/delete/ids [post]
func (h *teachHandler) DeleteByIDs(c *gin.Context) {
	form := &types.DeleteTeachsByIDsRequest{}
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
// @Summary update teach
// @Description update teach information by id
// @Tags teach
// @accept json
// @Produce json
// @Param id path string true "id"
// @Param data body types.UpdateTeachByIDRequest true "teach information"
// @Success 200 {object} types.UpdateTeachByIDRespond{}
// @Router /api/v1/teach/{id} [put]
func (h *teachHandler) UpdateByID(c *gin.Context) {
	_, id, isAbort := getTeachIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	form := &types.UpdateTeachByIDRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}
	form.ID = id

	teach := &model.Teach{}
	err = copier.Copy(teach, form)
	if err != nil {
		response.Error(c, ecode.ErrUpdateByIDTeach)
		return
	}

	ctx := middleware.WrapCtx(c)
	err = h.iDao.UpdateByID(ctx, teach)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	response.Success(c)
}

// GetByID get a record by id
// @Summary get teach detail
// @Description get teach detail by id
// @Tags teach
// @Param id path string true "id"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetTeachByIDRespond{}
// @Router /api/v1/teach/{id} [get]
func (h *teachHandler) GetByID(c *gin.Context) {
	idStr, id, isAbort := getTeachIDFromPath(c)
	if isAbort {
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	teach, err := h.iDao.GetByID(ctx, id)
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

	data := &types.TeachObjDetail{}
	err = copier.Copy(data, teach)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDTeach)
		return
	}
	data.ID = idStr

	response.Success(c, gin.H{"teach": data})
}

// GetByCondition get a record by condition
// @Summary get teach by condition
// @Description get teach by condition
// @Tags teach
// @Param data body types.Conditions true "query condition"
// @Accept json
// @Produce json
// @Success 200 {object} types.GetTeachByConditionRespond{}
// @Router /api/v1/teach/condition [post]
func (h *teachHandler) GetByCondition(c *gin.Context) {
	form := &types.GetTeachByConditionRequest{}
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
	teach, err := h.iDao.GetByCondition(ctx, &form.Conditions)
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

	data := &types.TeachObjDetail{}
	err = copier.Copy(data, teach)
	if err != nil {
		response.Error(c, ecode.ErrGetByIDTeach)
		return
	}
	data.ID = utils.Uint64ToStr(teach.ID)

	response.Success(c, gin.H{"teach": data})
}

// ListByIDs list of records by batch id
// @Summary list of teachs by batch id
// @Description list of teachs by batch id
// @Tags teach
// @Param data body types.ListTeachsByIDsRequest true "id array"
// @Accept json
// @Produce json
// @Success 200 {object} types.ListTeachsByIDsRespond{}
// @Router /api/v1/teach/list/ids [post]
func (h *teachHandler) ListByIDs(c *gin.Context) {
	form := &types.ListTeachsByIDsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams.WithOutMsg("参数错误"), "详细错误信息")
		response.Output(c, ecode.Unauthorized.WithOutMsg("错误简单描述").ToHTTPCode(), "详细错误信息")
		return
	}

	ctx := middleware.WrapCtx(c)
	teachMap, err := h.iDao.GetByIDs(ctx, form.IDs)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	teachs := []*types.TeachObjDetail{}
	for _, id := range form.IDs {
		if v, ok := teachMap[id]; ok {
			record, err := convertTeach(v)
			if err != nil {
				response.Error(c, ecode.ErrListTeach)
				return
			}
			teachs = append(teachs, record)
		}
	}

	response.Success(c, gin.H{
		"teachs": teachs,
	})
}

// List of records by query parameters
// @Summary list of teachs by query parameters
// @Description list of teachs by paging and conditions
// @Tags teach
// @accept json
// @Produce json
// @Param data body types.Params true "query parameters"
// @Success 200 {object} types.ListTeachsRespond{}
// @Router /api/v1/teach/list [post]
func (h *teachHandler) List(c *gin.Context) {
	form := &types.ListTeachsRequest{}
	err := c.ShouldBindJSON(form)
	if err != nil {
		logger.Warn("ShouldBindJSON error: ", logger.Err(err), middleware.GCtxRequestIDField(c))
		response.Error(c, ecode.InvalidParams)
		return
	}

	ctx := middleware.WrapCtx(c)
	teachs, total, err := h.iDao.GetByColumns(ctx, &form.Params)
	if err != nil {
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("form", form), middleware.GCtxRequestIDField(c))
		response.Output(c, ecode.InternalServerError.ToHTTPCode())
		return
	}

	data, err := convertTeachs(teachs)
	if err != nil {
		response.Error(c, ecode.ErrListTeach)
		return
	}

	response.Success(c, gin.H{
		"teachs": data,
		"total":  total,
	})
}

func getTeachIDFromPath(c *gin.Context) (string, uint64, bool) {
	idStr := c.Param("id")
	id, err := utils.StrToUint64E(idStr)
	if err != nil || id == 0 {
		logger.Warn("StrToUint64E error: ", logger.String("idStr", idStr), middleware.GCtxRequestIDField(c))
		return "", 0, true
	}

	return idStr, id, false
}

func convertTeach(teach *model.Teach) (*types.TeachObjDetail, error) {
	data := &types.TeachObjDetail{}
	err := copier.Copy(data, teach)
	if err != nil {
		return nil, err
	}
	data.ID = utils.Uint64ToStr(teach.ID)
	return data, nil
}

func convertTeachs(fromValues []*model.Teach) ([]*types.TeachObjDetail, error) {
	toValues := []*types.TeachObjDetail{}
	for _, v := range fromValues {
		data, err := convertTeach(v)
		if err != nil {
			return nil, err
		}
		toValues = append(toValues, data)
	}

	return toValues, nil
}
