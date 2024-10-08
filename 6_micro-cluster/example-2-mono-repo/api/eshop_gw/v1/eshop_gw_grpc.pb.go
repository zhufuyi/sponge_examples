// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: api/eshop_gw/v1/eshop_gw.proto

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
	EShopGw_GetDetailsByProductID_FullMethodName = "/api.eshop_gw.v1.EShopGw/GetDetailsByProductID"
)

// EShopGwClient is the client API for EShopGw service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EShopGwClient interface {
	// get page detail by product id
	GetDetailsByProductID(ctx context.Context, in *GetDetailsByProductIDRequest, opts ...grpc.CallOption) (*GetDetailsByProductIDReply, error)
}

type eShopGwClient struct {
	cc grpc.ClientConnInterface
}

func NewEShopGwClient(cc grpc.ClientConnInterface) EShopGwClient {
	return &eShopGwClient{cc}
}

func (c *eShopGwClient) GetDetailsByProductID(ctx context.Context, in *GetDetailsByProductIDRequest, opts ...grpc.CallOption) (*GetDetailsByProductIDReply, error) {
	out := new(GetDetailsByProductIDReply)
	err := c.cc.Invoke(ctx, EShopGw_GetDetailsByProductID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EShopGwServer is the server API for EShopGw service.
// All implementations must embed UnimplementedEShopGwServer
// for forward compatibility
type EShopGwServer interface {
	// get page detail by product id
	GetDetailsByProductID(context.Context, *GetDetailsByProductIDRequest) (*GetDetailsByProductIDReply, error)
	mustEmbedUnimplementedEShopGwServer()
}

// UnimplementedEShopGwServer must be embedded to have forward compatible implementations.
type UnimplementedEShopGwServer struct {
}

func (UnimplementedEShopGwServer) GetDetailsByProductID(context.Context, *GetDetailsByProductIDRequest) (*GetDetailsByProductIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetailsByProductID not implemented")
}
func (UnimplementedEShopGwServer) mustEmbedUnimplementedEShopGwServer() {}

// UnsafeEShopGwServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EShopGwServer will
// result in compilation errors.
type UnsafeEShopGwServer interface {
	mustEmbedUnimplementedEShopGwServer()
}

func RegisterEShopGwServer(s grpc.ServiceRegistrar, srv EShopGwServer) {
	s.RegisterService(&EShopGw_ServiceDesc, srv)
}

func _EShopGw_GetDetailsByProductID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailsByProductIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EShopGwServer).GetDetailsByProductID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EShopGw_GetDetailsByProductID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EShopGwServer).GetDetailsByProductID(ctx, req.(*GetDetailsByProductIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EShopGw_ServiceDesc is the grpc.ServiceDesc for EShopGw service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EShopGw_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.eshop_gw.v1.EShopGw",
	HandlerType: (*EShopGwServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetailsByProductID",
			Handler:    _EShopGw_GetDetailsByProductID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/eshop_gw/v1/eshop_gw.proto",
}
