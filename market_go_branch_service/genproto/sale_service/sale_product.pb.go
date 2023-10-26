// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: sale_product.proto

package sale_service

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

type SaleProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaleId    string  `protobuf:"bytes,1,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	ProductId string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SaleProduct) Reset() {
	*x = SaleProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProduct) ProtoMessage() {}

func (x *SaleProduct) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProduct.ProtoReflect.Descriptor instead.
func (*SaleProduct) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{0}
}

func (x *SaleProduct) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *SaleProduct) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *SaleProduct) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *SaleProduct) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SaleProductCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaleId    string  `protobuf:"bytes,1,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	ProductId string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SaleProductCreateReq) Reset() {
	*x = SaleProductCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductCreateReq) ProtoMessage() {}

func (x *SaleProductCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductCreateReq.ProtoReflect.Descriptor instead.
func (*SaleProductCreateReq) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{1}
}

func (x *SaleProductCreateReq) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *SaleProductCreateReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *SaleProductCreateReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *SaleProductCreateReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SaleProductCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SaleProductCreateResp) Reset() {
	*x = SaleProductCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductCreateResp) ProtoMessage() {}

func (x *SaleProductCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductCreateResp.ProtoReflect.Descriptor instead.
func (*SaleProductCreateResp) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{2}
}

func (x *SaleProductCreateResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SaleProductGetListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *SaleProductGetListReq) Reset() {
	*x = SaleProductGetListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductGetListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductGetListReq) ProtoMessage() {}

func (x *SaleProductGetListReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductGetListReq.ProtoReflect.Descriptor instead.
func (*SaleProductGetListReq) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{3}
}

func (x *SaleProductGetListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SaleProductGetListReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type SaleProductGetListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaleProducts []*SaleProduct `protobuf:"bytes,1,rep,name=saleProducts,proto3" json:"saleProducts,omitempty"`
	Count        int64          `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SaleProductGetListResp) Reset() {
	*x = SaleProductGetListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductGetListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductGetListResp) ProtoMessage() {}

func (x *SaleProductGetListResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductGetListResp.ProtoReflect.Descriptor instead.
func (*SaleProductGetListResp) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{4}
}

func (x *SaleProductGetListResp) GetSaleProducts() []*SaleProduct {
	if x != nil {
		return x.SaleProducts
	}
	return nil
}

func (x *SaleProductGetListResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SaleProductIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaleId string `protobuf:"bytes,1,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
}

func (x *SaleProductIdReq) Reset() {
	*x = SaleProductIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductIdReq) ProtoMessage() {}

func (x *SaleProductIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductIdReq.ProtoReflect.Descriptor instead.
func (*SaleProductIdReq) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{5}
}

func (x *SaleProductIdReq) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

type SaleProductUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SaleId    string  `protobuf:"bytes,1,opt,name=sale_id,json=saleId,proto3" json:"sale_id,omitempty"`
	ProductId string  `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Price     float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SaleProductUpdateReq) Reset() {
	*x = SaleProductUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductUpdateReq) ProtoMessage() {}

func (x *SaleProductUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductUpdateReq.ProtoReflect.Descriptor instead.
func (*SaleProductUpdateReq) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{6}
}

func (x *SaleProductUpdateReq) GetSaleId() string {
	if x != nil {
		return x.SaleId
	}
	return ""
}

func (x *SaleProductUpdateReq) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *SaleProductUpdateReq) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *SaleProductUpdateReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SaleProductUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SaleProductUpdateResp) Reset() {
	*x = SaleProductUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductUpdateResp) ProtoMessage() {}

func (x *SaleProductUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductUpdateResp.ProtoReflect.Descriptor instead.
func (*SaleProductUpdateResp) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{7}
}

func (x *SaleProductUpdateResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SaleProductDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SaleProductDeleteResp) Reset() {
	*x = SaleProductDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_product_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleProductDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleProductDeleteResp) ProtoMessage() {}

func (x *SaleProductDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_product_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleProductDeleteResp.ProtoReflect.Descriptor instead.
func (*SaleProductDeleteResp) Descriptor() ([]byte, []int) {
	return file_sale_product_proto_rawDescGZIP(), []int{8}
}

func (x *SaleProductDeleteResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sale_product_proto protoreflect.FileDescriptor

var file_sale_product_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x22, 0x77, 0x0a, 0x0b, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x14,
	0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x29,
	0x0a, 0x15, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x41, 0x0a, 0x15, 0x53, 0x61, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6d, 0x0a, 0x16,
	0x53, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x61, 0x6c, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0c, 0x73, 0x61, 0x6c, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x10, 0x53,
	0x61, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x53, 0x61, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x29, 0x0a, 0x15, 0x53,
	0x61, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x29, 0x0a, 0x15, 0x53, 0x61, 0x6c, 0x65, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_sale_product_proto_rawDescOnce sync.Once
	file_sale_product_proto_rawDescData = file_sale_product_proto_rawDesc
)

func file_sale_product_proto_rawDescGZIP() []byte {
	file_sale_product_proto_rawDescOnce.Do(func() {
		file_sale_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_sale_product_proto_rawDescData)
	})
	return file_sale_product_proto_rawDescData
}

var file_sale_product_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sale_product_proto_goTypes = []interface{}{
	(*SaleProduct)(nil),            // 0: sale_service.SaleProduct
	(*SaleProductCreateReq)(nil),   // 1: sale_service.SaleProductCreateReq
	(*SaleProductCreateResp)(nil),  // 2: sale_service.SaleProductCreateResp
	(*SaleProductGetListReq)(nil),  // 3: sale_service.SaleProductGetListReq
	(*SaleProductGetListResp)(nil), // 4: sale_service.SaleProductGetListResp
	(*SaleProductIdReq)(nil),       // 5: sale_service.SaleProductIdReq
	(*SaleProductUpdateReq)(nil),   // 6: sale_service.SaleProductUpdateReq
	(*SaleProductUpdateResp)(nil),  // 7: sale_service.SaleProductUpdateResp
	(*SaleProductDeleteResp)(nil),  // 8: sale_service.SaleProductDeleteResp
}
var file_sale_product_proto_depIdxs = []int32{
	0, // 0: sale_service.SaleProductGetListResp.saleProducts:type_name -> sale_service.SaleProduct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sale_product_proto_init() }
func file_sale_product_proto_init() {
	if File_sale_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sale_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProduct); i {
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
		file_sale_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductCreateReq); i {
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
		file_sale_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductCreateResp); i {
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
		file_sale_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductGetListReq); i {
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
		file_sale_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductGetListResp); i {
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
		file_sale_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductIdReq); i {
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
		file_sale_product_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductUpdateReq); i {
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
		file_sale_product_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductUpdateResp); i {
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
		file_sale_product_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleProductDeleteResp); i {
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
			RawDescriptor: file_sale_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sale_product_proto_goTypes,
		DependencyIndexes: file_sale_product_proto_depIdxs,
		MessageInfos:      file_sale_product_proto_msgTypes,
	}.Build()
	File_sale_product_proto = out.File
	file_sale_product_proto_rawDesc = nil
	file_sale_product_proto_goTypes = nil
	file_sale_product_proto_depIdxs = nil
}
