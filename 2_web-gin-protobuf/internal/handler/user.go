// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"

	userV1 "user/api/user/v1"
	"user/internal/ecode"

	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/logger"
)

var _ userV1.UserLogicer = (*userHandler)(nil)

type userHandler struct {
	// example:
	// 	userDao dao.UserDao
}

// NewUserHandler create a handler
func NewUserHandler() userV1.UserLogicer {
	return &userHandler{
		// example:
		// 	userDao: dao.NewUserDao(
		// 		model.GetDB(),
		// 		cache.NewUserCache(model.GetCacheType()),
		// 	),
	}
}

// Register 注册
func (h *userHandler) Register(ctx context.Context, req *userV1.RegisterRequest) (*userV1.RegisterReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	logger.Info("register successfully", middleware.CtxRequestIDField(ctx))

	return &userV1.RegisterReply{
		Id: 100,
	}, nil
}

// Login 登录
func (h *userHandler) Login(ctx context.Context, req *userV1.LoginRequest) (*userV1.LoginReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	logger.Info("login successfully", middleware.CtxRequestIDField(ctx))

	return &userV1.LoginReply{
		Id:    100,
		Token: "eydiewnafiaekdfaf......",
	}, nil
}

// Logout 登出
func (h *userHandler) Logout(ctx context.Context, req *userV1.LogoutRequest) (*userV1.LogoutReply, error) {
	err := req.Validate()
	if err != nil {
		logger.Warn("req.Validate error", logger.Err(err), logger.Any("req", req), middleware.CtxRequestIDField(ctx))
		return nil, ecode.InvalidParams.Err()
	}

	logger.Info("logout successfully", middleware.CtxRequestIDField(ctx))

	return &userV1.LogoutReply{}, nil
}