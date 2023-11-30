// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/stock/v1/stock.proto

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

type StockDeductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId    uint64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId"`
	ProductCount uint32 `protobuf:"varint,2,opt,name=productCount,proto3" json:"productCount"`
}

func (x *StockDeductRequest) Reset() {
	*x = StockDeductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockDeductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDeductRequest) ProtoMessage() {}

func (x *StockDeductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDeductRequest.ProtoReflect.Descriptor instead.
func (*StockDeductRequest) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_stock_proto_rawDescGZIP(), []int{0}
}

func (x *StockDeductRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *StockDeductRequest) GetProductCount() uint32 {
	if x != nil {
		return x.ProductCount
	}
	return 0
}

type StockDeductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StockDeductReply) Reset() {
	*x = StockDeductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockDeductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDeductReply) ProtoMessage() {}

func (x *StockDeductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDeductReply.ProtoReflect.Descriptor instead.
func (*StockDeductReply) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_stock_proto_rawDescGZIP(), []int{1}
}

type StockDeductRevertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId    uint64 `protobuf:"varint,1,opt,name=productId,proto3" json:"productId"`
	ProductCount uint32 `protobuf:"varint,2,opt,name=productCount,proto3" json:"productCount"`
}

func (x *StockDeductRevertRequest) Reset() {
	*x = StockDeductRevertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockDeductRevertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDeductRevertRequest) ProtoMessage() {}

func (x *StockDeductRevertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDeductRevertRequest.ProtoReflect.Descriptor instead.
func (*StockDeductRevertRequest) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_stock_proto_rawDescGZIP(), []int{2}
}

func (x *StockDeductRevertRequest) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *StockDeductRevertRequest) GetProductCount() uint32 {
	if x != nil {
		return x.ProductCount
	}
	return 0
}

type StockDeductRevertReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StockDeductRevertReply) Reset() {
	*x = StockDeductRevertReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_stock_v1_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockDeductRevertReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDeductRevertReply) ProtoMessage() {}

func (x *StockDeductRevertReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_stock_v1_stock_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDeductRevertReply.ProtoReflect.Descriptor instead.
func (*StockDeductRevertReply) Descriptor() ([]byte, []int) {
	return file_api_stock_v1_stock_proto_rawDescGZIP(), []int{3}
}

var File_api_stock_v1_stock_proto protoreflect.FileDescriptor

var file_api_stock_v1_stock_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x68, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x6e, 0x0a, 0x18, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20,
	0x00, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x18, 0x0a, 0x16, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xbf, 0x01, 0x0a, 0x05, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x12, 0x51, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44,
	0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_stock_v1_stock_proto_rawDescOnce sync.Once
	file_api_stock_v1_stock_proto_rawDescData = file_api_stock_v1_stock_proto_rawDesc
)

func file_api_stock_v1_stock_proto_rawDescGZIP() []byte {
	file_api_stock_v1_stock_proto_rawDescOnce.Do(func() {
		file_api_stock_v1_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_stock_v1_stock_proto_rawDescData)
	})
	return file_api_stock_v1_stock_proto_rawDescData
}

var file_api_stock_v1_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_stock_v1_stock_proto_goTypes = []interface{}{
	(*StockDeductRequest)(nil),       // 0: api.stock.v1.StockDeductRequest
	(*StockDeductReply)(nil),         // 1: api.stock.v1.StockDeductReply
	(*StockDeductRevertRequest)(nil), // 2: api.stock.v1.StockDeductRevertRequest
	(*StockDeductRevertReply)(nil),   // 3: api.stock.v1.StockDeductRevertReply
}
var file_api_stock_v1_stock_proto_depIdxs = []int32{
	0, // 0: api.stock.v1.stock.StockDeduct:input_type -> api.stock.v1.StockDeductRequest
	2, // 1: api.stock.v1.stock.StockDeductRevert:input_type -> api.stock.v1.StockDeductRevertRequest
	1, // 2: api.stock.v1.stock.StockDeduct:output_type -> api.stock.v1.StockDeductReply
	3, // 3: api.stock.v1.stock.StockDeductRevert:output_type -> api.stock.v1.StockDeductRevertReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_stock_v1_stock_proto_init() }
func file_api_stock_v1_stock_proto_init() {
	if File_api_stock_v1_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_stock_v1_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockDeductRequest); i {
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
		file_api_stock_v1_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockDeductReply); i {
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
		file_api_stock_v1_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockDeductRevertRequest); i {
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
		file_api_stock_v1_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockDeductRevertReply); i {
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
			RawDescriptor: file_api_stock_v1_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_stock_v1_stock_proto_goTypes,
		DependencyIndexes: file_api_stock_v1_stock_proto_depIdxs,
		MessageInfos:      file_api_stock_v1_stock_proto_msgTypes,
	}.Build()
	File_api_stock_v1_stock_proto = out.File
	file_api_stock_v1_stock_proto_rawDesc = nil
	file_api_stock_v1_stock_proto_goTypes = nil
	file_api_stock_v1_stock_proto_depIdxs = nil
}
