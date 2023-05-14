// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: sku.proto

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

type SKU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Item        *Item                  `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Name        string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64                  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SKU) Reset() {
	*x = SKU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKU) ProtoMessage() {}

func (x *SKU) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKU.ProtoReflect.Descriptor instead.
func (*SKU) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{0}
}

func (x *SKU) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SKU) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *SKU) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SKU) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SKU) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SKU) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SKU) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type SKUCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item        *Item  `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SKUCreateRequest) Reset() {
	*x = SKUCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUCreateRequest) ProtoMessage() {}

func (x *SKUCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUCreateRequest.ProtoReflect.Descriptor instead.
func (*SKUCreateRequest) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{1}
}

func (x *SKUCreateRequest) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *SKUCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SKUCreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SKUCreateRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SKUCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku *SKU `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *SKUCreateResponse) Reset() {
	*x = SKUCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUCreateResponse) ProtoMessage() {}

func (x *SKUCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUCreateResponse.ProtoReflect.Descriptor instead.
func (*SKUCreateResponse) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{2}
}

func (x *SKUCreateResponse) GetSku() *SKU {
	if x != nil {
		return x.Sku
	}
	return nil
}

type SKUGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *SKUGetRequest) Reset() {
	*x = SKUGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUGetRequest) ProtoMessage() {}

func (x *SKUGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUGetRequest.ProtoReflect.Descriptor instead.
func (*SKUGetRequest) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{3}
}

func (x *SKUGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type SKUGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku *SKU `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *SKUGetResponse) Reset() {
	*x = SKUGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUGetResponse) ProtoMessage() {}

func (x *SKUGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUGetResponse.ProtoReflect.Descriptor instead.
func (*SKUGetResponse) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{4}
}

func (x *SKUGetResponse) GetSku() *SKU {
	if x != nil {
		return x.Sku
	}
	return nil
}

type SKUUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Item        *Item  `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Price       int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SKUUpdateRequest) Reset() {
	*x = SKUUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUUpdateRequest) ProtoMessage() {}

func (x *SKUUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUUpdateRequest.ProtoReflect.Descriptor instead.
func (*SKUUpdateRequest) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{5}
}

func (x *SKUUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *SKUUpdateRequest) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *SKUUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SKUUpdateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SKUUpdateRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SKUUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sku *SKU `protobuf:"bytes,1,opt,name=sku,proto3" json:"sku,omitempty"`
}

func (x *SKUUpdateResponse) Reset() {
	*x = SKUUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUUpdateResponse) ProtoMessage() {}

func (x *SKUUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUUpdateResponse.ProtoReflect.Descriptor instead.
func (*SKUUpdateResponse) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{6}
}

func (x *SKUUpdateResponse) GetSku() *SKU {
	if x != nil {
		return x.Sku
	}
	return nil
}

type SKUDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *SKUDeleteRequest) Reset() {
	*x = SKUDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUDeleteRequest) ProtoMessage() {}

func (x *SKUDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUDeleteRequest.ProtoReflect.Descriptor instead.
func (*SKUDeleteRequest) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{7}
}

func (x *SKUDeleteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type SKUDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SKUDeleteResponse) Reset() {
	*x = SKUDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUDeleteResponse) ProtoMessage() {}

func (x *SKUDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUDeleteResponse.ProtoReflect.Descriptor instead.
func (*SKUDeleteResponse) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{8}
}

type SKUListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SKUListRequest) Reset() {
	*x = SKUListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUListRequest) ProtoMessage() {}

func (x *SKUListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUListRequest.ProtoReflect.Descriptor instead.
func (*SKUListRequest) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{9}
}

type SKUListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skus []*SKU `protobuf:"bytes,1,rep,name=skus,proto3" json:"skus,omitempty"`
}

func (x *SKUListResponse) Reset() {
	*x = SKUListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sku_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SKUListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SKUListResponse) ProtoMessage() {}

func (x *SKUListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sku_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SKUListResponse.ProtoReflect.Descriptor instead.
func (*SKUListResponse) Descriptor() ([]byte, []int) {
	return file_sku_proto_rawDescGZIP(), []int{10}
}

func (x *SKUListResponse) GetSkus() []*SKU {
	if x != nil {
		return x.Skus
	}
	return nil
}

var File_sku_proto protoreflect.FileDescriptor

var file_sku_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x6b, 0x75, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x01, 0x0a, 0x03, 0x53, 0x4b, 0x55, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7f,
	0x0a, 0x10, 0x53, 0x4b, 0x55, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69,
	0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x31, 0x0a, 0x11, 0x53, 0x4b, 0x55, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x52, 0x03, 0x73,
	0x6b, 0x75, 0x22, 0x23, 0x0a, 0x0d, 0x53, 0x4b, 0x55, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x0e, 0x53, 0x4b, 0x55, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x73, 0x6b, 0x75,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x4b, 0x55, 0x52, 0x03, 0x73, 0x6b, 0x75, 0x22, 0x93, 0x01, 0x0a, 0x10, 0x53, 0x4b, 0x55, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x1f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x31, 0x0a,
	0x11, 0x53, 0x4b, 0x55, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x73, 0x6b, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x52, 0x03, 0x73, 0x6b, 0x75,
	0x22, 0x26, 0x0a, 0x10, 0x53, 0x4b, 0x55, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x4b, 0x55, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x0a,
	0x0e, 0x53, 0x4b, 0x55, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x31, 0x0a, 0x0f, 0x53, 0x4b, 0x55, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x52, 0x04, 0x73, 0x6b,
	0x75, 0x73, 0x32, 0xae, 0x02, 0x0a, 0x0a, 0x53, 0x4b, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b,
	0x55, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b,
	0x55, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x4b, 0x55, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x72, 0x69, 0x73, 0x74, 0x69, 0x61, 0x6e, 0x63, 0x6c, 0x6c, 0x2f, 0x71, 0x72,
	0x70, 0x61, 0x79, 0x2d, 0x62, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sku_proto_rawDescOnce sync.Once
	file_sku_proto_rawDescData = file_sku_proto_rawDesc
)

func file_sku_proto_rawDescGZIP() []byte {
	file_sku_proto_rawDescOnce.Do(func() {
		file_sku_proto_rawDescData = protoimpl.X.CompressGZIP(file_sku_proto_rawDescData)
	})
	return file_sku_proto_rawDescData
}

var file_sku_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_sku_proto_goTypes = []interface{}{
	(*SKU)(nil),                   // 0: proto.SKU
	(*SKUCreateRequest)(nil),      // 1: proto.SKUCreateRequest
	(*SKUCreateResponse)(nil),     // 2: proto.SKUCreateResponse
	(*SKUGetRequest)(nil),         // 3: proto.SKUGetRequest
	(*SKUGetResponse)(nil),        // 4: proto.SKUGetResponse
	(*SKUUpdateRequest)(nil),      // 5: proto.SKUUpdateRequest
	(*SKUUpdateResponse)(nil),     // 6: proto.SKUUpdateResponse
	(*SKUDeleteRequest)(nil),      // 7: proto.SKUDeleteRequest
	(*SKUDeleteResponse)(nil),     // 8: proto.SKUDeleteResponse
	(*SKUListRequest)(nil),        // 9: proto.SKUListRequest
	(*SKUListResponse)(nil),       // 10: proto.SKUListResponse
	(*Item)(nil),                  // 11: proto.Item
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_sku_proto_depIdxs = []int32{
	11, // 0: proto.SKU.item:type_name -> proto.Item
	12, // 1: proto.SKU.created_at:type_name -> google.protobuf.Timestamp
	12, // 2: proto.SKU.updated_at:type_name -> google.protobuf.Timestamp
	11, // 3: proto.SKUCreateRequest.item:type_name -> proto.Item
	0,  // 4: proto.SKUCreateResponse.sku:type_name -> proto.SKU
	0,  // 5: proto.SKUGetResponse.sku:type_name -> proto.SKU
	11, // 6: proto.SKUUpdateRequest.item:type_name -> proto.Item
	0,  // 7: proto.SKUUpdateResponse.sku:type_name -> proto.SKU
	0,  // 8: proto.SKUListResponse.skus:type_name -> proto.SKU
	1,  // 9: proto.SKUService.Create:input_type -> proto.SKUCreateRequest
	3,  // 10: proto.SKUService.Get:input_type -> proto.SKUGetRequest
	5,  // 11: proto.SKUService.Update:input_type -> proto.SKUUpdateRequest
	7,  // 12: proto.SKUService.Delete:input_type -> proto.SKUDeleteRequest
	9,  // 13: proto.SKUService.List:input_type -> proto.SKUListRequest
	2,  // 14: proto.SKUService.Create:output_type -> proto.SKUCreateResponse
	4,  // 15: proto.SKUService.Get:output_type -> proto.SKUGetResponse
	6,  // 16: proto.SKUService.Update:output_type -> proto.SKUUpdateResponse
	8,  // 17: proto.SKUService.Delete:output_type -> proto.SKUDeleteResponse
	10, // 18: proto.SKUService.List:output_type -> proto.SKUListResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_sku_proto_init() }
func file_sku_proto_init() {
	if File_sku_proto != nil {
		return
	}
	file_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sku_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKU); i {
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
		file_sku_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUCreateRequest); i {
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
		file_sku_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUCreateResponse); i {
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
		file_sku_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUGetRequest); i {
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
		file_sku_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUGetResponse); i {
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
		file_sku_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUUpdateRequest); i {
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
		file_sku_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUUpdateResponse); i {
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
		file_sku_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUDeleteRequest); i {
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
		file_sku_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUDeleteResponse); i {
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
		file_sku_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUListRequest); i {
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
		file_sku_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SKUListResponse); i {
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
			RawDescriptor: file_sku_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sku_proto_goTypes,
		DependencyIndexes: file_sku_proto_depIdxs,
		MessageInfos:      file_sku_proto_msgTypes,
	}.Build()
	File_sku_proto = out.File
	file_sku_proto_rawDesc = nil
	file_sku_proto_goTypes = nil
	file_sku_proto_depIdxs = nil
}