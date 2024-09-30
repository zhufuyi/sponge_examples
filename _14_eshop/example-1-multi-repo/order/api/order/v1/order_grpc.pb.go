// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api/order/v1/order.proto

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
	Order_Submit_FullMethodName                = "/api.order.v1.order/Submit"
	Order_SendSubmitOrderNotify_FullMethodName = "/api.order.v1.order/SendSubmitOrderNotify"
	Order_Create_FullMethodName                = "/api.order.v1.order/Create"
	Order_CreateRevert_FullMethodName          = "/api.order.v1.order/CreateRevert"
	Order_DeleteByID_FullMethodName            = "/api.order.v1.order/DeleteByID"
	Order_UpdateByID_FullMethodName            = "/api.order.v1.order/UpdateByID"
	Order_GetByID_FullMethodName               = "/api.order.v1.order/GetByID"
	Order_List_FullMethodName                  = "/api.order.v1.order/List"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	// 提交订单(分布式事务)
	Submit(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderReply, error)
	// 发送提交订单通知
	SendSubmitOrderNotify(ctx context.Context, in *SendSubmitOrderNotifyRequest, opts ...grpc.CallOption) (*SendSubmitOrderNotifyReply, error)
	// 创建订单
	Create(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error)
	// 取消创建订单
	CreateRevert(ctx context.Context, in *CreateOrderRevertRequest, opts ...grpc.CallOption) (*CreateOrderRevertReply, error)
	// delete order by id
	DeleteByID(ctx context.Context, in *DeleteOrderByIDRequest, opts ...grpc.CallOption) (*DeleteOrderByIDReply, error)
	// update order by id
	UpdateByID(ctx context.Context, in *UpdateOrderByIDRequest, opts ...grpc.CallOption) (*UpdateOrderByIDReply, error)
	// get order by id
	GetByID(ctx context.Context, in *GetOrderByIDRequest, opts ...grpc.CallOption) (*GetOrderByIDReply, error)
	// list of order by query parameters
	List(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) Submit(ctx context.Context, in *SubmitOrderRequest, opts ...grpc.CallOption) (*SubmitOrderReply, error) {
	out := new(SubmitOrderReply)
	err := c.cc.Invoke(ctx, Order_Submit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) SendSubmitOrderNotify(ctx context.Context, in *SendSubmitOrderNotifyRequest, opts ...grpc.CallOption) (*SendSubmitOrderNotifyReply, error) {
	out := new(SendSubmitOrderNotifyReply)
	err := c.cc.Invoke(ctx, Order_SendSubmitOrderNotify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Create(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error) {
	out := new(CreateOrderReply)
	err := c.cc.Invoke(ctx, Order_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateRevert(ctx context.Context, in *CreateOrderRevertRequest, opts ...grpc.CallOption) (*CreateOrderRevertReply, error) {
	out := new(CreateOrderRevertReply)
	err := c.cc.Invoke(ctx, Order_CreateRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteByID(ctx context.Context, in *DeleteOrderByIDRequest, opts ...grpc.CallOption) (*DeleteOrderByIDReply, error) {
	out := new(DeleteOrderByIDReply)
	err := c.cc.Invoke(ctx, Order_DeleteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateByID(ctx context.Context, in *UpdateOrderByIDRequest, opts ...grpc.CallOption) (*UpdateOrderByIDReply, error) {
	out := new(UpdateOrderByIDReply)
	err := c.cc.Invoke(ctx, Order_UpdateByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetByID(ctx context.Context, in *GetOrderByIDRequest, opts ...grpc.CallOption) (*GetOrderByIDReply, error) {
	out := new(GetOrderByIDReply)
	err := c.cc.Invoke(ctx, Order_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) List(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error) {
	out := new(ListOrderReply)
	err := c.cc.Invoke(ctx, Order_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	// 提交订单(分布式事务)
	Submit(context.Context, *SubmitOrderRequest) (*SubmitOrderReply, error)
	// 发送提交订单通知
	SendSubmitOrderNotify(context.Context, *SendSubmitOrderNotifyRequest) (*SendSubmitOrderNotifyReply, error)
	// 创建订单
	Create(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
	// 取消创建订单
	CreateRevert(context.Context, *CreateOrderRevertRequest) (*CreateOrderRevertReply, error)
	// delete order by id
	DeleteByID(context.Context, *DeleteOrderByIDRequest) (*DeleteOrderByIDReply, error)
	// update order by id
	UpdateByID(context.Context, *UpdateOrderByIDRequest) (*UpdateOrderByIDReply, error)
	// get order by id
	GetByID(context.Context, *GetOrderByIDRequest) (*GetOrderByIDReply, error)
	// list of order by query parameters
	List(context.Context, *ListOrderRequest) (*ListOrderReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) Submit(context.Context, *SubmitOrderRequest) (*SubmitOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedOrderServer) SendSubmitOrderNotify(context.Context, *SendSubmitOrderNotifyRequest) (*SendSubmitOrderNotifyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSubmitOrderNotify not implemented")
}
func (UnimplementedOrderServer) Create(context.Context, *CreateOrderRequest) (*CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderServer) CreateRevert(context.Context, *CreateOrderRevertRequest) (*CreateOrderRevertReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRevert not implemented")
}
func (UnimplementedOrderServer) DeleteByID(context.Context, *DeleteOrderByIDRequest) (*DeleteOrderByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedOrderServer) UpdateByID(context.Context, *UpdateOrderByIDRequest) (*UpdateOrderByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByID not implemented")
}
func (UnimplementedOrderServer) GetByID(context.Context, *GetOrderByIDRequest) (*GetOrderByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedOrderServer) List(context.Context, *ListOrderRequest) (*ListOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Submit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Submit(ctx, req.(*SubmitOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_SendSubmitOrderNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSubmitOrderNotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).SendSubmitOrderNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_SendSubmitOrderNotify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).SendSubmitOrderNotify(ctx, req.(*SendSubmitOrderNotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Create(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRevertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateRevert(ctx, req.(*CreateOrderRevertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_DeleteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteByID(ctx, req.(*DeleteOrderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_UpdateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateByID(ctx, req.(*UpdateOrderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetByID(ctx, req.(*GetOrderByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).List(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.order.v1.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _Order_Submit_Handler,
		},
		{
			MethodName: "SendSubmitOrderNotify",
			Handler:    _Order_SendSubmitOrderNotify_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Order_Create_Handler,
		},
		{
			MethodName: "CreateRevert",
			Handler:    _Order_CreateRevert_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _Order_DeleteByID_Handler,
		},
		{
			MethodName: "UpdateByID",
			Handler:    _Order_UpdateByID_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Order_GetByID_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Order_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/order/v1/order.proto",
}
