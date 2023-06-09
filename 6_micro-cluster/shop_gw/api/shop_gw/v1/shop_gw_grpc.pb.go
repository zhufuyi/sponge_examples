// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api/shop_gw/v1/shop_gw.proto

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

// ShopGwClient is the client API for ShopGw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopGwClient interface {
	// get page detail by product id
	GetDetailsByProductID(ctx context.Context, in *GetDetailsByProductIDRequest, opts ...grpc.CallOption) (*GetDetailsByProductIDReply, error)
}

type shopGwClient struct {
	cc grpc.ClientConnInterface
}

func NewShopGwClient(cc grpc.ClientConnInterface) ShopGwClient {
	return &shopGwClient{cc}
}

func (c *shopGwClient) GetDetailsByProductID(ctx context.Context, in *GetDetailsByProductIDRequest, opts ...grpc.CallOption) (*GetDetailsByProductIDReply, error) {
	out := new(GetDetailsByProductIDReply)
	err := c.cc.Invoke(ctx, "/api.shop_gw.v1.ShopGw/GetDetailsByProductID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopGwServer is the server API for ShopGw service.
// All implementations must embed UnimplementedShopGwServer
// for forward compatibility
type ShopGwServer interface {
	// get page detail by product id
	GetDetailsByProductID(context.Context, *GetDetailsByProductIDRequest) (*GetDetailsByProductIDReply, error)
	mustEmbedUnimplementedShopGwServer()
}

// UnimplementedShopGwServer must be embedded to have forward compatible implementations.
type UnimplementedShopGwServer struct {
}

func (UnimplementedShopGwServer) GetDetailsByProductID(context.Context, *GetDetailsByProductIDRequest) (*GetDetailsByProductIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailsByProductID not implemented")
}
func (UnimplementedShopGwServer) mustEmbedUnimplementedShopGwServer() {}

// UnsafeShopGwServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopGwServer will
// result in compilation errors.
type UnsafeShopGwServer interface {
	mustEmbedUnimplementedShopGwServer()
}

func RegisterShopGwServer(s grpc.ServiceRegistrar, srv ShopGwServer) {
	s.RegisterService(&ShopGw_ServiceDesc, srv)
}

func _ShopGw_GetDetailsByProductID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailsByProductIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopGwServer).GetDetailsByProductID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.shop_gw.v1.ShopGw/GetDetailsByProductID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopGwServer).GetDetailsByProductID(ctx, req.(*GetDetailsByProductIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShopGw_ServiceDesc is the grpc.ServiceDesc for ShopGw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShopGw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.shop_gw.v1.ShopGw",
	HandlerType: (*ShopGwServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetailsByProductID",
			Handler:    _ShopGw_GetDetailsByProductID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/shop_gw/v1/shop_gw.proto",
}
