package service

import (
	"context"
	"errors"
	"strings"

	userV1 "user/api/user/v1"
	"user/internal/cache"
	"user/internal/dao"
	"user/internal/ecode"
	"user/internal/model"

	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/mysql/query"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		userV1.RegisterTeacherServer(server, NewTeacherServer()) // register service to the rpc service
	})
}

var _ userV1.TeacherServer = (*teacher)(nil)

type teacher struct {
	userV1.UnimplementedTeacherServer

	iDao dao.TeacherDao
}

// NewTeacherServer create a new service
func NewTeacherServer() userV1.TeacherServer {
	return &teacher{
		iDao: dao.NewTeacherDao(
			model.GetDB(),
			cache.NewTeacherCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *teacher) Create(ctx context.Context, req *userV1.CreateTeacherRequest) (*userV1.CreateTeacherReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record := &model.Teacher{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusCreateTeacher.Err()
	}

	//ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.Create(ctx, record)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("teacher", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.CreateTeacherReply{Id: record.ID}, nil
}

// DeleteByID delete a record by id
func (s *teacher) DeleteByID(ctx context.Context, req *userV1.DeleteTeacherByIDRequest) (*userV1.DeleteTeacherByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	//ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteTeacherByIDReply{}, nil
}

// DeleteByIDs delete records by batch id
func (s *teacher) DeleteByIDs(ctx context.Context, req *userV1.DeleteTeacherByIDsRequest) (*userV1.DeleteTeacherByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	//ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.DeleteByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteTeacherByIDsReply{}, nil
}

// UpdateByID update a record by id
func (s *teacher) UpdateByID(ctx context.Context, req *userV1.UpdateTeacherByIDRequest) (*userV1.UpdateTeacherByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record := &model.Teacher{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDTeacher.Err()
	}
	record.ID = req.Id

	//ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("teacher", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.UpdateTeacherByIDReply{}, nil
}

// GetByID get a record by id
func (s *teacher) GetByID(ctx context.Context, req *userV1.GetTeacherByIDRequest) (*userV1.GetTeacherByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	//ctx = interceptor.WrapServerCtx(ctx)
	record, err := s.iDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertTeacher(record)
	if err != nil {
		logger.Warn("convertTeacher error", logger.Err(err), logger.Any("teacher", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDTeacher.Err()
	}

	return &userV1.GetTeacherByIDReply{Teacher: data}, nil
}

// GetByCondition get a record by id
func (s *teacher) GetByCondition(ctx context.Context, req *userV1.GetTeacherByConditionRequest) (*userV1.GetTeacherByConditionReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	//ctx = interceptor.WrapServerCtx(ctx)
	conditions := &query.Conditions{}
	for _, v := range req.Conditions.GetColumns() {
		column := query.Column{}
		_ = copier.Copy(&column, v)
		conditions.Columns = append(conditions.Columns, column)
	}
	err = conditions.CheckValid()
	if err != nil {
		logger.Warn("Parameters error", logger.Err(err), logger.Any("conditions", conditions), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record, err := s.iDao.GetByCondition(ctx, conditions)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			logger.Warn("GetByCondition error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("GetByCondition error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertTeacher(record)
	if err != nil {
		logger.Warn("convertTeacher error", logger.Err(err), logger.Any("teacher", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByConditionTeacher.Err()
	}

	return &userV1.GetTeacherByConditionReply{
		Teacher: data,
	}, nil
}

// ListByIDs list of records by batch id
func (s *teacher) ListByIDs(ctx context.Context, req *userV1.ListTeacherByIDsRequest) (*userV1.ListTeacherByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	//ctx = interceptor.WrapServerCtx(ctx)
	teacherMap, err := s.iDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	teachers := []*userV1.Teacher{}
	for _, id := range req.Ids {
		if v, ok := teacherMap[id]; ok {
			record, err := convertTeacher(v)
			if err != nil {
				logger.Warn("convertTeacher error", logger.Err(err), logger.Any("teacher", v), interceptor.ServerCtxRequestIDField(ctx))
				return nil, ecode.StatusInternalServerError.ToRPCErr()
			}
			teachers = append(teachers, record)
		}
	}

	return &userV1.ListTeacherByIDsReply{Teachers: teachers}, nil
}

// List of records by query parameters
func (s *teacher) List(ctx context.Context, req *userV1.ListTeacherRequest) (*userV1.ListTeacherReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListTeacher.Err()
	}
	params.Size = int(req.Params.Limit)

	//ctx = interceptor.WrapServerCtx(ctx)
	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	teachers := []*userV1.Teacher{}
	for _, record := range records {
		data, err := convertTeacher(record)
		if err != nil {
			logger.Warn("convertTeacher error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		teachers = append(teachers, data)
	}

	return &userV1.ListTeacherReply{
		Total:    total,
		Teachers: teachers,
	}, nil
}

func convertTeacher(record *model.Teacher) (*userV1.Teacher, error) {
	value := &userV1.Teacher{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	value.Id = record.ID
	value.CreatedAt = record.CreatedAt.Unix()
	value.UpdatedAt = record.UpdatedAt.Unix()
	return value, nil
}
