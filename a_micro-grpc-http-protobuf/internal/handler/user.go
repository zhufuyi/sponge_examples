// Code generated by https://github.com/zhufuyi/sponge

package handler

import (
	"context"

	userV1 "user/api/user/v1"
	"user/internal/service"
)

var _ userV1.UserLogicer = (*userHandler)(nil)

type userHandler struct {
	server userV1.UserServer
}

// NewUserHandler create a handler
func NewUserHandler() userV1.UserLogicer {
	return &userHandler{
		server: service.NewUserServer(),
	}
}

// Register 注册
func (h *userHandler) Register(ctx context.Context, req *userV1.RegisterRequest) (*userV1.RegisterReply, error) {
	return h.server.Register(ctx, req)
}

// Login 登录
func (h *userHandler) Login(ctx context.Context, req *userV1.LoginRequest) (*userV1.LoginReply, error) {
	return h.server.Login(ctx, req)
}

// Logout 登出
func (h *userHandler) Logout(ctx context.Context, req *userV1.LogoutRequest) (*userV1.LogoutReply, error) {
	return h.server.Logout(ctx, req)
}

// ChangePassword 修改密码
func (h *userHandler) ChangePassword(ctx context.Context, req *userV1.ChangePasswordRequest) (*userV1.ChangeRegisterReply, error) {
	return h.server.ChangePassword(ctx, req)
}
