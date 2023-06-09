// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/relation/v1/relation.proto

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

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	FollowedUid uint64 `protobuf:"varint,2,opt,name=followedUid,proto3" json:"followedUid"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FollowRequest) GetFollowedUid() uint64 {
	if x != nil {
		return x.FollowedUid
	}
	return 0
}

type FollowReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FollowReply) Reset() {
	*x = FollowReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowReply) ProtoMessage() {}

func (x *FollowReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowReply.ProtoReflect.Descriptor instead.
func (*FollowReply) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{1}
}

type UnfollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	FollowedUid uint64 `protobuf:"varint,2,opt,name=followedUid,proto3" json:"followedUid"`
}

func (x *UnfollowRequest) Reset() {
	*x = UnfollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowRequest) ProtoMessage() {}

func (x *UnfollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowRequest.ProtoReflect.Descriptor instead.
func (*UnfollowRequest) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{2}
}

func (x *UnfollowRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UnfollowRequest) GetFollowedUid() uint64 {
	if x != nil {
		return x.FollowedUid
	}
	return 0
}

type UnfollowReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnfollowReply) Reset() {
	*x = UnfollowReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnfollowReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnfollowReply) ProtoMessage() {}

func (x *UnfollowReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnfollowReply.ProtoReflect.Descriptor instead.
func (*UnfollowReply) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{3}
}

type ListFollowingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit"`
}

func (x *ListFollowingRequest) Reset() {
	*x = ListFollowingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowingRequest) ProtoMessage() {}

func (x *ListFollowingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowingRequest.ProtoReflect.Descriptor instead.
func (*ListFollowingRequest) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{4}
}

func (x *ListFollowingRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListFollowingRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListFollowingRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListFollowingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowedUids []uint64 `protobuf:"varint,1,rep,packed,name=followedUids,proto3" json:"followedUids"`
	Total        int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *ListFollowingReply) Reset() {
	*x = ListFollowingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowingReply) ProtoMessage() {}

func (x *ListFollowingReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowingReply.ProtoReflect.Descriptor instead.
func (*ListFollowingReply) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{5}
}

func (x *ListFollowingReply) GetFollowedUids() []uint64 {
	if x != nil {
		return x.FollowedUids
	}
	return nil
}

func (x *ListFollowingReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ListFollowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	Page   int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit"`
}

func (x *ListFollowerRequest) Reset() {
	*x = ListFollowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowerRequest) ProtoMessage() {}

func (x *ListFollowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowerRequest.ProtoReflect.Descriptor instead.
func (*ListFollowerRequest) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{6}
}

func (x *ListFollowerRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListFollowerRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListFollowerRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListFollowerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FollowerUids []uint64 `protobuf:"varint,1,rep,packed,name=followerUids,proto3" json:"followerUids"`
	Total        int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *ListFollowerReply) Reset() {
	*x = ListFollowerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFollowerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFollowerReply) ProtoMessage() {}

func (x *ListFollowerReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFollowerReply.ProtoReflect.Descriptor instead.
func (*ListFollowerReply) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{7}
}

func (x *ListFollowerReply) GetFollowerUids() []uint64 {
	if x != nil {
		return x.FollowerUids
	}
	return nil
}

func (x *ListFollowerReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type BatchGetRelationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId"`
	Uids   []uint64 `protobuf:"varint,2,rep,packed,name=uids,proto3" json:"uids"`
}

func (x *BatchGetRelationRequest) Reset() {
	*x = BatchGetRelationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetRelationRequest) ProtoMessage() {}

func (x *BatchGetRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetRelationRequest.ProtoReflect.Descriptor instead.
func (*BatchGetRelationRequest) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{8}
}

func (x *BatchGetRelationRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BatchGetRelationRequest) GetUids() []uint64 {
	if x != nil {
		return x.Uids
	}
	return nil
}

type BatchGetRelationReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result map[uint64]int32 `protobuf:"bytes,1,rep,name=result,proto3" json:"result" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *BatchGetRelationReply) Reset() {
	*x = BatchGetRelationReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_relation_v1_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetRelationReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetRelationReply) ProtoMessage() {}

func (x *BatchGetRelationReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_relation_v1_relation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetRelationReply.ProtoReflect.Descriptor instead.
func (*BatchGetRelationReply) Descriptor() ([]byte, []int) {
	return file_api_relation_v1_relation_proto_rawDescGZIP(), []int{9}
}

func (x *BatchGetRelationReply) GetResult() map[uint64]int32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_api_relation_v1_relation_proto protoreflect.FileDescriptor

var file_api_relation_v1_relation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0d, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x32, 0x02, 0x28, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0b,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x55, 0x69, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5d, 0x0a, 0x0f, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0b, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x55, 0x69, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x73, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x1a, 0x02, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x4e, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x55, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x55, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x72, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1d, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x4d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x55, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0c, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x55, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x58,
	0x0a, 0x17, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x69,
	0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x04, 0x75, 0x69, 0x64, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x15, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39,
	0x0a, 0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xce, 0x03, 0x0a, 0x0f, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a,
	0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x08, 0x55, 0x6e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x12, 0x66, 0x0a, 0x10, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_relation_v1_relation_proto_rawDescOnce sync.Once
	file_api_relation_v1_relation_proto_rawDescData = file_api_relation_v1_relation_proto_rawDesc
)

func file_api_relation_v1_relation_proto_rawDescGZIP() []byte {
	file_api_relation_v1_relation_proto_rawDescOnce.Do(func() {
		file_api_relation_v1_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_relation_v1_relation_proto_rawDescData)
	})
	return file_api_relation_v1_relation_proto_rawDescData
}

var file_api_relation_v1_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_relation_v1_relation_proto_goTypes = []interface{}{
	(*FollowRequest)(nil),           // 0: api.relation.v1.FollowRequest
	(*FollowReply)(nil),             // 1: api.relation.v1.FollowReply
	(*UnfollowRequest)(nil),         // 2: api.relation.v1.UnfollowRequest
	(*UnfollowReply)(nil),           // 3: api.relation.v1.UnfollowReply
	(*ListFollowingRequest)(nil),    // 4: api.relation.v1.ListFollowingRequest
	(*ListFollowingReply)(nil),      // 5: api.relation.v1.ListFollowingReply
	(*ListFollowerRequest)(nil),     // 6: api.relation.v1.ListFollowerRequest
	(*ListFollowerReply)(nil),       // 7: api.relation.v1.ListFollowerReply
	(*BatchGetRelationRequest)(nil), // 8: api.relation.v1.BatchGetRelationRequest
	(*BatchGetRelationReply)(nil),   // 9: api.relation.v1.BatchGetRelationReply
	nil,                             // 10: api.relation.v1.BatchGetRelationReply.ResultEntry
}
var file_api_relation_v1_relation_proto_depIdxs = []int32{
	10, // 0: api.relation.v1.BatchGetRelationReply.result:type_name -> api.relation.v1.BatchGetRelationReply.ResultEntry
	0,  // 1: api.relation.v1.RelationService.Follow:input_type -> api.relation.v1.FollowRequest
	2,  // 2: api.relation.v1.RelationService.Unfollow:input_type -> api.relation.v1.UnfollowRequest
	4,  // 3: api.relation.v1.RelationService.ListFollowing:input_type -> api.relation.v1.ListFollowingRequest
	6,  // 4: api.relation.v1.RelationService.ListFollower:input_type -> api.relation.v1.ListFollowerRequest
	8,  // 5: api.relation.v1.RelationService.BatchGetRelation:input_type -> api.relation.v1.BatchGetRelationRequest
	1,  // 6: api.relation.v1.RelationService.Follow:output_type -> api.relation.v1.FollowReply
	3,  // 7: api.relation.v1.RelationService.Unfollow:output_type -> api.relation.v1.UnfollowReply
	5,  // 8: api.relation.v1.RelationService.ListFollowing:output_type -> api.relation.v1.ListFollowingReply
	7,  // 9: api.relation.v1.RelationService.ListFollower:output_type -> api.relation.v1.ListFollowerReply
	9,  // 10: api.relation.v1.RelationService.BatchGetRelation:output_type -> api.relation.v1.BatchGetRelationReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_relation_v1_relation_proto_init() }
func file_api_relation_v1_relation_proto_init() {
	if File_api_relation_v1_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_relation_v1_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_api_relation_v1_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowReply); i {
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
		file_api_relation_v1_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowRequest); i {
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
		file_api_relation_v1_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnfollowReply); i {
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
		file_api_relation_v1_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowingRequest); i {
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
		file_api_relation_v1_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowingReply); i {
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
		file_api_relation_v1_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowerRequest); i {
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
		file_api_relation_v1_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFollowerReply); i {
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
		file_api_relation_v1_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetRelationRequest); i {
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
		file_api_relation_v1_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetRelationReply); i {
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
			RawDescriptor: file_api_relation_v1_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_relation_v1_relation_proto_goTypes,
		DependencyIndexes: file_api_relation_v1_relation_proto_depIdxs,
		MessageInfos:      file_api_relation_v1_relation_proto_msgTypes,
	}.Build()
	File_api_relation_v1_relation_proto = out.File
	file_api_relation_v1_relation_proto_rawDesc = nil
	file_api_relation_v1_relation_proto_goTypes = nil
	file_api_relation_v1_relation_proto_depIdxs = nil
}
