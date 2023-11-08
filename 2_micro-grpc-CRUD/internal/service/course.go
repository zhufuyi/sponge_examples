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
		userV1.RegisterCourseServer(server, NewCourseServer()) // register service to the rpc service
	})
}

var _ userV1.CourseServer = (*course)(nil)

type course struct {
	userV1.UnimplementedCourseServer

	iDao dao.CourseDao
}

// NewCourseServer create a new service
func NewCourseServer() userV1.CourseServer {
	return &course{
		iDao: dao.NewCourseDao(
			model.GetDB(),
			cache.NewCourseCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *course) Create(ctx context.Context, req *userV1.CreateCourseRequest) (*userV1.CreateCourseReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record := &model.Course{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusCreateCourse.Err()
	}

	ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.Create(ctx, record)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("course", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.CreateCourseReply{Id: record.ID}, nil
}

// DeleteByID delete a record by id
func (s *course) DeleteByID(ctx context.Context, req *userV1.DeleteCourseByIDRequest) (*userV1.DeleteCourseByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteCourseByIDReply{}, nil
}

// DeleteByIDs delete records by batch id
func (s *course) DeleteByIDs(ctx context.Context, req *userV1.DeleteCourseByIDsRequest) (*userV1.DeleteCourseByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.DeleteByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteCourseByIDsReply{}, nil
}

// UpdateByID update a record by id
func (s *course) UpdateByID(ctx context.Context, req *userV1.UpdateCourseByIDRequest) (*userV1.UpdateCourseByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record := &model.Course{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDCourse.Err()
	}
	record.ID = req.Id

	ctx = interceptor.WrapServerCtx(ctx)
	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("course", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.UpdateCourseByIDReply{}, nil
}

// GetByID get a record by id
func (s *course) GetByID(ctx context.Context, req *userV1.GetCourseByIDRequest) (*userV1.GetCourseByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = interceptor.WrapServerCtx(ctx)
	record, err := s.iDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertCourse(record)
	if err != nil {
		logger.Warn("convertCourse error", logger.Err(err), logger.Any("course", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDCourse.Err()
	}

	return &userV1.GetCourseByIDReply{Course: data}, nil
}

// GetByCondition get a record by id
func (s *course) GetByCondition(ctx context.Context, req *userV1.GetCourseByConditionRequest) (*userV1.GetCourseByConditionReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = interceptor.WrapServerCtx(ctx)
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

	data, err := convertCourse(record)
	if err != nil {
		logger.Warn("convertCourse error", logger.Err(err), logger.Any("course", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByConditionCourse.Err()
	}

	return &userV1.GetCourseByConditionReply{
		Course: data,
	}, nil
}

// ListByIDs list of records by batch id
func (s *course) ListByIDs(ctx context.Context, req *userV1.ListCourseByIDsRequest) (*userV1.ListCourseByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = interceptor.WrapServerCtx(ctx)
	courseMap, err := s.iDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	courses := []*userV1.Course{}
	for _, id := range req.Ids {
		if v, ok := courseMap[id]; ok {
			record, err := convertCourse(v)
			if err != nil {
				logger.Warn("convertCourse error", logger.Err(err), logger.Any("course", v), interceptor.ServerCtxRequestIDField(ctx))
				return nil, ecode.StatusInternalServerError.ToRPCErr()
			}
			courses = append(courses, record)
		}
	}

	return &userV1.ListCourseByIDsReply{Courses: courses}, nil
}

// List of records by query parameters
func (s *course) List(ctx context.Context, req *userV1.ListCourseRequest) (*userV1.ListCourseReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListCourse.Err()
	}
	params.Size = int(req.Params.Limit)

	ctx = interceptor.WrapServerCtx(ctx)
	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	courses := []*userV1.Course{}
	for _, record := range records {
		data, err := convertCourse(record)
		if err != nil {
			logger.Warn("convertCourse error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		courses = append(courses, data)
	}

	return &userV1.ListCourseReply{
		Total:   total,
		Courses: courses,
	}, nil
}

func convertCourse(record *model.Course) (*userV1.Course, error) {
	value := &userV1.Course{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	value.Id = record.ID
	value.CreatedAt = record.CreatedAt.Unix()
	value.UpdatedAt = record.UpdatedAt.Unix()
	return value, nil
}
