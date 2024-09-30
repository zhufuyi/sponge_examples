// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.2
// source: api/eshop_gw/v1/eshop_gw.proto

package v1

import (
	v13 "eshop_gw/api/coupon/v1"
	v15 "eshop_gw/api/flashSale/v1"
	v12 "eshop_gw/api/order/v1"
	v11 "eshop_gw/api/product/v1"
	v14 "eshop_gw/api/stock/v1"
	v1 "eshop_gw/api/user/v1"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PageParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page number, start from 0
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page" form:"page"`
	// Size per page
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit" form:"limit"`
}

func (x *PageParam) Reset() {
	*x = PageParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_eshop_gw_v1_eshop_gw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageParam) ProtoMessage() {}

func (x *PageParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_eshop_gw_v1_eshop_gw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageParam.ProtoReflect.Descriptor instead.
func (*PageParam) Descriptor() ([]byte, []int) {
	return file_api_eshop_gw_v1_eshop_gw_proto_rawDescGZIP(), []int{0}
}

func (x *PageParam) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageParam) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_api_eshop_gw_v1_eshop_gw_proto protoreflect.FileDescriptor

var file_api_eshop_gw_v1_eshop_gw_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2e, 0x76,
	0x31, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x6c, 0x61,
	0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x53,
	0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x18, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00,
	0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0xe8, 0x13, 0x0a, 0x05, 0x65, 0x73, 0x68,
	0x6f, 0x70, 0x12, 0xa8, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x5c, 0x92, 0x41, 0x42, 0x0a, 0x10, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe6, 0x9c, 0x8d, 0xe5,
	0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x75,
	0x73, 0x65, 0x72, 0x1a, 0x21, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xb7, 0x01,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x73, 0x92, 0x41, 0x57, 0x0a, 0x10, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe6,
	0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12, 0x1b, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x26, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xc0, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x65, 0x92, 0x41, 0x48, 0x0a, 0x10, 0xe5, 0x95, 0x86, 0xe5, 0x93, 0x81,
	0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x24, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74,
	0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0xc9, 0x01, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x7c, 0x92, 0x41, 0x5d, 0x0a, 0x10, 0xe5,
	0x95, 0x86, 0xe5, 0x93, 0x81, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12,
	0x1e, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a,
	0x29, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xb9, 0x01, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x68, 0x92, 0x41, 0x46, 0x0a, 0x10, 0xe8,
	0xae, 0xa2, 0xe5, 0x8d, 0x95, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12,
	0x0c, 0xe6, 0x8f, 0x90, 0xe4, 0xba, 0xa4, 0xe8, 0xae, 0xa2, 0xe5, 0x8d, 0x95, 0x1a, 0x24, 0xe6,
	0x8f, 0x90, 0xe4, 0xba, 0xa4, 0xe8, 0xae, 0xa2, 0xe5, 0x8d, 0x95, 0xef, 0xbc, 0x8c, 0x64, 0x74,
	0x6d, 0x20, 0x73, 0x61, 0x67, 0x61, 0x20, 0xe4, 0xba, 0x8b, 0xe5, 0x8a, 0xa1, 0xe6, 0xa8, 0xa1,
	0xe5, 0xbc, 0x8f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0xfb, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x8b, 0x01, 0x92, 0x41, 0x5f, 0x0a, 0x10, 0xe8, 0xae, 0xa2, 0xe5, 0x8d, 0x95,
	0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12, 0x1d, 0xe6, 0x8f, 0x90, 0xe4,
	0xba, 0xa4, 0xe8, 0xae, 0xa2, 0xe5, 0x8d, 0x95, 0x28, 0xe6, 0x9c, 0x89, 0xe9, 0x98, 0x9f, 0xe5,
	0x88, 0x97, 0xe7, 0xbc, 0x93, 0xe5, 0x86, 0xb2, 0x29, 0x1a, 0x2c, 0xe5, 0x8f, 0x91, 0xe9, 0x80,
	0x81, 0xe6, 0x8f, 0x90, 0xe4, 0xba, 0xa4, 0xe8, 0xae, 0xa2, 0xe5, 0x8d, 0x95, 0xe9, 0x80, 0x9a,
	0xe7, 0x9f, 0xa5, 0xe5, 0x88, 0xb0, 0xe6, 0xb6, 0x88, 0xe6, 0x81, 0xaf, 0xe9, 0x98, 0x9f, 0xe5,
	0x88, 0x97, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x6e,
	0x64, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x3a, 0x01, 0x2a,
	0x12, 0xbd, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x76, 0x92, 0x41, 0x59, 0x0a, 0x10, 0xe8,
	0xae, 0xa2, 0xe5, 0x8d, 0x95, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12,
	0x1c, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x20,
	0x62, 0x79, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x27, 0x6c,
	0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x20, 0x62, 0x79,
	0x20, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x12, 0xbb, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x65, 0x92, 0x41, 0x49, 0x0a, 0x13, 0xe4, 0xbc,
	0x98, 0xe6, 0x83, 0xa0, 0xe5, 0x88, 0xb8, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70,
	0x69, 0x12, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x1a, 0x23, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xc6,
	0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x7c, 0x92, 0x41, 0x5e, 0x0a, 0x13,
	0xe4, 0xbc, 0x98, 0xe6, 0x83, 0xa0, 0xe5, 0x88, 0xb8, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20,
	0x61, 0x70, 0x69, 0x12, 0x1d, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x1a, 0x28, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e,
	0x64, 0x20, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xb0, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5f, 0x92, 0x41, 0x44, 0x0a, 0x10,
	0xe5, 0xba, 0x93, 0xe5, 0xad, 0x98, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69,
	0x12, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x1a, 0x22,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0xbd, 0x01, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x76, 0x92, 0x41, 0x59, 0x0a, 0x10, 0xe5, 0xba, 0x93, 0xe5, 0xad, 0x98, 0xe6,
	0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12, 0x1c, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x27, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66,
	0x20, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x20, 0x62, 0x79, 0x20, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xd3, 0x01, 0x0a, 0x11, 0x53,
	0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x53, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x70,
	0x92, 0x41, 0x48, 0x0a, 0x10, 0xe5, 0xba, 0x93, 0xe5, 0xad, 0x98, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a,
	0xa1, 0x20, 0x61, 0x70, 0x69, 0x12, 0x15, 0xe8, 0xae, 0xbe, 0xe7, 0xbd, 0xae, 0xe4, 0xba, 0xa7,
	0xe5, 0x93, 0x81, 0xe7, 0x9a, 0x84, 0xe5, 0xba, 0x93, 0xe5, 0xad, 0x98, 0x1a, 0x1d, 0xe8, 0xae,
	0xbe, 0xe7, 0xbd, 0xae, 0xe4, 0xba, 0xa7, 0xe5, 0x93, 0x81, 0xe7, 0x9a, 0x84, 0xe5, 0xba, 0x93,
	0xe5, 0xad, 0x98, 0x28, 0xe5, 0xb9, 0x82, 0xe7, 0xad, 0x89, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x22, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2f, 0x73, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x3a, 0x01, 0x2a,
	0x12, 0xa0, 0x01, 0x0a, 0x09, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x4d, 0x92, 0x41, 0x2e, 0x0a, 0x10, 0xe7, 0xa7, 0x92, 0xe6, 0x9d,
	0x80, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x20, 0x61, 0x70, 0x69, 0x12, 0x0c, 0xe7, 0xa7, 0x92,
	0xe6, 0x9d, 0x80, 0xe6, 0x8a, 0xa2, 0xe8, 0xb4, 0xad, 0x1a, 0x0c, 0xe7, 0xa7, 0x92, 0xe6, 0x9d,
	0x80, 0xe6, 0x8a, 0xa2, 0xe8, 0xb4, 0xad, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x61, 0x6c, 0x65,
	0x3a, 0x01, 0x2a, 0x42, 0xc2, 0x01, 0x5a, 0x1b, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x77, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x92, 0x41, 0xa1, 0x01, 0x12, 0x1d, 0x0a, 0x16, 0x65, 0x73, 0x68, 0x6f, 0x70,
	0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x6f, 0x63,
	0x73, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x3a, 0x38, 0x30, 0x38, 0x30, 0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x48,
	0x0a, 0x46, 0x0a, 0x0a, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x38,
	0x08, 0x02, 0x12, 0x23, 0x54, 0x79, 0x70, 0x65, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20,
	0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a, 0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x74,
	0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_eshop_gw_v1_eshop_gw_proto_rawDescOnce sync.Once
	file_api_eshop_gw_v1_eshop_gw_proto_rawDescData = file_api_eshop_gw_v1_eshop_gw_proto_rawDesc
)

func file_api_eshop_gw_v1_eshop_gw_proto_rawDescGZIP() []byte {
	file_api_eshop_gw_v1_eshop_gw_proto_rawDescOnce.Do(func() {
		file_api_eshop_gw_v1_eshop_gw_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_eshop_gw_v1_eshop_gw_proto_rawDescData)
	})
	return file_api_eshop_gw_v1_eshop_gw_proto_rawDescData
}

var file_api_eshop_gw_v1_eshop_gw_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_eshop_gw_v1_eshop_gw_proto_goTypes = []interface{}{
	(*PageParam)(nil),                        // 0: api.eshop_gw.v1.PageParam
	(*v1.CreateUserRequest)(nil),             // 1: api.user.v1.CreateUserRequest
	(*v11.CreateProductRequest)(nil),         // 2: api.product.v1.CreateProductRequest
	(*v12.SubmitOrderRequest)(nil),           // 3: api.order.v1.SubmitOrderRequest
	(*v12.SendSubmitOrderNotifyRequest)(nil), // 4: api.order.v1.SendSubmitOrderNotifyRequest
	(*v13.CreateCouponRequest)(nil),          // 5: api.coupon.v1.CreateCouponRequest
	(*v14.CreateStockRequest)(nil),           // 6: api.stock.v1.CreateStockRequest
	(*v14.SetFlashSaleStockRequest)(nil),     // 7: api.stock.v1.SetFlashSaleStockRequest
	(*v15.FlashSaleRequest)(nil),             // 8: api.flashSale.v1.FlashSaleRequest
	(*v1.CreateUserReply)(nil),               // 9: api.user.v1.CreateUserReply
	(*v1.ListUserReply)(nil),                 // 10: api.user.v1.ListUserReply
	(*v11.CreateProductReply)(nil),           // 11: api.product.v1.CreateProductReply
	(*v11.ListProductReply)(nil),             // 12: api.product.v1.ListProductReply
	(*v12.SubmitOrderReply)(nil),             // 13: api.order.v1.SubmitOrderReply
	(*v12.SendSubmitOrderNotifyReply)(nil),   // 14: api.order.v1.SendSubmitOrderNotifyReply
	(*v12.ListOrderReply)(nil),               // 15: api.order.v1.ListOrderReply
	(*v13.CreateCouponReply)(nil),            // 16: api.coupon.v1.CreateCouponReply
	(*v13.ListCouponReply)(nil),              // 17: api.coupon.v1.ListCouponReply
	(*v14.CreateStockReply)(nil),             // 18: api.stock.v1.CreateStockReply
	(*v14.ListStockReply)(nil),               // 19: api.stock.v1.ListStockReply
	(*v14.SetFlashSaleStockReply)(nil),       // 20: api.stock.v1.SetFlashSaleStockReply
	(*v15.FlashSaleReply)(nil),               // 21: api.flashSale.v1.FlashSaleReply
}
var file_api_eshop_gw_v1_eshop_gw_proto_depIdxs = []int32{
	1,  // 0: api.eshop_gw.v1.eshop.CreateUser:input_type -> api.user.v1.CreateUserRequest
	0,  // 1: api.eshop_gw.v1.eshop.ListUser:input_type -> api.eshop_gw.v1.PageParam
	2,  // 2: api.eshop_gw.v1.eshop.CreateProduct:input_type -> api.product.v1.CreateProductRequest
	0,  // 3: api.eshop_gw.v1.eshop.ListProduct:input_type -> api.eshop_gw.v1.PageParam
	3,  // 4: api.eshop_gw.v1.eshop.SubmitOrder:input_type -> api.order.v1.SubmitOrderRequest
	4,  // 5: api.eshop_gw.v1.eshop.SendSubmitOrderNotify:input_type -> api.order.v1.SendSubmitOrderNotifyRequest
	0,  // 6: api.eshop_gw.v1.eshop.ListOrder:input_type -> api.eshop_gw.v1.PageParam
	5,  // 7: api.eshop_gw.v1.eshop.CreateCoupon:input_type -> api.coupon.v1.CreateCouponRequest
	0,  // 8: api.eshop_gw.v1.eshop.ListCoupon:input_type -> api.eshop_gw.v1.PageParam
	6,  // 9: api.eshop_gw.v1.eshop.CreateStock:input_type -> api.stock.v1.CreateStockRequest
	0,  // 10: api.eshop_gw.v1.eshop.ListStock:input_type -> api.eshop_gw.v1.PageParam
	7,  // 11: api.eshop_gw.v1.eshop.SetFlashSaleStock:input_type -> api.stock.v1.SetFlashSaleStockRequest
	8,  // 12: api.eshop_gw.v1.eshop.FlashSale:input_type -> api.flashSale.v1.FlashSaleRequest
	9,  // 13: api.eshop_gw.v1.eshop.CreateUser:output_type -> api.user.v1.CreateUserReply
	10, // 14: api.eshop_gw.v1.eshop.ListUser:output_type -> api.user.v1.ListUserReply
	11, // 15: api.eshop_gw.v1.eshop.CreateProduct:output_type -> api.product.v1.CreateProductReply
	12, // 16: api.eshop_gw.v1.eshop.ListProduct:output_type -> api.product.v1.ListProductReply
	13, // 17: api.eshop_gw.v1.eshop.SubmitOrder:output_type -> api.order.v1.SubmitOrderReply
	14, // 18: api.eshop_gw.v1.eshop.SendSubmitOrderNotify:output_type -> api.order.v1.SendSubmitOrderNotifyReply
	15, // 19: api.eshop_gw.v1.eshop.ListOrder:output_type -> api.order.v1.ListOrderReply
	16, // 20: api.eshop_gw.v1.eshop.CreateCoupon:output_type -> api.coupon.v1.CreateCouponReply
	17, // 21: api.eshop_gw.v1.eshop.ListCoupon:output_type -> api.coupon.v1.ListCouponReply
	18, // 22: api.eshop_gw.v1.eshop.CreateStock:output_type -> api.stock.v1.CreateStockReply
	19, // 23: api.eshop_gw.v1.eshop.ListStock:output_type -> api.stock.v1.ListStockReply
	20, // 24: api.eshop_gw.v1.eshop.SetFlashSaleStock:output_type -> api.stock.v1.SetFlashSaleStockReply
	21, // 25: api.eshop_gw.v1.eshop.FlashSale:output_type -> api.flashSale.v1.FlashSaleReply
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_eshop_gw_v1_eshop_gw_proto_init() }
func file_api_eshop_gw_v1_eshop_gw_proto_init() {
	if File_api_eshop_gw_v1_eshop_gw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_eshop_gw_v1_eshop_gw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_eshop_gw_v1_eshop_gw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_eshop_gw_v1_eshop_gw_proto_goTypes,
		DependencyIndexes: file_api_eshop_gw_v1_eshop_gw_proto_depIdxs,
		MessageInfos:      file_api_eshop_gw_v1_eshop_gw_proto_msgTypes,
	}.Build()
	File_api_eshop_gw_v1_eshop_gw_proto = out.File
	file_api_eshop_gw_v1_eshop_gw_proto_rawDesc = nil
	file_api_eshop_gw_v1_eshop_gw_proto_goTypes = nil
	file_api_eshop_gw_v1_eshop_gw_proto_depIdxs = nil
}
