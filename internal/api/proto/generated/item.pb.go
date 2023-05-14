// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: item.proto

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

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Category  *Category              `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Name      string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Item) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Item) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ItemCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Name     string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ItemCreateRequest) Reset() {
	*x = ItemCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCreateRequest) ProtoMessage() {}

func (x *ItemCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCreateRequest.ProtoReflect.Descriptor instead.
func (*ItemCreateRequest) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{1}
}

func (x *ItemCreateRequest) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ItemCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ItemCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *ItemCreateResponse) Reset() {
	*x = ItemCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemCreateResponse) ProtoMessage() {}

func (x *ItemCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemCreateResponse.ProtoReflect.Descriptor instead.
func (*ItemCreateResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{2}
}

func (x *ItemCreateResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type ItemUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string    `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Category *Category `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Name     string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ItemUpdateRequest) Reset() {
	*x = ItemUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemUpdateRequest) ProtoMessage() {}

func (x *ItemUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemUpdateRequest.ProtoReflect.Descriptor instead.
func (*ItemUpdateRequest) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{3}
}

func (x *ItemUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *ItemUpdateRequest) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ItemUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ItemUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *ItemUpdateResponse) Reset() {
	*x = ItemUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemUpdateResponse) ProtoMessage() {}

func (x *ItemUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemUpdateResponse.ProtoReflect.Descriptor instead.
func (*ItemUpdateResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{4}
}

func (x *ItemUpdateResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type ItemDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ItemDeleteRequest) Reset() {
	*x = ItemDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDeleteRequest) ProtoMessage() {}

func (x *ItemDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDeleteRequest.ProtoReflect.Descriptor instead.
func (*ItemDeleteRequest) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{5}
}

func (x *ItemDeleteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ItemDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ItemDeleteResponse) Reset() {
	*x = ItemDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDeleteResponse) ProtoMessage() {}

func (x *ItemDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDeleteResponse.ProtoReflect.Descriptor instead.
func (*ItemDeleteResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{6}
}

type ItemGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ItemGetRequest) Reset() {
	*x = ItemGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemGetRequest) ProtoMessage() {}

func (x *ItemGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemGetRequest.ProtoReflect.Descriptor instead.
func (*ItemGetRequest) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{7}
}

func (x *ItemGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ItemGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *Item `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *ItemGetResponse) Reset() {
	*x = ItemGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemGetResponse) ProtoMessage() {}

func (x *ItemGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemGetResponse.ProtoReflect.Descriptor instead.
func (*ItemGetResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{8}
}

func (x *ItemGetResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type ItemListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ItemListRequest) Reset() {
	*x = ItemListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemListRequest) ProtoMessage() {}

func (x *ItemListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemListRequest.ProtoReflect.Descriptor instead.
func (*ItemListRequest) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{9}
}

type ItemListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ItemListResponse) Reset() {
	*x = ItemListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_item_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemListResponse) ProtoMessage() {}

func (x *ItemListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_item_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemListResponse.ProtoReflect.Descriptor instead.
func (*ItemListResponse) Descriptor() ([]byte, []int) {
	return file_item_proto_rawDescGZIP(), []int{10}
}

func (x *ItemListResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_item_proto protoreflect.FileDescriptor

var file_item_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x54, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35,
	0x0a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x68, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x2b,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x35, 0x0a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x27, 0x0a, 0x11, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22,
	0x14, 0x0a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0f, 0x49,
	0x74, 0x65, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22,
	0x11, 0x0a, 0x0f, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x35, 0x0a, 0x10, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xb9, 0x02, 0x0a, 0x0b, 0x49, 0x74,
	0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x69, 0x73, 0x74, 0x69, 0x61, 0x6e, 0x63, 0x6c, 0x6c, 0x2f,
	0x71, 0x72, 0x70, 0x61, 0x79, 0x2d, 0x62, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_item_proto_rawDescOnce sync.Once
	file_item_proto_rawDescData = file_item_proto_rawDesc
)

func file_item_proto_rawDescGZIP() []byte {
	file_item_proto_rawDescOnce.Do(func() {
		file_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_item_proto_rawDescData)
	})
	return file_item_proto_rawDescData
}

var file_item_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_item_proto_goTypes = []interface{}{
	(*Item)(nil),                  // 0: proto.Item
	(*ItemCreateRequest)(nil),     // 1: proto.ItemCreateRequest
	(*ItemCreateResponse)(nil),    // 2: proto.ItemCreateResponse
	(*ItemUpdateRequest)(nil),     // 3: proto.ItemUpdateRequest
	(*ItemUpdateResponse)(nil),    // 4: proto.ItemUpdateResponse
	(*ItemDeleteRequest)(nil),     // 5: proto.ItemDeleteRequest
	(*ItemDeleteResponse)(nil),    // 6: proto.ItemDeleteResponse
	(*ItemGetRequest)(nil),        // 7: proto.ItemGetRequest
	(*ItemGetResponse)(nil),       // 8: proto.ItemGetResponse
	(*ItemListRequest)(nil),       // 9: proto.ItemListRequest
	(*ItemListResponse)(nil),      // 10: proto.ItemListResponse
	(*Category)(nil),              // 11: proto.Category
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_item_proto_depIdxs = []int32{
	11, // 0: proto.Item.category:type_name -> proto.Category
	12, // 1: proto.Item.created_at:type_name -> google.protobuf.Timestamp
	12, // 2: proto.Item.updated_at:type_name -> google.protobuf.Timestamp
	11, // 3: proto.ItemCreateRequest.category:type_name -> proto.Category
	0,  // 4: proto.ItemCreateResponse.item:type_name -> proto.Item
	11, // 5: proto.ItemUpdateRequest.category:type_name -> proto.Category
	0,  // 6: proto.ItemUpdateResponse.item:type_name -> proto.Item
	0,  // 7: proto.ItemGetResponse.item:type_name -> proto.Item
	0,  // 8: proto.ItemListResponse.items:type_name -> proto.Item
	1,  // 9: proto.ItemService.Create:input_type -> proto.ItemCreateRequest
	3,  // 10: proto.ItemService.Update:input_type -> proto.ItemUpdateRequest
	5,  // 11: proto.ItemService.Delete:input_type -> proto.ItemDeleteRequest
	7,  // 12: proto.ItemService.Get:input_type -> proto.ItemGetRequest
	9,  // 13: proto.ItemService.List:input_type -> proto.ItemListRequest
	2,  // 14: proto.ItemService.Create:output_type -> proto.ItemCreateResponse
	4,  // 15: proto.ItemService.Update:output_type -> proto.ItemUpdateResponse
	6,  // 16: proto.ItemService.Delete:output_type -> proto.ItemDeleteResponse
	8,  // 17: proto.ItemService.Get:output_type -> proto.ItemGetResponse
	10, // 18: proto.ItemService.List:output_type -> proto.ItemListResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_item_proto_init() }
func file_item_proto_init() {
	if File_item_proto != nil {
		return
	}
	file_category_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCreateRequest); i {
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
		file_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemCreateResponse); i {
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
		file_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemUpdateRequest); i {
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
		file_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemUpdateResponse); i {
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
		file_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDeleteRequest); i {
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
		file_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDeleteResponse); i {
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
		file_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemGetRequest); i {
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
		file_item_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemGetResponse); i {
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
		file_item_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemListRequest); i {
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
		file_item_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemListResponse); i {
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
			RawDescriptor: file_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_item_proto_goTypes,
		DependencyIndexes: file_item_proto_depIdxs,
		MessageInfos:      file_item_proto_msgTypes,
	}.Build()
	File_item_proto = out.File
	file_item_proto_rawDesc = nil
	file_item_proto_goTypes = nil
	file_item_proto_depIdxs = nil
}