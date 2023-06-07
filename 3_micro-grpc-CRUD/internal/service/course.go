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

// nolint
func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		userV1.RegisterCourseServiceServer(server, NewCourseServiceServer()) // register service to the rpc service
	})
}

var _ userV1.CourseServiceServer = (*courseService)(nil)

type courseService struct {
	userV1.UnimplementedCourseServiceServer

	iDao dao.CourseDao
}

// NewCourseServiceServer create a new service
func NewCourseServiceServer() userV1.CourseServiceServer {
	return &courseService{
		iDao: dao.NewCourseDao(
			model.GetDB(),
			cache.NewCourseCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *courseService) Create(ctx context.Context, req *userV1.CreateCourseRequest) (*userV1.CreateCourseReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	course := &model.Course{}
	err = copier.Copy(course, req)
	if err != nil {
		return nil, ecode.StatusCreateCourse.Err()
	}

	err = s.iDao.Create(ctx, course)
	if err != nil {
		logger.Error("s.iDao.Create error", logger.Err(err), logger.Any("course", course), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.CreateCourseReply{Id: course.ID}, nil
}

// DeleteByID delete a record by id
func (s *courseService) DeleteByID(ctx context.Context, req *userV1.DeleteCourseByIDRequest) (*userV1.DeleteCourseByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	err = s.iDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Error("s.iDao.DeleteByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteCourseByIDReply{}, nil
}

// DeleteByIDs delete records by batch id
func (s *courseService) DeleteByIDs(ctx context.Context, req *userV1.DeleteCourseByIDsRequest) (*userV1.DeleteCourseByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	err = s.iDao.DeleteByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("s.iDao.DeleteByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteCourseByIDsReply{}, nil
}

// UpdateByID update a record by id
func (s *courseService) UpdateByID(ctx context.Context, req *userV1.UpdateCourseByIDRequest) (*userV1.UpdateCourseByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	course := &model.Course{}
	err = copier.Copy(course, req)
	if err != nil {
		return nil, ecode.StatusUpdateCourse.Err()
	}
	course.ID = req.Id

	err = s.iDao.UpdateByID(ctx, course)
	if err != nil {
		logger.Error("s.iDao.UpdateByID error", logger.Err(err), logger.Any("course", course), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.UpdateCourseByIDReply{}, nil
}

// GetByID get a record by id
func (s *courseService) GetByID(ctx context.Context, req *userV1.GetCourseByIDRequest) (*userV1.GetCourseByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record, err := s.iDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			logger.Warn("s.iDao.GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("s.iDao.GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertCourse(record)
	if err != nil {
		logger.Warn("convertCourse error", logger.Err(err), logger.Any("record", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetCourse.Err()
	}

	return &userV1.GetCourseByIDReply{Course: data}, nil
}

// ListByIDs list of records by batch id
func (s *courseService) ListByIDs(ctx context.Context, req *userV1.ListCourseByIDsRequest) (*userV1.ListCourseByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	courseMap, err := s.iDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("s.iDao.GetByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
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
func (s *courseService) List(ctx context.Context, req *userV1.ListCourseRequest) (*userV1.ListCourseReply, error) {
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

	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("s.iDao.GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("s.iDao.GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	courses := []*userV1.Course{}
	for _, record := range records {
		course, err := convertCourse(record)
		if err != nil {
			logger.Warn("convertCourse error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		courses = append(courses, course)
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
