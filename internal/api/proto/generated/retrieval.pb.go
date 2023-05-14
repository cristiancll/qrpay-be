// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: retrieval.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Retrieval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User      *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Seller    *User                  `protobuf:"bytes,3,opt,name=seller,proto3" json:"seller,omitempty"`
	SaleItem  *SaleItem              `protobuf:"bytes,4,opt,name=sale_item,json=saleItem,proto3" json:"sale_item,omitempty"`
	Quantity  int64                  `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Retrieval) Reset() {
	*x = Retrieval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Retrieval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Retrieval) ProtoMessage() {}

func (x *Retrieval) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Retrieval.ProtoReflect.Descriptor instead.
func (*Retrieval) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{0}
}

func (x *Retrieval) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Retrieval) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Retrieval) GetSeller() *User {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *Retrieval) GetSaleItem() *SaleItem {
	if x != nil {
		return x.SaleItem
	}
	return nil
}

func (x *Retrieval) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Retrieval) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Retrieval) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type RetrievalCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     *User     `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Seller   *User     `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	SaleItem *SaleItem `protobuf:"bytes,3,opt,name=sale_item,json=saleItem,proto3" json:"sale_item,omitempty"`
	Quantity int64     `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *RetrievalCreateRequest) Reset() {
	*x = RetrievalCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalCreateRequest) ProtoMessage() {}

func (x *RetrievalCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalCreateRequest.ProtoReflect.Descriptor instead.
func (*RetrievalCreateRequest) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{1}
}

func (x *RetrievalCreateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RetrievalCreateRequest) GetSeller() *User {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *RetrievalCreateRequest) GetSaleItem() *SaleItem {
	if x != nil {
		return x.SaleItem
	}
	return nil
}

func (x *RetrievalCreateRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RetrievalCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retrieval *Retrieval `protobuf:"bytes,1,opt,name=retrieval,proto3" json:"retrieval,omitempty"`
}

func (x *RetrievalCreateResponse) Reset() {
	*x = RetrievalCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalCreateResponse) ProtoMessage() {}

func (x *RetrievalCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalCreateResponse.ProtoReflect.Descriptor instead.
func (*RetrievalCreateResponse) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{2}
}

func (x *RetrievalCreateResponse) GetRetrieval() *Retrieval {
	if x != nil {
		return x.Retrieval
	}
	return nil
}

type RetrievalGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *RetrievalGetRequest) Reset() {
	*x = RetrievalGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalGetRequest) ProtoMessage() {}

func (x *RetrievalGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalGetRequest.ProtoReflect.Descriptor instead.
func (*RetrievalGetRequest) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{3}
}

func (x *RetrievalGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type RetrievalGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retrieval *Retrieval `protobuf:"bytes,1,opt,name=retrieval,proto3" json:"retrieval,omitempty"`
}

func (x *RetrievalGetResponse) Reset() {
	*x = RetrievalGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalGetResponse) ProtoMessage() {}

func (x *RetrievalGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalGetResponse.ProtoReflect.Descriptor instead.
func (*RetrievalGetResponse) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{4}
}

func (x *RetrievalGetResponse) GetRetrieval() *Retrieval {
	if x != nil {
		return x.Retrieval
	}
	return nil
}

type RetrievalListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetrievalListRequest) Reset() {
	*x = RetrievalListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalListRequest) ProtoMessage() {}

func (x *RetrievalListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalListRequest.ProtoReflect.Descriptor instead.
func (*RetrievalListRequest) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{5}
}

type RetrievalListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retrievals []*Retrieval `protobuf:"bytes,1,rep,name=retrievals,proto3" json:"retrievals,omitempty"`
}

func (x *RetrievalListResponse) Reset() {
	*x = RetrievalListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalListResponse) ProtoMessage() {}

func (x *RetrievalListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalListResponse.ProtoReflect.Descriptor instead.
func (*RetrievalListResponse) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{6}
}

func (x *RetrievalListResponse) GetRetrievals() []*Retrieval {
	if x != nil {
		return x.Retrievals
	}
	return nil
}

type RetrievalUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User     *User     `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Seller   *User     `protobuf:"bytes,3,opt,name=seller,proto3" json:"seller,omitempty"`
	SaleItem *SaleItem `protobuf:"bytes,4,opt,name=sale_item,json=saleItem,proto3" json:"sale_item,omitempty"`
	Quantity int64     `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *RetrievalUpdateRequest) Reset() {
	*x = RetrievalUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalUpdateRequest) ProtoMessage() {}

func (x *RetrievalUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalUpdateRequest.ProtoReflect.Descriptor instead.
func (*RetrievalUpdateRequest) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{7}
}

func (x *RetrievalUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *RetrievalUpdateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RetrievalUpdateRequest) GetSeller() *User {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *RetrievalUpdateRequest) GetSaleItem() *SaleItem {
	if x != nil {
		return x.SaleItem
	}
	return nil
}

func (x *RetrievalUpdateRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type RetrievalUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retrieval *Retrieval `protobuf:"bytes,1,opt,name=retrieval,proto3" json:"retrieval,omitempty"`
}

func (x *RetrievalUpdateResponse) Reset() {
	*x = RetrievalUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalUpdateResponse) ProtoMessage() {}

func (x *RetrievalUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalUpdateResponse.ProtoReflect.Descriptor instead.
func (*RetrievalUpdateResponse) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{8}
}

func (x *RetrievalUpdateResponse) GetRetrieval() *Retrieval {
	if x != nil {
		return x.Retrieval
	}
	return nil
}

type RetrievalDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *RetrievalDeleteRequest) Reset() {
	*x = RetrievalDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalDeleteRequest) ProtoMessage() {}

func (x *RetrievalDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalDeleteRequest.ProtoReflect.Descriptor instead.
func (*RetrievalDeleteRequest) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{9}
}

func (x *RetrievalDeleteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type RetrievalDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RetrievalDeleteResponse) Reset() {
	*x = RetrievalDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalDeleteResponse) ProtoMessage() {}

func (x *RetrievalDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalDeleteResponse.ProtoReflect.Descriptor instead.
func (*RetrievalDeleteResponse) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{10}
}

var File_retrieval_proto protoreflect.FileDescriptor

var file_retrieval_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2c, 0x0a,
	0x09, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa8, 0x01,
	0x0a, 0x16, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x49, 0x0a, 0x17, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x22, 0x29, 0x0a, 0x13, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x46,
	0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x52, 0x09, 0x72, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x49,
	0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x52, 0x0a, 0x72,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x16, 0x52, 0x65,
	0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x09, 0x73, 0x61, 0x6c, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x6c, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x08, 0x73, 0x61, 0x6c, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x49, 0x0a, 0x17, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x52, 0x09, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x16, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x22, 0x19, 0x0a, 0x17, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf0, 0x02, 0x0a,
	0x10, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x47, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x61, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61,
	0x6c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x61, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61,
	0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72,
	0x69, 0x73, 0x74, 0x69, 0x61, 0x6e, 0x63, 0x6c, 0x6c, 0x2f, 0x71, 0x72, 0x70, 0x61, 0x79, 0x2d,
	0x62, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_retrieval_proto_rawDescOnce sync.Once
	file_retrieval_proto_rawDescData = file_retrieval_proto_rawDesc
)

func file_retrieval_proto_rawDescGZIP() []byte {
	file_retrieval_proto_rawDescOnce.Do(func() {
		file_retrieval_proto_rawDescData = protoimpl.X.CompressGZIP(file_retrieval_proto_rawDescData)
	})
	return file_retrieval_proto_rawDescData
}

var file_retrieval_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_retrieval_proto_goTypes = []interface{}{
	(*Retrieval)(nil),               // 0: proto.Retrieval
	(*RetrievalCreateRequest)(nil),  // 1: proto.RetrievalCreateRequest
	(*RetrievalCreateResponse)(nil), // 2: proto.RetrievalCreateResponse
	(*RetrievalGetRequest)(nil),     // 3: proto.RetrievalGetRequest
	(*RetrievalGetResponse)(nil),    // 4: proto.RetrievalGetResponse
	(*RetrievalListRequest)(nil),    // 5: proto.RetrievalListRequest
	(*RetrievalListResponse)(nil),   // 6: proto.RetrievalListResponse
	(*RetrievalUpdateRequest)(nil),  // 7: proto.RetrievalUpdateRequest
	(*RetrievalUpdateResponse)(nil), // 8: proto.RetrievalUpdateResponse
	(*RetrievalDeleteRequest)(nil),  // 9: proto.RetrievalDeleteRequest
	(*RetrievalDeleteResponse)(nil), // 10: proto.RetrievalDeleteResponse
	(*User)(nil),                    // 11: proto.User
	(*SaleItem)(nil),                // 12: proto.SaleItem
	(*timestamppb.Timestamp)(nil),   // 13: google.protobuf.Timestamp
}
var file_retrieval_proto_depIdxs = []int32{
	11, // 0: proto.Retrieval.user:type_name -> proto.User
	11, // 1: proto.Retrieval.seller:type_name -> proto.User
	12, // 2: proto.Retrieval.sale_item:type_name -> proto.SaleItem
	13, // 3: proto.Retrieval.created_at:type_name -> google.protobuf.Timestamp
	13, // 4: proto.Retrieval.updated_at:type_name -> google.protobuf.Timestamp
	11, // 5: proto.RetrievalCreateRequest.user:type_name -> proto.User
	11, // 6: proto.RetrievalCreateRequest.seller:type_name -> proto.User
	12, // 7: proto.RetrievalCreateRequest.sale_item:type_name -> proto.SaleItem
	0,  // 8: proto.RetrievalCreateResponse.retrieval:type_name -> proto.Retrieval
	0,  // 9: proto.RetrievalGetResponse.retrieval:type_name -> proto.Retrieval
	0,  // 10: proto.RetrievalListResponse.retrievals:type_name -> proto.Retrieval
	11, // 11: proto.RetrievalUpdateRequest.user:type_name -> proto.User
	11, // 12: proto.RetrievalUpdateRequest.seller:type_name -> proto.User
	12, // 13: proto.RetrievalUpdateRequest.sale_item:type_name -> proto.SaleItem
	0,  // 14: proto.RetrievalUpdateResponse.retrieval:type_name -> proto.Retrieval
	1,  // 15: proto.RetrievalService.Create:input_type -> proto.RetrievalCreateRequest
	3,  // 16: proto.RetrievalService.Get:input_type -> proto.RetrievalGetRequest
	5,  // 17: proto.RetrievalService.List:input_type -> proto.RetrievalListRequest
	7,  // 18: proto.RetrievalService.Update:input_type -> proto.RetrievalUpdateRequest
	9,  // 19: proto.RetrievalService.Delete:input_type -> proto.RetrievalDeleteRequest
	2,  // 20: proto.RetrievalService.Create:output_type -> proto.RetrievalCreateResponse
	4,  // 21: proto.RetrievalService.Get:output_type -> proto.RetrievalGetResponse
	6,  // 22: proto.RetrievalService.List:output_type -> proto.RetrievalListResponse
	8,  // 23: proto.RetrievalService.Update:output_type -> proto.RetrievalUpdateResponse
	10, // 24: proto.RetrievalService.Delete:output_type -> proto.RetrievalDeleteResponse
	20, // [20:25] is the sub-list for method output_type
	15, // [15:20] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_retrieval_proto_init() }
func file_retrieval_proto_init() {
	if File_retrieval_proto != nil {
		return
	}
	file_user_proto_init()
	file_saleItem_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_retrieval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Retrieval); i {
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
		file_retrieval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalCreateRequest); i {
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
		file_retrieval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalCreateResponse); i {
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
		file_retrieval_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalGetRequest); i {
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
		file_retrieval_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalGetResponse); i {
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
		file_retrieval_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalListRequest); i {
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
		file_retrieval_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalListResponse); i {
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
		file_retrieval_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalUpdateRequest); i {
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
		file_retrieval_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalUpdateResponse); i {
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
		file_retrieval_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalDeleteRequest); i {
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
		file_retrieval_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalDeleteResponse); i {
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
			RawDescriptor: file_retrieval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_retrieval_proto_goTypes,
		DependencyIndexes: file_retrieval_proto_depIdxs,
		MessageInfos:      file_retrieval_proto_msgTypes,
	}.Build()
	File_retrieval_proto = out.File
	file_retrieval_proto_rawDesc = nil
	file_retrieval_proto_goTypes = nil
	file_retrieval_proto_depIdxs = nil
}
