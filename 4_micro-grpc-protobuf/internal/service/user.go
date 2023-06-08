// Code generated by https://github.com/zhufuyi/sponge

package service

import (
	"context"
	"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	"github.com/zhufuyi/sponge/pkg/logger"
	"user/internal/ecode"

	userV1 "user/api/user/v1"

	//"user/internal/cache"
	//"user/internal/dao"
	//"user/internal/ecode"
	//"user/internal/model"

	//"github.com/zhufuyi/sponge/pkg/grpc/interceptor"
	//"github.com/zhufuyi/sponge/pkg/logger"

	"google.golang.org/grpc"
)

func init() {
	registerFns = append(registerFns, func(server *grpc.Server) {
		userV1.RegisterUserServer(server, NewUserServer())
	})
}

var _ userV1.UserServer = (*user)(nil)

type user struct {
	userV1.UnimplementedUserServer

	// example:
	//		iDao dao.UserDao
}

// NewUserServer create a server
func NewUserServer() userV1.UserServer {
	return &user{
		// example:
		//		iDao: dao.NewUserDao(
		//			model.GetDB(),
		//			cache.NewUserCache(model.GetCacheType()),
		//		),
	}
}

// Register ......
func (s *user) Register(ctx context.Context, req *userV1.RegisterRequest) (*userV1.RegisterReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	// fill in the business logic code here

	return &userV1.RegisterReply{Id: 1}, nil
}

// Login ......
func (s *user) Login(ctx context.Context, req *userV1.LoginRequest) (*userV1.LoginReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	// fill in the business logic code here

	return &userV1.LoginReply{
		Id:    1,
		Token: "abcd......",
	}, nil
}

// Logout ......
func (s *user) Logout(ctx context.Context, req *userV1.LogoutRequest) (*userV1.LogoutReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	// fill in the business logic code here

	return &userV1.LogoutReply{}, nil
}

// ChangePassword ......
func (s *user) ChangePassword(ctx context.Context, req *userV1.ChangePasswordRequest) (*userV1.ChangeRegisterReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), interceptor.ServerCtxRequestIDField(ctx))
		return nil, ecode.StatusInvalidParams.Err()
	}

	// fill in the business logic code here

	return &userV1.ChangeRegisterReply{}, nil
}