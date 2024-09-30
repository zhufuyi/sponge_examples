// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.2
// source: api/coupon/v1/coupon.proto

package v1

import (
	
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

type CouponUseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponID uint64 `protobuf:"varint,1,opt,name=couponID,proto3" json:"couponID"`
}

func (x *CouponUseRequest) Reset() {
	*x = CouponUseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_coupon_v1_coupon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponUseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponUseRequest) ProtoMessage() {}

func (x *CouponUseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_coupon_v1_coupon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponUseRequest.ProtoReflect.Descriptor instead.
func (*CouponUseRequest) Descriptor() ([]byte, []int) {
	return file_api_coupon_v1_coupon_proto_rawDescGZIP(), []int{0}
}

func (x *CouponUseRequest) GetCouponID() uint64 {
	if x != nil {
		return x.CouponID
	}
	return 0
}

type CouponUseReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CouponUseReply) Reset() {
	*x = CouponUseReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_coupon_v1_coupon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponUseReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponUseReply) ProtoMessage() {}

func (x *CouponUseReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_coupon_v1_coupon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponUseReply.ProtoReflect.Descriptor instead.
func (*CouponUseReply) Descriptor() ([]byte, []int) {
	return file_api_coupon_v1_coupon_proto_rawDescGZIP(), []int{1}
}

type CouponUseRevertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponID uint64 `protobuf:"varint,1,opt,name=couponID,proto3" json:"couponID"`
}

func (x *CouponUseRevertRequest) Reset() {
	*x = CouponUseRevertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_coupon_v1_coupon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponUseRevertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponUseRevertRequest) ProtoMessage() {}

func (x *CouponUseRevertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_coupon_v1_coupon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponUseRevertRequest.ProtoReflect.Descriptor instead.
func (*CouponUseRevertRequest) Descriptor() ([]byte, []int) {
	return file_api_coupon_v1_coupon_proto_rawDescGZIP(), []int{2}
}

func (x *CouponUseRevertRequest) GetCouponID() uint64 {
	if x != nil {
		return x.CouponID
	}
	return 0
}

type CouponUseRevertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CouponUseRevertReply) Reset() {
	*x = CouponUseRevertReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_coupon_v1_coupon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponUseRevertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponUseRevertReply) ProtoMessage() {}

func (x *CouponUseRevertReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_coupon_v1_coupon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponUseRevertReply.ProtoReflect.Descriptor instead.
func (*CouponUseRevertReply) Descriptor() ([]byte, []int) {
	return file_api_coupon_v1_coupon_proto_rawDescGZIP(), []int{3}
}

var File_api_coupon_v1_coupon_proto protoreflect.FileDescriptor

var file_api_coupon_v1_coupon_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x10, 0x0a,
	0x0e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x3d, 0x0a, 0x16, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x76, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x16,
	0x0a, 0x14, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x76, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xb8, 0x01, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x12, 0x4d, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x12, 0x1f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x76,
	0x65, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x76,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x18, 0x5a, 0x16, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_coupon_v1_coupon_proto_rawDescOnce sync.Once
	file_api_coupon_v1_coupon_proto_rawDescData = file_api_coupon_v1_coupon_proto_rawDesc
)

func file_api_coupon_v1_coupon_proto_rawDescGZIP() []byte {
	file_api_coupon_v1_coupon_proto_rawDescOnce.Do(func() {
		file_api_coupon_v1_coupon_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_coupon_v1_coupon_proto_rawDescData)
	})
	return file_api_coupon_v1_coupon_proto_rawDescData
}

var file_api_coupon_v1_coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_coupon_v1_coupon_proto_goTypes = []interface{}{
	(*CouponUseRequest)(nil),       // 0: api.coupon.v1.CouponUseRequest
	(*CouponUseReply)(nil),         // 1: api.coupon.v1.CouponUseReply
	(*CouponUseRevertRequest)(nil), // 2: api.coupon.v1.couponUseRevertRequest
	(*CouponUseRevertReply)(nil),   // 3: api.coupon.v1.couponUseRevertReply
}
var file_api_coupon_v1_coupon_proto_depIdxs = []int32{
	0, // 0: api.coupon.v1.coupon.CouponUse:input_type -> api.coupon.v1.CouponUseRequest
	2, // 1: api.coupon.v1.coupon.CouponUseRevert:input_type -> api.coupon.v1.couponUseRevertRequest
	1, // 2: api.coupon.v1.coupon.CouponUse:output_type -> api.coupon.v1.CouponUseReply
	3, // 3: api.coupon.v1.coupon.CouponUseRevert:output_type -> api.coupon.v1.couponUseRevertReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_coupon_v1_coupon_proto_init() }
func file_api_coupon_v1_coupon_proto_init() {
	if File_api_coupon_v1_coupon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_coupon_v1_coupon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponUseRequest); i {
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
		file_api_coupon_v1_coupon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponUseReply); i {
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
		file_api_coupon_v1_coupon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponUseRevertRequest); i {
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
		file_api_coupon_v1_coupon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponUseRevertReply); i {
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
			RawDescriptor: file_api_coupon_v1_coupon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_coupon_v1_coupon_proto_goTypes,
		DependencyIndexes: file_api_coupon_v1_coupon_proto_depIdxs,
		MessageInfos:      file_api_coupon_v1_coupon_proto_msgTypes,
	}.Build()
	File_api_coupon_v1_coupon_proto = out.File
	file_api_coupon_v1_coupon_proto_rawDesc = nil
	file_api_coupon_v1_coupon_proto_goTypes = nil
	file_api_coupon_v1_coupon_proto_depIdxs = nil
}