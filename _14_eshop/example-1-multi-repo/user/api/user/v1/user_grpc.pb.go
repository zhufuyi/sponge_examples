// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api/user/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	User_Create_FullMethodName     = "/api.user.v1.user/Create"
	User_DeleteByID_FullMethodName = "/api.user.v1.user/DeleteByID"
	User_UpdateByID_FullMethodName = "/api.user.v1.user/UpdateByID"
	User_GetByID_FullMethodName    = "/api.user.v1.user/GetByID"
	User_List_FullMethodName       = "/api.user.v1.user/List"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// create user
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error)
	// delete user by id
	DeleteByID(ctx context.Context, in *DeleteUserByIDRequest, opts ...grpc.CallOption) (*DeleteUserByIDReply, error)
	// update user by id
	UpdateByID(ctx context.Context, in *UpdateUserByIDRequest, opts ...grpc.CallOption) (*UpdateUserByIDReply, error)
	// get user by id
	GetByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDReply, error)
	// list of user by query parameters
	List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserReply, error) {
	out := new(CreateUserReply)
	err := c.cc.Invoke(ctx, User_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteByID(ctx context.Context, in *DeleteUserByIDRequest, opts ...grpc.CallOption) (*DeleteUserByIDReply, error) {
	out := new(DeleteUserByIDReply)
	err := c.cc.Invoke(ctx, User_DeleteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateByID(ctx context.Context, in *UpdateUserByIDRequest, opts ...grpc.CallOption) (*UpdateUserByIDReply, error) {
	out := new(UpdateUserByIDReply)
	err := c.cc.Invoke(ctx, User_UpdateByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetByID(ctx context.Context, in *GetUserByIDRequest, opts ...grpc.CallOption) (*GetUserByIDReply, error) {
	out := new(GetUserByIDReply)
	err := c.cc.Invoke(ctx, User_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserReply, error) {
	out := new(ListUserReply)
	err := c.cc.Invoke(ctx, User_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// create user
	Create(context.Context, *CreateUserRequest) (*CreateUserReply, error)
	// delete user by id
	DeleteByID(context.Context, *DeleteUserByIDRequest) (*DeleteUserByIDReply, error)
	// update user by id
	UpdateByID(context.Context, *UpdateUserByIDRequest) (*UpdateUserByIDReply, error)
	// get user by id
	GetByID(context.Context, *GetUserByIDRequest) (*GetUserByIDReply, error)
	// list of user by query parameters
	List(context.Context, *ListUserRequest) (*ListUserReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Create(context.Context, *CreateUserRequest) (*CreateUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServer) DeleteByID(context.Context, *DeleteUserByIDRequest) (*DeleteUserByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedUserServer) UpdateByID(context.Context, *UpdateUserByIDRequest) (*UpdateUserByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByID not implemented")
}
func (UnimplementedUserServer) GetByID(context.Context, *GetUserByIDRequest) (*GetUserByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedUserServer) List(context.Context, *ListUserRequest) (*ListUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteByID(ctx, req.(*DeleteUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UpdateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateByID(ctx, req.(*UpdateUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetByID(ctx, req.(*GetUserByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).List(ctx, req.(*ListUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _User_Create_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _User_DeleteByID_Handler,
		},
		{
			MethodName: "UpdateByID",
			Handler:    _User_UpdateByID_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _User_GetByID_Handler,
		},
		{
			MethodName: "List",
			Handler:    _User_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/user/v1/user.proto",
}
