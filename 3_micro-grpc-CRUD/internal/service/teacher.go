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
		userV1.RegisterTeacherServiceServer(server, NewTeacherServiceServer()) // register service to the rpc service
	})
}

var _ userV1.TeacherServiceServer = (*teacherService)(nil)

type teacherService struct {
	userV1.UnimplementedTeacherServiceServer

	iDao dao.TeacherDao
}

// NewTeacherServiceServer create a new service
func NewTeacherServiceServer() userV1.TeacherServiceServer {
	return &teacherService{
		iDao: dao.NewTeacherDao(
			model.GetDB(),
			cache.NewTeacherCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *teacherService) Create(ctx context.Context, req *userV1.CreateTeacherRequest) (*userV1.CreateTeacherReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	teacher := &model.Teacher{}
	err = copier.Copy(teacher, req)
	if err != nil {
		return nil, ecode.StatusCreateTeacher.Err()
	}

	err = s.iDao.Create(ctx, teacher)
	if err != nil {
		logger.Error("s.iDao.Create error", logger.Err(err), logger.Any("teacher", teacher), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.CreateTeacherReply{Id: teacher.ID}, nil
}

// DeleteByID delete a record by id
func (s *teacherService) DeleteByID(ctx context.Context, req *userV1.DeleteTeacherByIDRequest) (*userV1.DeleteTeacherByIDReply, error) {
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

	return &userV1.DeleteTeacherByIDReply{}, nil
}

// DeleteByIDs delete records by batch id
func (s *teacherService) DeleteByIDs(ctx context.Context, req *userV1.DeleteTeacherByIDsRequest) (*userV1.DeleteTeacherByIDsReply, error) {
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

	return &userV1.DeleteTeacherByIDsReply{}, nil
}

// UpdateByID update a record by id
func (s *teacherService) UpdateByID(ctx context.Context, req *userV1.UpdateTeacherByIDRequest) (*userV1.UpdateTeacherByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	teacher := &model.Teacher{}
	err = copier.Copy(teacher, req)
	if err != nil {
		return nil, ecode.StatusUpdateTeacher.Err()
	}
	teacher.ID = req.Id

	err = s.iDao.UpdateByID(ctx, teacher)
	if err != nil {
		logger.Error("s.iDao.UpdateByID error", logger.Err(err), logger.Any("teacher", teacher), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.UpdateTeacherByIDReply{}, nil
}

// GetByID get a record by id
func (s *teacherService) GetByID(ctx context.Context, req *userV1.GetTeacherByIDRequest) (*userV1.GetTeacherByIDReply, error) {
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

	data, err := convertTeacher(record)
	if err != nil {
		logger.Warn("convertTeacher error", logger.Err(err), logger.Any("record", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetTeacher.Err()
	}

	return &userV1.GetTeacherByIDReply{Teacher: data}, nil
}

// ListByIDs list of records by batch id
func (s *teacherService) ListByIDs(ctx context.Context, req *userV1.ListTeacherByIDsRequest) (*userV1.ListTeacherByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	teacherMap, err := s.iDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("s.iDao.GetByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
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
func (s *teacherService) List(ctx context.Context, req *userV1.ListTeacherRequest) (*userV1.ListTeacherReply, error) {
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

	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("s.iDao.GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("s.iDao.GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	teachers := []*userV1.Teacher{}
	for _, record := range records {
		teacher, err := convertTeacher(record)
		if err != nil {
			logger.Warn("convertTeacher error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		teachers = append(teachers, teacher)
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
