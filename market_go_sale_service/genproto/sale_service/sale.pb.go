// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: sale.proto

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

type Sale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId        string  `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ShopAssistantId string  `protobuf:"bytes,3,opt,name=shop_assistant_id,json=shopAssistantId,proto3" json:"shop_assistant_id,omitempty"`
	CashierId       string  `protobuf:"bytes,4,opt,name=cashier_id,json=cashierId,proto3" json:"cashier_id,omitempty"`
	Price           float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	PaymentType     string  `protobuf:"bytes,6,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	Status          string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	ClientName      string  `protobuf:"bytes,8,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	CreatedAt       string  `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       string  `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       string  `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Sale) Reset() {
	*x = Sale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sale) ProtoMessage() {}

func (x *Sale) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sale.ProtoReflect.Descriptor instead.
func (*Sale) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{0}
}

func (x *Sale) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sale) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *Sale) GetShopAssistantId() string {
	if x != nil {
		return x.ShopAssistantId
	}
	return ""
}

func (x *Sale) GetCashierId() string {
	if x != nil {
		return x.CashierId
	}
	return ""
}

func (x *Sale) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Sale) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *Sale) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Sale) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Sale) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Sale) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Sale) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type SaleCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId        string  `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ShopAssistantId string  `protobuf:"bytes,2,opt,name=shop_assistant_id,json=shopAssistantId,proto3" json:"shop_assistant_id,omitempty"`
	CashierId       string  `protobuf:"bytes,3,opt,name=cashier_id,json=cashierId,proto3" json:"cashier_id,omitempty"`
	Price           float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	PaymentType     string  `protobuf:"bytes,5,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	Status          string  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	ClientName      string  `protobuf:"bytes,7,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
}

func (x *SaleCreateReq) Reset() {
	*x = SaleCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleCreateReq) ProtoMessage() {}

func (x *SaleCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleCreateReq.ProtoReflect.Descriptor instead.
func (*SaleCreateReq) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{1}
}

func (x *SaleCreateReq) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *SaleCreateReq) GetShopAssistantId() string {
	if x != nil {
		return x.ShopAssistantId
	}
	return ""
}

func (x *SaleCreateReq) GetCashierId() string {
	if x != nil {
		return x.CashierId
	}
	return ""
}

func (x *SaleCreateReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SaleCreateReq) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *SaleCreateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SaleCreateReq) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

type SaleCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SaleCreateResp) Reset() {
	*x = SaleCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleCreateResp) ProtoMessage() {}

func (x *SaleCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleCreateResp.ProtoReflect.Descriptor instead.
func (*SaleCreateResp) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{2}
}

func (x *SaleCreateResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SaleGetListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page            int64   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit           int64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	BranchId        string  `protobuf:"bytes,3,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ClientName      string  `protobuf:"bytes,4,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	PaymentType     string  `protobuf:"bytes,5,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	ShopAssistantId string  `protobuf:"bytes,6,opt,name=shop_assistant_id,json=shopAssistantId,proto3" json:"shop_assistant_id,omitempty"`
	CashierId       string  `protobuf:"bytes,7,opt,name=cashier_id,json=cashierId,proto3" json:"cashier_id,omitempty"`
	Status          string  `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAtFrom   string  `protobuf:"bytes,9,opt,name=created_at_from,json=createdAtFrom,proto3" json:"created_at_from,omitempty"`
	CreatedAtTo     string  `protobuf:"bytes,10,opt,name=created_at_to,json=createdAtTo,proto3" json:"created_at_to,omitempty"`
	PriceFrom       float32 `protobuf:"fixed32,11,opt,name=price_from,json=priceFrom,proto3" json:"price_from,omitempty"`
	PriceTo         float32 `protobuf:"fixed32,12,opt,name=price_to,json=priceTo,proto3" json:"price_to,omitempty"`
}

func (x *SaleGetListReq) Reset() {
	*x = SaleGetListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleGetListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleGetListReq) ProtoMessage() {}

func (x *SaleGetListReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleGetListReq.ProtoReflect.Descriptor instead.
func (*SaleGetListReq) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{3}
}

func (x *SaleGetListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SaleGetListReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SaleGetListReq) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *SaleGetListReq) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *SaleGetListReq) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *SaleGetListReq) GetShopAssistantId() string {
	if x != nil {
		return x.ShopAssistantId
	}
	return ""
}

func (x *SaleGetListReq) GetCashierId() string {
	if x != nil {
		return x.CashierId
	}
	return ""
}

func (x *SaleGetListReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SaleGetListReq) GetCreatedAtFrom() string {
	if x != nil {
		return x.CreatedAtFrom
	}
	return ""
}

func (x *SaleGetListReq) GetCreatedAtTo() string {
	if x != nil {
		return x.CreatedAtTo
	}
	return ""
}

func (x *SaleGetListReq) GetPriceFrom() float32 {
	if x != nil {
		return x.PriceFrom
	}
	return 0
}

func (x *SaleGetListReq) GetPriceTo() float32 {
	if x != nil {
		return x.PriceTo
	}
	return 0
}

type SaleGetListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sales []*Sale `protobuf:"bytes,1,rep,name=sales,proto3" json:"sales,omitempty"`
	Count int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *SaleGetListResp) Reset() {
	*x = SaleGetListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleGetListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleGetListResp) ProtoMessage() {}

func (x *SaleGetListResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleGetListResp.ProtoReflect.Descriptor instead.
func (*SaleGetListResp) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{4}
}

func (x *SaleGetListResp) GetSales() []*Sale {
	if x != nil {
		return x.Sales
	}
	return nil
}

func (x *SaleGetListResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type SaleIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaleIdReq) Reset() {
	*x = SaleIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleIdReq) ProtoMessage() {}

func (x *SaleIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleIdReq.ProtoReflect.Descriptor instead.
func (*SaleIdReq) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{5}
}

func (x *SaleIdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SaleUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BranchId        string  `protobuf:"bytes,2,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	ShopAssistantId string  `protobuf:"bytes,3,opt,name=shop_assistant_id,json=shopAssistantId,proto3" json:"shop_assistant_id,omitempty"`
	CashierId       string  `protobuf:"bytes,4,opt,name=cashier_id,json=cashierId,proto3" json:"cashier_id,omitempty"`
	Price           float32 `protobuf:"fixed32,5,opt,name=price,proto3" json:"price,omitempty"`
	PaymentType     string  `protobuf:"bytes,6,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	Status          string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	ClientName      string  `protobuf:"bytes,8,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
}

func (x *SaleUpdateReq) Reset() {
	*x = SaleUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleUpdateReq) ProtoMessage() {}

func (x *SaleUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleUpdateReq.ProtoReflect.Descriptor instead.
func (*SaleUpdateReq) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{6}
}

func (x *SaleUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SaleUpdateReq) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *SaleUpdateReq) GetShopAssistantId() string {
	if x != nil {
		return x.ShopAssistantId
	}
	return ""
}

func (x *SaleUpdateReq) GetCashierId() string {
	if x != nil {
		return x.CashierId
	}
	return ""
}

func (x *SaleUpdateReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SaleUpdateReq) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *SaleUpdateReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SaleUpdateReq) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

type SaleUpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SaleUpdateResp) Reset() {
	*x = SaleUpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleUpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleUpdateResp) ProtoMessage() {}

func (x *SaleUpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleUpdateResp.ProtoReflect.Descriptor instead.
func (*SaleUpdateResp) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{7}
}

func (x *SaleUpdateResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SaleDeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SaleDeleteResp) Reset() {
	*x = SaleDeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sale_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaleDeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaleDeleteResp) ProtoMessage() {}

func (x *SaleDeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_sale_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaleDeleteResp.ProtoReflect.Descriptor instead.
func (*SaleDeleteResp) Descriptor() ([]byte, []int) {
	return file_sale_proto_rawDescGZIP(), []int{8}
}

func (x *SaleDeleteResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sale_proto protoreflect.FileDescriptor

var file_sale_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x61,
	0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xcd, 0x02, 0x0a, 0x04, 0x53,
	0x61, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x6f,
	0x70, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xe9, 0x01, 0x0a, 0x0d, 0x53,
	0x61, 0x6c, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f,
	0x70, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x73, 0x68, 0x69,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x53, 0x61, 0x6c, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x84, 0x03, 0x0a, 0x0e, 0x53,
	0x61, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x70,
	0x5f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x5f, 0x74, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x54, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f,
	0x74, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x54,
	0x6f, 0x22, 0x51, 0x0a, 0x0f, 0x53, 0x61, 0x6c, 0x65, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0xf9, 0x01, 0x0a, 0x0d, 0x53, 0x61, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x68, 0x6f,
	0x70, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a,
	0x0e, 0x53, 0x61, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x22, 0x0a, 0x0e, 0x53, 0x61, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sale_proto_rawDescOnce sync.Once
	file_sale_proto_rawDescData = file_sale_proto_rawDesc
)

func file_sale_proto_rawDescGZIP() []byte {
	file_sale_proto_rawDescOnce.Do(func() {
		file_sale_proto_rawDescData = protoimpl.X.CompressGZIP(file_sale_proto_rawDescData)
	})
	return file_sale_proto_rawDescData
}

var file_sale_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_sale_proto_goTypes = []interface{}{
	(*Sale)(nil),            // 0: sale_service.Sale
	(*SaleCreateReq)(nil),   // 1: sale_service.SaleCreateReq
	(*SaleCreateResp)(nil),  // 2: sale_service.SaleCreateResp
	(*SaleGetListReq)(nil),  // 3: sale_service.SaleGetListReq
	(*SaleGetListResp)(nil), // 4: sale_service.SaleGetListResp
	(*SaleIdReq)(nil),       // 5: sale_service.SaleIdReq
	(*SaleUpdateReq)(nil),   // 6: sale_service.SaleUpdateReq
	(*SaleUpdateResp)(nil),  // 7: sale_service.SaleUpdateResp
	(*SaleDeleteResp)(nil),  // 8: sale_service.SaleDeleteResp
}
var file_sale_proto_depIdxs = []int32{
	0, // 0: sale_service.SaleGetListResp.sales:type_name -> sale_service.Sale
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sale_proto_init() }
func file_sale_proto_init() {
	if File_sale_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sale_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sale); i {
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
		file_sale_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleCreateReq); i {
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
		file_sale_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleCreateResp); i {
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
		file_sale_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleGetListReq); i {
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
		file_sale_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleGetListResp); i {
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
		file_sale_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleIdReq); i {
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
		file_sale_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleUpdateReq); i {
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
		file_sale_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleUpdateResp); i {
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
		file_sale_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaleDeleteResp); i {
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
			RawDescriptor: file_sale_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sale_proto_goTypes,
		DependencyIndexes: file_sale_proto_depIdxs,
		MessageInfos:      file_sale_proto_msgTypes,
	}.Build()
	File_sale_proto = out.File
	file_sale_proto_rawDesc = nil
	file_sale_proto_goTypes = nil
	file_sale_proto_depIdxs = nil
}