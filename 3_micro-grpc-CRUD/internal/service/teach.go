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
		userV1.RegisterTeachServiceServer(server, NewTeachServiceServer()) // register service to the rpc service
	})
}

var _ userV1.TeachServiceServer = (*teachService)(nil)

type teachService struct {
	userV1.UnimplementedTeachServiceServer

	iDao dao.TeachDao
}

// NewTeachServiceServer create a new service
func NewTeachServiceServer() userV1.TeachServiceServer {
	return &teachService{
		iDao: dao.NewTeachDao(
			model.GetDB(),
			cache.NewTeachCache(model.GetCacheType()),
		),
	}
}

// Create a record
func (s *teachService) Create(ctx context.Context, req *userV1.CreateTeachRequest) (*userV1.CreateTeachReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	teach := &model.Teach{}
	err = copier.Copy(teach, req)
	if err != nil {
		return nil, ecode.StatusCreateTeach.Err()
	}

	err = s.iDao.Create(ctx, teach)
	if err != nil {
		logger.Error("s.iDao.Create error", logger.Err(err), logger.Any("teach", teach), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.CreateTeachReply{Id: teach.ID}, nil
}

// DeleteByID delete a record by id
func (s *teachService) DeleteByID(ctx context.Context, req *userV1.DeleteTeachByIDRequest) (*userV1.DeleteTeachByIDReply, error) {
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

	return &userV1.DeleteTeachByIDReply{}, nil
}

// DeleteByIDs delete records by batch id
func (s *teachService) DeleteByIDs(ctx context.Context, req *userV1.DeleteTeachByIDsRequest) (*userV1.DeleteTeachByIDsReply, error) {
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

	return &userV1.DeleteTeachByIDsReply{}, nil
}

// UpdateByID update a record by id
func (s *teachService) UpdateByID(ctx context.Context, req *userV1.UpdateTeachByIDRequest) (*userV1.UpdateTeachByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	teach := &model.Teach{}
	err = copier.Copy(teach, req)
	if err != nil {
		return nil, ecode.StatusUpdateTeach.Err()
	}
	teach.ID = req.Id

	err = s.iDao.UpdateByID(ctx, teach)
	if err != nil {
		logger.Error("s.iDao.UpdateByID error", logger.Err(err), logger.Any("teach", teach), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.UpdateTeachByIDReply{}, nil
}

// GetByID get a record by id
func (s *teachService) GetByID(ctx context.Context, req *userV1.GetTeachByIDRequest) (*userV1.GetTeachByIDReply, error) {
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

	data, err := convertTeach(record)
	if err != nil {
		logger.Warn("convertTeach error", logger.Err(err), logger.Any("record", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetTeach.Err()
	}

	return &userV1.GetTeachByIDReply{Teach: data}, nil
}

// ListByIDs list of records by batch id
func (s *teachService) ListByIDs(ctx context.Context, req *userV1.ListTeachByIDsRequest) (*userV1.ListTeachByIDsReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	teachMap, err := s.iDao.GetByIDs(ctx, req.Ids)
	if err != nil {
		logger.Error("s.iDao.GetByID error", logger.Err(err), logger.Any("ids", req.Ids), interceptor.ServerCtxRequestIDField(ctx))
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
func (s *teachService) List(ctx context.Context, req *userV1.ListTeachRequest) (*userV1.ListTeachReply, error) {
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

	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("s.iDao.GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("s.iDao.GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	teachs := []*userV1.Teach{}
	for _, record := range records {
		teach, err := convertTeach(record)
		if err != nil {
			logger.Warn("convertTeach error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		teachs = append(teachs, teach)
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
