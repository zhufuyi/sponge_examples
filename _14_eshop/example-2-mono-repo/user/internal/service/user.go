// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	"github.com/zhufuyi/sponge/pkg/ggorm/query"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"github.com/zhufuyi/sponge/pkg/utils"

	userV1 "eshop/api/user/v1"
	"eshop/user/internal/cache"
	"eshop/user/internal/dao"
	"eshop/user/internal/ecode"
	"eshop/user/internal/model"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		userV1.RegisterUserServer(server, NewUserServer())
	})
}

var _ userV1.UserServer = (*user)(nil)
var _ time.Time

type user struct {
	userV1.UnimplementedUserServer

	iDao dao.UserDao
}

// NewUserServer create a server
func NewUserServer() userV1.UserServer {
	return &user{
		iDao: dao.NewUserDao(
			model.GetDB(),
			cache.NewUserCache(model.GetCacheType()),
		),
	}
}

// Create user
func (s *user) Create(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.CreateUserReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.User{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusCreateUser.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here

	err = s.iDao.Create(ctx, record)
	if err != nil {
		logger.Error("Create error", logger.Err(err), logger.Any("user", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.CreateUserReply{Id: record.ID}, nil
}

// DeleteByID delete user by id
func (s *user) DeleteByID(ctx context.Context, req *userV1.DeleteUserByIDRequest) (*userV1.DeleteUserByIDReply, error) {
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

	return &userV1.DeleteUserByIDReply{}, nil
}

// UpdateByID update user by id
func (s *user) UpdateByID(ctx context.Context, req *userV1.UpdateUserByIDRequest) (*userV1.UpdateUserByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record := &model.User{}
	err = copier.Copy(record, req)
	if err != nil {
		return nil, ecode.StatusUpdateByIDUser.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	record.ID = req.Id

	err = s.iDao.UpdateByID(ctx, record)
	if err != nil {
		logger.Error("UpdateByID error", logger.Err(err), logger.Any("user", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	return &userV1.UpdateUserByIDReply{}, nil
}

// GetByID get user by id
func (s *user) GetByID(ctx context.Context, req *userV1.GetUserByIDRequest) (*userV1.GetUserByIDReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	record, err := s.iDao.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrRecordNotFound) {
			logger.Warn("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusNotFound.Err()
		}
		logger.Error("GetByID error", logger.Err(err), logger.Any("id", req.Id), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	data, err := convertUser(record)
	if err != nil {
		logger.Warn("convertUser error", logger.Err(err), logger.Any("user", record), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusGetByIDUser.Err()
	}

	return &userV1.GetUserByIDReply{User: data}, nil
}

// List of user by query parameters
func (s *user) List(ctx context.Context, req *userV1.ListUserRequest) (*userV1.ListUserReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}
	ctx = interceptor.WrapServerCtx(ctx)

	params := &query.Params{}
	err = copier.Copy(params, req.Params)
	if err != nil {
		return nil, ecode.StatusListUser.Err()
	}
	// Note: if copier.Copy cannot assign a value to a field, add it here
	params.Size = int(req.Params.Limit)

	records, total, err := s.iDao.GetByColumns(ctx, params)
	if err != nil {
		if strings.Contains(err.Error(), "query params error:") {
			logger.Warn("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
			return nil, ecode.StatusInvalidParams.Err()
		}
		logger.Error("GetByColumns error", logger.Err(err), logger.Any("params", params), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInternalServerError.ToRPCErr()
	}

	users := []*userV1.User{}
	for _, record := range records {
		data, err := convertUser(record)
		if err != nil {
			logger.Warn("convertUser error", logger.Err(err), logger.Any("id", record.ID), interceptor.ServerCtxRequestIDField(ctx))
			continue
		}
		users = append(users, data)
	}

	return &userV1.ListUserReply{
		Total: total,
		Users: users,
	}, nil
}

func convertUser(record *model.User) (*userV1.User, error) {
	value := &userV1.User{}
	err := copier.Copy(value, record)
	if err != nil {
		return nil, err
	}
	// todo if copier.Copy cannot assign a value to a field, add it here, e.g. CreatedAt, UpdatedAt
	value.Id = record.ID
	value.CreatedAt = utils.FormatDateTimeRFC3339(record.CreatedAt)
	value.UpdatedAt = utils.FormatDateTimeRFC3339(record.UpdatedAt)

	return value, nil
}
