// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.2
// source: manage_product.proto

package frontend_product

import (
	_ "github.com/777continue/gomall/app/frontend/hertz_gen/api"
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" form:"id" query:"id"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name" query:"name"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" form:"description" query:"description"`
	Picture     string   `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty" form:"picture" query:"picture"`
	Price       float32  `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty" form:"price" query:"price"`
	Stock       uint32   `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty" form:"stock" query:"stock"`
	Categories  []string `protobuf:"bytes,7,rep,name=categories,proto3" json:"categories,omitempty" form:"categories" query:"categories"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[0]
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
	return file_manage_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *Product) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

type ListProductsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page         int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" query:"page"`
	PageSize     int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" query:"page_size"`
	CategoryName string `protobuf:"bytes,3,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty" query:"category_name"`
}

func (x *ListProductsReq) Reset() {
	*x = ListProductsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsReq) ProtoMessage() {}

func (x *ListProductsReq) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsReq.ProtoReflect.Descriptor instead.
func (*ListProductsReq) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{1}
}

func (x *ListProductsReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListProductsReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListProductsReq) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

type ListProductsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty" form:"products" query:"products"`
}

func (x *ListProductsResp) Reset() {
	*x = ListProductsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResp) ProtoMessage() {}

func (x *ListProductsResp) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResp.ProtoReflect.Descriptor instead.
func (*ListProductsResp) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{2}
}

func (x *ListProductsResp) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type GetProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" path:"id"`
}

func (x *GetProductReq) Reset() {
	*x = GetProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReq) ProtoMessage() {}

func (x *GetProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReq.ProtoReflect.Descriptor instead.
func (*GetProductReq) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty" form:"product" query:"product"`
}

func (x *GetProductResp) Reset() {
	*x = GetProductResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResp) ProtoMessage() {}

func (x *GetProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResp.ProtoReflect.Descriptor instead.
func (*GetProductResp) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductResp) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

type AddProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" form:"description"`
	Picture     []byte  `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty" form:"picture"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty" form:"price"`
	Stock       uint32  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty" form:"stock"`
}

func (x *AddProductReq) Reset() {
	*x = AddProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductReq) ProtoMessage() {}

func (x *AddProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductReq.ProtoReflect.Descriptor instead.
func (*AddProductReq) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{5}
}

func (x *AddProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddProductReq) GetPicture() []byte {
	if x != nil {
		return x.Picture
	}
	return nil
}

func (x *AddProductReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AddProductReq) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type AddProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddProductResp) Reset() {
	*x = AddProductResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProductResp) ProtoMessage() {}

func (x *AddProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProductResp.ProtoReflect.Descriptor instead.
func (*AddProductResp) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{6}
}

type UpdateProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" path:"id"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" form:"name"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" form:"description"`
	Picture     []byte  `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty" form:"picture"`
	Price       float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty" form:"price"`
	Stock       uint32  `protobuf:"varint,6,opt,name=stock,proto3" json:"stock,omitempty" form:"stock"`
}

func (x *UpdateProductReq) Reset() {
	*x = UpdateProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductReq) ProtoMessage() {}

func (x *UpdateProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductReq.ProtoReflect.Descriptor instead.
func (*UpdateProductReq) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProductReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateProductReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateProductReq) GetPicture() []byte {
	if x != nil {
		return x.Picture
	}
	return nil
}

func (x *UpdateProductReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateProductReq) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

type UpdateProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateProductResp) Reset() {
	*x = UpdateProductResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductResp) ProtoMessage() {}

func (x *UpdateProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductResp.ProtoReflect.Descriptor instead.
func (*UpdateProductResp) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{8}
}

type DeleteProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" path:"id"`
}

func (x *DeleteProductReq) Reset() {
	*x = DeleteProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductReq) ProtoMessage() {}

func (x *DeleteProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductReq.ProtoReflect.Descriptor instead.
func (*DeleteProductReq) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteProductReq) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProductResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteProductResp) Reset() {
	*x = DeleteProductResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manage_product_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductResp) ProtoMessage() {}

func (x *DeleteProductResp) ProtoReflect() protoreflect.Message {
	mi := &file_manage_product_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductResp.ProtoReflect.Descriptor instead.
func (*DeleteProductResp) Descriptor() ([]byte, []int) {
	return file_manage_product_proto_rawDescGZIP(), []int{10}
}

var File_manage_product_proto protoreflect.FileDescriptor

var file_manage_product_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xb2,
	0xbb, 0x18, 0x04, 0x70, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0d, 0xb2, 0xbb, 0x18, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x11, 0xb2, 0xbb, 0x18, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x49, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xd2, 0xbb, 0x18, 0x02, 0x69,
	0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0xc9, 0x01, 0x0a,
	0x0d, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1c,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xbb,
	0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0f, 0xe2, 0xbb, 0x18, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x0b, 0xe2, 0xbb, 0x18, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x42, 0x09, 0xe2, 0xbb, 0x18, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x09, 0xe2, 0xbb, 0x18, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x10, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0xe4, 0x01, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x16, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xd2, 0xbb, 0x18,
	0x02, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xe2, 0xbb, 0x18, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe2, 0xbb, 0x18, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0b, 0xe2, 0xbb, 0x18, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1f, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x42, 0x09,
	0xe2, 0xbb, 0x18, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1f, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x09, 0xe2, 0xbb, 0x18, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2a, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x06, 0xd2, 0xbb, 0x18, 0x02, 0x69, 0x64, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x32, 0x9d, 0x04, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x64, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x12, 0x21, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0xca, 0xc1, 0x18, 0x09, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x63, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0xca, 0xc1, 0x18, 0x0e, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5e, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x41, 0x64,
	0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x0d, 0xd2, 0xc1,
	0x18, 0x09, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x6c, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x66,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x23, 0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0xda, 0xc1, 0x18, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x23,
	0x2e, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x12, 0xe2, 0xc1, 0x18, 0x0e, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x37, 0x37, 0x37, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x65, 0x2f, 0x67, 0x6f, 0x6d, 0x61, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2f, 0x68, 0x65, 0x72, 0x74, 0x7a, 0x5f, 0x67, 0x65, 0x6e, 0x2f,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manage_product_proto_rawDescOnce sync.Once
	file_manage_product_proto_rawDescData = file_manage_product_proto_rawDesc
)

func file_manage_product_proto_rawDescGZIP() []byte {
	file_manage_product_proto_rawDescOnce.Do(func() {
		file_manage_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_manage_product_proto_rawDescData)
	})
	return file_manage_product_proto_rawDescData
}

var file_manage_product_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_manage_product_proto_goTypes = []interface{}{
	(*Product)(nil),           // 0: frontend.product.Product
	(*ListProductsReq)(nil),   // 1: frontend.product.ListProductsReq
	(*ListProductsResp)(nil),  // 2: frontend.product.ListProductsResp
	(*GetProductReq)(nil),     // 3: frontend.product.GetProductReq
	(*GetProductResp)(nil),    // 4: frontend.product.GetProductResp
	(*AddProductReq)(nil),     // 5: frontend.product.AddProductReq
	(*AddProductResp)(nil),    // 6: frontend.product.AddProductResp
	(*UpdateProductReq)(nil),  // 7: frontend.product.UpdateProductReq
	(*UpdateProductResp)(nil), // 8: frontend.product.UpdateProductResp
	(*DeleteProductReq)(nil),  // 9: frontend.product.DeleteProductReq
	(*DeleteProductResp)(nil), // 10: frontend.product.DeleteProductResp
}
var file_manage_product_proto_depIdxs = []int32{
	0,  // 0: frontend.product.ListProductsResp.products:type_name -> frontend.product.Product
	0,  // 1: frontend.product.GetProductResp.product:type_name -> frontend.product.Product
	1,  // 2: frontend.product.ProductManageService.ListProducts:input_type -> frontend.product.ListProductsReq
	3,  // 3: frontend.product.ProductManageService.GetProduct:input_type -> frontend.product.GetProductReq
	5,  // 4: frontend.product.ProductManageService.AddProduct:input_type -> frontend.product.AddProductReq
	7,  // 5: frontend.product.ProductManageService.UpdateProduct:input_type -> frontend.product.UpdateProductReq
	9,  // 6: frontend.product.ProductManageService.DeleteProduct:input_type -> frontend.product.DeleteProductReq
	2,  // 7: frontend.product.ProductManageService.ListProducts:output_type -> frontend.product.ListProductsResp
	4,  // 8: frontend.product.ProductManageService.GetProduct:output_type -> frontend.product.GetProductResp
	6,  // 9: frontend.product.ProductManageService.AddProduct:output_type -> frontend.product.AddProductResp
	8,  // 10: frontend.product.ProductManageService.UpdateProduct:output_type -> frontend.product.UpdateProductResp
	10, // 11: frontend.product.ProductManageService.DeleteProduct:output_type -> frontend.product.DeleteProductResp
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_manage_product_proto_init() }
func file_manage_product_proto_init() {
	if File_manage_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manage_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_manage_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsReq); i {
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
		file_manage_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsResp); i {
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
		file_manage_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReq); i {
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
		file_manage_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductResp); i {
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
		file_manage_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductReq); i {
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
		file_manage_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProductResp); i {
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
		file_manage_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductReq); i {
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
		file_manage_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductResp); i {
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
		file_manage_product_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductReq); i {
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
		file_manage_product_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductResp); i {
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
			RawDescriptor: file_manage_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manage_product_proto_goTypes,
		DependencyIndexes: file_manage_product_proto_depIdxs,
		MessageInfos:      file_manage_product_proto_msgTypes,
	}.Build()
	File_manage_product_proto = out.File
	file_manage_product_proto_rawDesc = nil
	file_manage_product_proto_goTypes = nil
	file_manage_product_proto_depIdxs = nil
}
