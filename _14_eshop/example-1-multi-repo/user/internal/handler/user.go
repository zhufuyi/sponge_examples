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

// Create a record
func (h *userHandler) Create(ctx context.Context, req *userV1.CreateUserRequest) (*userV1.CreateUserReply, error) {
	return h.server.Create(ctx, req)
}

// DeleteByID delete a record by id
func (h *userHandler) DeleteByID(ctx context.Context, req *userV1.DeleteUserByIDRequest) (*userV1.DeleteUserByIDReply, error) {
	return h.server.DeleteByID(ctx, req)
}

// UpdateByID update a record by id
func (h *userHandler) UpdateByID(ctx context.Context, req *userV1.UpdateUserByIDRequest) (*userV1.UpdateUserByIDReply, error) {
	return h.server.UpdateByID(ctx, req)
}

// GetByID get a record by id
func (h *userHandler) GetByID(ctx context.Context, req *userV1.GetUserByIDRequest) (*userV1.GetUserByIDReply, error) {
	return h.server.GetByID(ctx, req)
}

// List of records by query parameters
func (h *userHandler) List(ctx context.Context, req *userV1.ListUserRequest) (*userV1.ListUserReply, error) {
	return h.server.List(ctx, req)
}
