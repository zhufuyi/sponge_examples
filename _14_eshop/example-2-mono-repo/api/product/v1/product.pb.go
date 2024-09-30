// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.2
// source: api/product/v1/product.proto

package v1

import (
	types "eshop/api/types"
	
	
	
	
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

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`               // 商品名称
	Price       int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price"`            // 商品价格
	Photo       string `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo"`             // 商品图片
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description"` // 商品描述
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateProductRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *CreateProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
}

func (x *CreateProductReply) Reset() {
	*x = CreateProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductReply) ProtoMessage() {}

func (x *CreateProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductReply.ProtoReflect.Descriptor instead.
func (*CreateProductReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductReply) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *DeleteProductByIDRequest) Reset() {
	*x = DeleteProductByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductByIDRequest) ProtoMessage() {}

func (x *DeleteProductByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteProductByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProductByIDReply) Reset() {
	*x = DeleteProductByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductByIDReply) ProtoMessage() {}

func (x *DeleteProductByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductByIDReply.ProtoReflect.Descriptor instead.
func (*DeleteProductByIDReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{3}
}

type UpdateProductByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`               // 商品名称
	Price       int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price"`            // 商品价格
	Photo       string `protobuf:"bytes,4,opt,name=photo,proto3" json:"photo"`             // 商品图片
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description"` // 商品描述
}

func (x *UpdateProductByIDRequest) Reset() {
	*x = UpdateProductByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductByIDRequest) ProtoMessage() {}

func (x *UpdateProductByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateProductByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProductByIDRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductByIDRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductByIDRequest) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *UpdateProductByIDRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateProductByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProductByIDReply) Reset() {
	*x = UpdateProductByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductByIDReply) ProtoMessage() {}

func (x *UpdateProductByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductByIDReply.ProtoReflect.Descriptor instead.
func (*UpdateProductByIDReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{5}
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	CreatedAt   string `protobuf:"bytes,2,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt   string `protobuf:"bytes,3,opt,name=updatedAt,proto3" json:"updatedAt"`
	Name        string `protobuf:"bytes,4,opt,name=name,proto3" json:"name"`               // 商品名称
	Price       int32  `protobuf:"varint,5,opt,name=price,proto3" json:"price"`            // 商品价格
	Photo       string `protobuf:"bytes,6,opt,name=photo,proto3" json:"photo"`             // 商品图片
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description"` // 商品描述
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{6}
}

func (x *Product) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Product) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetPhoto() string {
	if x != nil {
		return x.Photo
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetProductByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id" uri:"id"`
}

func (x *GetProductByIDRequest) Reset() {
	*x = GetProductByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIDRequest) ProtoMessage() {}

func (x *GetProductByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIDRequest.ProtoReflect.Descriptor instead.
func (*GetProductByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{7}
}

func (x *GetProductByIDRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProductByIDReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product"`
}

func (x *GetProductByIDReply) Reset() {
	*x = GetProductByIDReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIDReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIDReply) ProtoMessage() {}

func (x *GetProductByIDReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIDReply.ProtoReflect.Descriptor instead.
func (*GetProductByIDReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{8}
}

func (x *GetProductByIDReply) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type ListProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params *types.Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (x *ListProductRequest) Reset() {
	*x = ListProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductRequest) ProtoMessage() {}

func (x *ListProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductRequest.ProtoReflect.Descriptor instead.
func (*ListProductRequest) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{9}
}

func (x *ListProductRequest) GetParams() *types.Params {
	if x != nil {
		return x.Params
	}
	return nil
}

type ListProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total    int64      `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products"`
}

func (x *ListProductReply) Reset() {
	*x = ListProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_product_v1_product_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductReply) ProtoMessage() {}

func (x *ListProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_product_v1_product_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductReply.ProtoReflect.Descriptor instead.
func (*ListProductReply) Descriptor() ([]byte, []int) {
	return file_api_product_v1_product_proto_rawDescGZIP(), []int{10}
}

func (x *ListProductReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListProductReply) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_api_product_v1_product_proto protoreflect.FileDescriptor

var file_api_product_v1_product_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x15,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x78, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x40, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x9b, 0x01,
	0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0d, 0x9a, 0x84, 0x9e, 0x03, 0x08, 0x75, 0x72, 0x69,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x18, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0xb7, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x14, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x9a, 0x84, 0x9e,
	0x03, 0x08, 0x75, 0x72, 0x69, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x3f, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x5d, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0xeb, 0x06, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0xa7, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x53, 0x92, 0x41, 0x36, 0x12, 0x0e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x24,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0xa5,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x28, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x45, 0x92, 0x41, 0x26, 0x12, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x1a, 0x14, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x2a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xa8, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x48, 0x92, 0x41, 0x26, 0x12, 0x0e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x14, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x20, 0x62, 0x79, 0x20,
	0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01,
	0x2a, 0x12, 0xa4, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x4d, 0x92, 0x41, 0x2e, 0x12, 0x12,
	0x67, 0x65, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x20, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x1a, 0x18, 0x67, 0x65, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x20,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x20, 0x62, 0x79, 0x20, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0xbb, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x6d, 0x92, 0x41, 0x4b, 0x12, 0x1e, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x20, 0x62,
	0x79, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x1a, 0x29, 0x6c, 0x69,
	0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x20, 0x62,
	0x79, 0x20, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x63, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x6c,
	0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0xb9, 0x01, 0x5a, 0x17, 0x65, 0x73, 0x68, 0x6f, 0x70,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x92, 0x41, 0x9c, 0x01, 0x12, 0x17, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x20, 0x61, 0x70, 0x69, 0x20, 0x64, 0x6f, 0x63, 0x73, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x1a,
	0x0f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x33, 0x30, 0x31, 0x38, 0x30,
	0x2a, 0x02, 0x01, 0x02, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x48, 0x0a, 0x46, 0x0a, 0x0a, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x38, 0x08, 0x02, 0x12, 0x23, 0x54, 0x79,
	0x70, 0x65, 0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x2d, 0x6a,
	0x77, 0x74, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_product_v1_product_proto_rawDescOnce sync.Once
	file_api_product_v1_product_proto_rawDescData = file_api_product_v1_product_proto_rawDesc
)

func file_api_product_v1_product_proto_rawDescGZIP() []byte {
	file_api_product_v1_product_proto_rawDescOnce.Do(func() {
		file_api_product_v1_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_product_v1_product_proto_rawDescData)
	})
	return file_api_product_v1_product_proto_rawDescData
}

var file_api_product_v1_product_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_product_v1_product_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),     // 0: api.product.v1.CreateProductRequest
	(*CreateProductReply)(nil),       // 1: api.product.v1.CreateProductReply
	(*DeleteProductByIDRequest)(nil), // 2: api.product.v1.DeleteProductByIDRequest
	(*DeleteProductByIDReply)(nil),   // 3: api.product.v1.DeleteProductByIDReply
	(*UpdateProductByIDRequest)(nil), // 4: api.product.v1.UpdateProductByIDRequest
	(*UpdateProductByIDReply)(nil),   // 5: api.product.v1.UpdateProductByIDReply
	(*Product)(nil),                  // 6: api.product.v1.Product
	(*GetProductByIDRequest)(nil),    // 7: api.product.v1.GetProductByIDRequest
	(*GetProductByIDReply)(nil),      // 8: api.product.v1.GetProductByIDReply
	(*ListProductRequest)(nil),       // 9: api.product.v1.ListProductRequest
	(*ListProductReply)(nil),         // 10: api.product.v1.ListProductReply
	(*types.Params)(nil),             // 11: api.types.Params
}
var file_api_product_v1_product_proto_depIdxs = []int32{
	6,  // 0: api.product.v1.GetProductByIDReply.product:type_name -> api.product.v1.Product
	11, // 1: api.product.v1.ListProductRequest.params:type_name -> api.types.Params
	6,  // 2: api.product.v1.ListProductReply.products:type_name -> api.product.v1.Product
	0,  // 3: api.product.v1.product.Create:input_type -> api.product.v1.CreateProductRequest
	2,  // 4: api.product.v1.product.DeleteByID:input_type -> api.product.v1.DeleteProductByIDRequest
	4,  // 5: api.product.v1.product.UpdateByID:input_type -> api.product.v1.UpdateProductByIDRequest
	7,  // 6: api.product.v1.product.GetByID:input_type -> api.product.v1.GetProductByIDRequest
	9,  // 7: api.product.v1.product.List:input_type -> api.product.v1.ListProductRequest
	1,  // 8: api.product.v1.product.Create:output_type -> api.product.v1.CreateProductReply
	3,  // 9: api.product.v1.product.DeleteByID:output_type -> api.product.v1.DeleteProductByIDReply
	5,  // 10: api.product.v1.product.UpdateByID:output_type -> api.product.v1.UpdateProductByIDReply
	8,  // 11: api.product.v1.product.GetByID:output_type -> api.product.v1.GetProductByIDReply
	10, // 12: api.product.v1.product.List:output_type -> api.product.v1.ListProductReply
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_product_v1_product_proto_init() }
func file_api_product_v1_product_proto_init() {
	if File_api_product_v1_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_product_v1_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductReply); i {
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
		file_api_product_v1_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductByIDRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductByIDReply); i {
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
		file_api_product_v1_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductByIDRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductByIDReply); i {
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
		file_api_product_v1_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_api_product_v1_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIDRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIDReply); i {
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
		file_api_product_v1_product_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductRequest); i {
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
		file_api_product_v1_product_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductReply); i {
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
			RawDescriptor: file_api_product_v1_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_product_v1_product_proto_goTypes,
		DependencyIndexes: file_api_product_v1_product_proto_depIdxs,
		MessageInfos:      file_api_product_v1_product_proto_msgTypes,
	}.Build()
	File_api_product_v1_product_proto = out.File
	file_api_product_v1_product_proto_rawDesc = nil
	file_api_product_v1_product_proto_goTypes = nil
	file_api_product_v1_product_proto_depIdxs = nil
}