// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api/product/v1/product.proto

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
	Product_Create_FullMethodName     = "/api.product.v1.product/Create"
	Product_DeleteByID_FullMethodName = "/api.product.v1.product/DeleteByID"
	Product_UpdateByID_FullMethodName = "/api.product.v1.product/UpdateByID"
	Product_GetByID_FullMethodName    = "/api.product.v1.product/GetByID"
	Product_List_FullMethodName       = "/api.product.v1.product/List"
)

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductClient interface {
	// create product
	Create(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductReply, error)
	// delete product by id
	DeleteByID(ctx context.Context, in *DeleteProductByIDRequest, opts ...grpc.CallOption) (*DeleteProductByIDReply, error)
	// update product by id
	UpdateByID(ctx context.Context, in *UpdateProductByIDRequest, opts ...grpc.CallOption) (*UpdateProductByIDReply, error)
	// get product by id
	GetByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDReply, error)
	// list of product by query parameters
	List(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductReply, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) Create(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductReply, error) {
	out := new(CreateProductReply)
	err := c.cc.Invoke(ctx, Product_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) DeleteByID(ctx context.Context, in *DeleteProductByIDRequest, opts ...grpc.CallOption) (*DeleteProductByIDReply, error) {
	out := new(DeleteProductByIDReply)
	err := c.cc.Invoke(ctx, Product_DeleteByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) UpdateByID(ctx context.Context, in *UpdateProductByIDRequest, opts ...grpc.CallOption) (*UpdateProductByIDReply, error) {
	out := new(UpdateProductByIDReply)
	err := c.cc.Invoke(ctx, Product_UpdateByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) GetByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*GetProductByIDReply, error) {
	out := new(GetProductByIDReply)
	err := c.cc.Invoke(ctx, Product_GetByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) List(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductReply, error) {
	out := new(ListProductReply)
	err := c.cc.Invoke(ctx, Product_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServer is the server API for Product service.
// All implementations must embed UnimplementedProductServer
// for forward compatibility
type ProductServer interface {
	// create product
	Create(context.Context, *CreateProductRequest) (*CreateProductReply, error)
	// delete product by id
	DeleteByID(context.Context, *DeleteProductByIDRequest) (*DeleteProductByIDReply, error)
	// update product by id
	UpdateByID(context.Context, *UpdateProductByIDRequest) (*UpdateProductByIDReply, error)
	// get product by id
	GetByID(context.Context, *GetProductByIDRequest) (*GetProductByIDReply, error)
	// list of product by query parameters
	List(context.Context, *ListProductRequest) (*ListProductReply, error)
	mustEmbedUnimplementedProductServer()
}

// UnimplementedProductServer must be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (UnimplementedProductServer) Create(context.Context, *CreateProductRequest) (*CreateProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProductServer) DeleteByID(context.Context, *DeleteProductByIDRequest) (*DeleteProductByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByID not implemented")
}
func (UnimplementedProductServer) UpdateByID(context.Context, *UpdateProductByIDRequest) (*UpdateProductByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByID not implemented")
}
func (UnimplementedProductServer) GetByID(context.Context, *GetProductByIDRequest) (*GetProductByIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedProductServer) List(context.Context, *ListProductRequest) (*ListProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProductServer) mustEmbedUnimplementedProductServer() {}

// UnsafeProductServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServer will
// result in compilation errors.
type UnsafeProductServer interface {
	mustEmbedUnimplementedProductServer()
}

func RegisterProductServer(s grpc.ServiceRegistrar, srv ProductServer) {
	s.RegisterService(&Product_ServiceDesc, srv)
}

func _Product_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).Create(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_DeleteByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).DeleteByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_DeleteByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).DeleteByID(ctx, req.(*DeleteProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_UpdateByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).UpdateByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_UpdateByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).UpdateByID(ctx, req.(*UpdateProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_GetByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).GetByID(ctx, req.(*GetProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Product_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).List(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Product_ServiceDesc is the grpc.ServiceDesc for Product service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Product_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.product.v1.product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Product_Create_Handler,
		},
		{
			MethodName: "DeleteByID",
			Handler:    _Product_DeleteByID_Handler,
		},
		{
			MethodName: "UpdateByID",
			Handler:    _Product_UpdateByID_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _Product_GetByID_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Product_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/product/v1/product.proto",
}
