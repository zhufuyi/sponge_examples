// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.2
// source: api/inventory/v1/inventory.proto

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

type GetByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *GetByIDRequest) Reset() {
	*x = GetByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_inventory_v1_inventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDRequest) ProtoMessage() {}

func (x *GetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_inventory_v1_inventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDRequest.ProtoReflect.Descriptor instead.
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_inventory_v1_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *GetByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type InventoryDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Num     float32 `protobuf:"fixed32,4,opt,name=num,proto3" json:"num"`
	SoldNum int32   `protobuf:"varint,3,opt,name=soldNum,proto3" json:"soldNum"`
}

func (x *InventoryDetail) Reset() {
	*x = InventoryDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_inventory_v1_inventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryDetail) ProtoMessage() {}

func (x *InventoryDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_inventory_v1_inventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryDetail.ProtoReflect.Descriptor instead.
func (*InventoryDetail) Descriptor() ([]byte, []int) {
	return file_api_inventory_v1_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *InventoryDetail) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InventoryDetail) GetNum() float32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *InventoryDetail) GetSoldNum() int32 {
	if x != nil {
		return x.SoldNum
	}
	return 0
}

type GetByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InventoryDetail *InventoryDetail `protobuf:"bytes,1,opt,name=inventoryDetail,proto3" json:"inventoryDetail"`
}

func (x *GetByIDReply) Reset() {
	*x = GetByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_inventory_v1_inventory_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIDReply) ProtoMessage() {}

func (x *GetByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_inventory_v1_inventory_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIDReply.ProtoReflect.Descriptor instead.
func (*GetByIDReply) Descriptor() ([]byte, []int) {
	return file_api_inventory_v1_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *GetByIDReply) GetInventoryDetail() *InventoryDetail {
	if x != nil {
		return x.InventoryDetail
	}
	return nil
}

var File_api_inventory_v1_inventory_proto protoreflect.FileDescriptor

var file_api_inventory_v1_inventory_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e,
	0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x6f, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x73, 0x6f, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x22, 0x5b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4b, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x32, 0x5a, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x4d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_inventory_v1_inventory_proto_rawDescOnce sync.Once
	file_api_inventory_v1_inventory_proto_rawDescData = file_api_inventory_v1_inventory_proto_rawDesc
)

func file_api_inventory_v1_inventory_proto_rawDescGZIP() []byte {
	file_api_inventory_v1_inventory_proto_rawDescOnce.Do(func() {
		file_api_inventory_v1_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_inventory_v1_inventory_proto_rawDescData)
	})
	return file_api_inventory_v1_inventory_proto_rawDescData
}

var file_api_inventory_v1_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_inventory_v1_inventory_proto_goTypes = []interface{}{
	(*GetByIDRequest)(nil),  // 0: api.inventory.v1.GetByIDRequest
	(*InventoryDetail)(nil), // 1: api.inventory.v1.InventoryDetail
	(*GetByIDReply)(nil),    // 2: api.inventory.v1.GetByIDReply
}
var file_api_inventory_v1_inventory_proto_depIdxs = []int32{
	1, // 0: api.inventory.v1.GetByIDReply.inventoryDetail:type_name -> api.inventory.v1.InventoryDetail
	0, // 1: api.inventory.v1.Inventory.GetByID:input_type -> api.inventory.v1.GetByIDRequest
	2, // 2: api.inventory.v1.Inventory.GetByID:output_type -> api.inventory.v1.GetByIDReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_inventory_v1_inventory_proto_init() }
func file_api_inventory_v1_inventory_proto_init() {
	if File_api_inventory_v1_inventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_inventory_v1_inventory_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDRequest); i {
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
		file_api_inventory_v1_inventory_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryDetail); i {
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
		file_api_inventory_v1_inventory_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIDReply); i {
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
			RawDescriptor: file_api_inventory_v1_inventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_inventory_v1_inventory_proto_goTypes,
		DependencyIndexes: file_api_inventory_v1_inventory_proto_depIdxs,
		MessageInfos:      file_api_inventory_v1_inventory_proto_msgTypes,
	}.Build()
	File_api_inventory_v1_inventory_proto = out.File
	file_api_inventory_v1_inventory_proto_rawDesc = nil
	file_api_inventory_v1_inventory_proto_goTypes = nil
	file_api_inventory_v1_inventory_proto_depIdxs = nil
}
