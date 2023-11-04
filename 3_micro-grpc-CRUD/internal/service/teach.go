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
		userV1.RegisterTeachServer(server, NewTeachServer()) // register service to the rpc service
	})
}

var _ userV1.TeachServer = (*teach)(nil)

type teach struct {
	userV1.UnimplementedTeachServer

	iDao dao.TeachDao
}

// NewTeachServer create a new service
func NewTeachServer() userV1.TeachServer {
	return &teach{
		iDao: dao.NewTeachDao(
			model.GetDB(),
			cache.NewTeachCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *teach) Create(ctx context.Context, req *userV1.CreateTeachRequest) (*userV1.CreateTeachReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record := &model.Teach{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusCreateTeach.Err()
	}

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
	err = s.iDao.Create(ctx, record)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("teach", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.CreateTeachReply{Id: record.ID}, nil
}

// DeleteByID delete a record by id
func (s *teach) DeleteByID(ctx context.Context, req *userV1.DeleteTeachByIDRequest) (*userV1.DeleteTeachByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
	err = s.iDao.DeleteByID(ctx, req.Id)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteTeachByIDReply{}, nil
}

// DeleteByIDs delete records by batch id
func (s *teach) DeleteByIDs(ctx context.Context, req *userV1.DeleteTeachByIDsRequest) (*userV1.DeleteTeachByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
	err = s.iDao.DeleteByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("DeleteByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.DeleteTeachByIDsReply{}, nil
}

// UpdateByID update a record by id
func (s *teach) UpdateByID(ctx context.Context, req *userV1.UpdateTeachByIDRequest) (*userV1.UpdateTeachByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	record := &model.Teach{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDTeach.Err()
	}
	record.ID = req.Id

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("teach", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.UpdateTeachByIDReply{}, nil
}

// GetByID get a record by id
func (s *teach) GetByID(ctx context.Context, req *userV1.GetTeachByIDRequest) (*userV1.GetTeachByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
	record, err := s.iDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, query.ErrNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertTeach(record)
	if err != nil {
		logger.Warn("convertTeach error", logger.Err(err), logger.Any("teach", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDTeach.Err()
	}

	return &userV1.GetTeachByIDReply{Teach: data}, nil
}

// GetByCondition get a record by id
func (s *teach) GetByCondition(ctx context.Context, req *userV1.GetTeachByConditionRequest) (*userV1.GetTeachByConditionReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
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

	data, err := convertTeach(record)
	if err != nil {
		logger.Warn("convertTeach error", logger.Err(err), logger.Any("teach", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByConditionTeach.Err()
	}

	return &userV1.GetTeachByConditionReply{
		Teach: data,
	}, nil
}

// ListByIDs list of records by batch id
func (s *teach) ListByIDs(ctx context.Context, req *userV1.ListTeachByIDsRequest) (*userV1.ListTeachByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
	teachMap, err := s.iDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("GetByIDs error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	teachs := []*userV1.Teach{}
	for _, id := range req.Ids {
		if v, ok := teachMap[id]; ok {
			record, err := convertTeach(v)
			if err != nil {
				logger.Warn("convertTeach error", logger.Err(err), logger.Any("teach", v), interceptor.ServerCtxRequestIDField(ctx))
				return nil, ecode.StatusInternalServerError.ToRPCErr()
			}
			teachs = append(teachs, record)
		}
	}

	return &userV1.ListTeachByIDsReply{Teachs: teachs}, nil
}

// List of records by query parameters
func (s *teach) List(ctx context.Context, req *userV1.ListTeachRequest) (*userV1.ListTeachReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListTeach.Err()
	}
	params.Size = int(req.Params.Limit)

	ctx = context.WithValue(ctx, interceptor.ContextRequestIDKey, interceptor.ServerCtxRequestID(ctx)) //nolint
	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	teachs := []*userV1.Teach{}
	for _, record := range records {
		data, err := convertTeach(record)
		if err != nil {
			logger.Warn("convertTeach error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		teachs = append(teachs, data)
	}

	return &userV1.ListTeachReply{
		Total:  total,
		Teachs: teachs,
	}, nil
}

func convertTeach(record *model.Teach) (*userV1.Teach, error) {
	value := &userV1.Teach{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	value.Id = record.ID
	value.CreatedAt = record.CreatedAt.Unix()
	value.UpdatedAt = record.UpdatedAt.Unix()
	return value, nil
}
