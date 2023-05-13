// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: operationLog.proto

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

type OperationLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User        *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Seller      *User                  `protobuf:"bytes,3,opt,name=seller,proto3" json:"seller,omitempty"`
	Operation   string                 `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	OperationId int64                  `protobuf:"varint,5,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	Metadata    string                 `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *OperationLog) Reset() {
	*x = OperationLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLog) ProtoMessage() {}

func (x *OperationLog) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLog.ProtoReflect.Descriptor instead.
func (*OperationLog) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{0}
}

func (x *OperationLog) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *OperationLog) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *OperationLog) GetSeller() *User {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *OperationLog) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *OperationLog) GetOperationId() int64 {
	if x != nil {
		return x.OperationId
	}
	return 0
}

func (x *OperationLog) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *OperationLog) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OperationLog) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type OperationLogCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Seller      *User  `protobuf:"bytes,2,opt,name=seller,proto3" json:"seller,omitempty"`
	Operation   string `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	OperationId int64  `protobuf:"varint,4,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	Metadata    string `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *OperationLogCreateRequest) Reset() {
	*x = OperationLogCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogCreateRequest) ProtoMessage() {}

func (x *OperationLogCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogCreateRequest.ProtoReflect.Descriptor instead.
func (*OperationLogCreateRequest) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{1}
}

func (x *OperationLogCreateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *OperationLogCreateRequest) GetSeller() *User {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *OperationLogCreateRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *OperationLogCreateRequest) GetOperationId() int64 {
	if x != nil {
		return x.OperationId
	}
	return 0
}

func (x *OperationLogCreateRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type OperationLogCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationLog *OperationLog `protobuf:"bytes,1,opt,name=operation_log,json=operationLog,proto3" json:"operation_log,omitempty"`
}

func (x *OperationLogCreateResponse) Reset() {
	*x = OperationLogCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogCreateResponse) ProtoMessage() {}

func (x *OperationLogCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogCreateResponse.ProtoReflect.Descriptor instead.
func (*OperationLogCreateResponse) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{2}
}

func (x *OperationLogCreateResponse) GetOperationLog() *OperationLog {
	if x != nil {
		return x.OperationLog
	}
	return nil
}

type OperationLogGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *OperationLogGetRequest) Reset() {
	*x = OperationLogGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogGetRequest) ProtoMessage() {}

func (x *OperationLogGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogGetRequest.ProtoReflect.Descriptor instead.
func (*OperationLogGetRequest) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{3}
}

func (x *OperationLogGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type OperationLogGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationLog *OperationLog `protobuf:"bytes,1,opt,name=operation_log,json=operationLog,proto3" json:"operation_log,omitempty"`
}

func (x *OperationLogGetResponse) Reset() {
	*x = OperationLogGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogGetResponse) ProtoMessage() {}

func (x *OperationLogGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogGetResponse.ProtoReflect.Descriptor instead.
func (*OperationLogGetResponse) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{4}
}

func (x *OperationLogGetResponse) GetOperationLog() *OperationLog {
	if x != nil {
		return x.OperationLog
	}
	return nil
}

type OperationLogListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OperationLogListRequest) Reset() {
	*x = OperationLogListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogListRequest) ProtoMessage() {}

func (x *OperationLogListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogListRequest.ProtoReflect.Descriptor instead.
func (*OperationLogListRequest) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{5}
}

type OperationLogListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationLogs []*OperationLog `protobuf:"bytes,1,rep,name=operation_logs,json=operationLogs,proto3" json:"operation_logs,omitempty"`
}

func (x *OperationLogListResponse) Reset() {
	*x = OperationLogListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogListResponse) ProtoMessage() {}

func (x *OperationLogListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogListResponse.ProtoReflect.Descriptor instead.
func (*OperationLogListResponse) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{6}
}

func (x *OperationLogListResponse) GetOperationLogs() []*OperationLog {
	if x != nil {
		return x.OperationLogs
	}
	return nil
}

type OperationLogUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	User        *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Seller      *User  `protobuf:"bytes,3,opt,name=seller,proto3" json:"seller,omitempty"`
	Operation   string `protobuf:"bytes,4,opt,name=operation,proto3" json:"operation,omitempty"`
	OperationId int64  `protobuf:"varint,5,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	Metadata    string `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *OperationLogUpdateRequest) Reset() {
	*x = OperationLogUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogUpdateRequest) ProtoMessage() {}

func (x *OperationLogUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogUpdateRequest.ProtoReflect.Descriptor instead.
func (*OperationLogUpdateRequest) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{7}
}

func (x *OperationLogUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *OperationLogUpdateRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *OperationLogUpdateRequest) GetSeller() *User {
	if x != nil {
		return x.Seller
	}
	return nil
}

func (x *OperationLogUpdateRequest) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *OperationLogUpdateRequest) GetOperationId() int64 {
	if x != nil {
		return x.OperationId
	}
	return 0
}

func (x *OperationLogUpdateRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type OperationLogUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationLog *OperationLog `protobuf:"bytes,1,opt,name=operation_log,json=operationLog,proto3" json:"operation_log,omitempty"`
}

func (x *OperationLogUpdateResponse) Reset() {
	*x = OperationLogUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogUpdateResponse) ProtoMessage() {}

func (x *OperationLogUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogUpdateResponse.ProtoReflect.Descriptor instead.
func (*OperationLogUpdateResponse) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{8}
}

func (x *OperationLogUpdateResponse) GetOperationLog() *OperationLog {
	if x != nil {
		return x.OperationLog
	}
	return nil
}

type OperationLogDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *OperationLogDeleteRequest) Reset() {
	*x = OperationLogDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogDeleteRequest) ProtoMessage() {}

func (x *OperationLogDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogDeleteRequest.ProtoReflect.Descriptor instead.
func (*OperationLogDeleteRequest) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{9}
}

func (x *OperationLogDeleteRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type OperationLogDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OperationLogDeleteResponse) Reset() {
	*x = OperationLogDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_operationLog_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLogDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLogDeleteResponse) ProtoMessage() {}

func (x *OperationLogDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operationLog_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLogDeleteResponse.ProtoReflect.Descriptor instead.
func (*OperationLogDeleteResponse) Descriptor() ([]byte, []int) {
	return file_operationLog_proto_rawDescGZIP(), []int{10}
}

var File_operationLog_proto protoreflect.FileDescriptor

var file_operationLog_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x0c, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23,
	0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x1a, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x22,
	0x2c, 0x0a, 0x16, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x53, 0x0a,
	0x17, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x22, 0x19, 0x0a, 0x17, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a,
	0x18, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x19, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x1a, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x22, 0x2f, 0x0a, 0x19, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x91, 0x03, 0x0a, 0x13, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x72, 0x69, 0x73, 0x74, 0x69, 0x61, 0x6e, 0x63, 0x6c, 0x6c, 0x2f,
	0x71, 0x72, 0x70, 0x61, 0x79, 0x2d, 0x62, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_operationLog_proto_rawDescOnce sync.Once
	file_operationLog_proto_rawDescData = file_operationLog_proto_rawDesc
)

func file_operationLog_proto_rawDescGZIP() []byte {
	file_operationLog_proto_rawDescOnce.Do(func() {
		file_operationLog_proto_rawDescData = protoimpl.X.CompressGZIP(file_operationLog_proto_rawDescData)
	})
	return file_operationLog_proto_rawDescData
}

var file_operationLog_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_operationLog_proto_goTypes = []interface{}{
	(*OperationLog)(nil),               // 0: proto.OperationLog
	(*OperationLogCreateRequest)(nil),  // 1: proto.OperationLogCreateRequest
	(*OperationLogCreateResponse)(nil), // 2: proto.OperationLogCreateResponse
	(*OperationLogGetRequest)(nil),     // 3: proto.OperationLogGetRequest
	(*OperationLogGetResponse)(nil),    // 4: proto.OperationLogGetResponse
	(*OperationLogListRequest)(nil),    // 5: proto.OperationLogListRequest
	(*OperationLogListResponse)(nil),   // 6: proto.OperationLogListResponse
	(*OperationLogUpdateRequest)(nil),  // 7: proto.OperationLogUpdateRequest
	(*OperationLogUpdateResponse)(nil), // 8: proto.OperationLogUpdateResponse
	(*OperationLogDeleteRequest)(nil),  // 9: proto.OperationLogDeleteRequest
	(*OperationLogDeleteResponse)(nil), // 10: proto.OperationLogDeleteResponse
	(*User)(nil),                       // 11: proto.User
	(*timestamppb.Timestamp)(nil),      // 12: google.protobuf.Timestamp
}
var file_operationLog_proto_depIdxs = []int32{
	11, // 0: proto.OperationLog.user:type_name -> proto.User
	11, // 1: proto.OperationLog.seller:type_name -> proto.User
	12, // 2: proto.OperationLog.created_at:type_name -> google.protobuf.Timestamp
	12, // 3: proto.OperationLog.updated_at:type_name -> google.protobuf.Timestamp
	11, // 4: proto.OperationLogCreateRequest.user:type_name -> proto.User
	11, // 5: proto.OperationLogCreateRequest.seller:type_name -> proto.User
	0,  // 6: proto.OperationLogCreateResponse.operation_log:type_name -> proto.OperationLog
	0,  // 7: proto.OperationLogGetResponse.operation_log:type_name -> proto.OperationLog
	0,  // 8: proto.OperationLogListResponse.operation_logs:type_name -> proto.OperationLog
	11, // 9: proto.OperationLogUpdateRequest.user:type_name -> proto.User
	11, // 10: proto.OperationLogUpdateRequest.seller:type_name -> proto.User
	0,  // 11: proto.OperationLogUpdateResponse.operation_log:type_name -> proto.OperationLog
	1,  // 12: proto.OperationLogService.Create:input_type -> proto.OperationLogCreateRequest
	3,  // 13: proto.OperationLogService.Get:input_type -> proto.OperationLogGetRequest
	5,  // 14: proto.OperationLogService.List:input_type -> proto.OperationLogListRequest
	7,  // 15: proto.OperationLogService.Update:input_type -> proto.OperationLogUpdateRequest
	9,  // 16: proto.OperationLogService.Delete:input_type -> proto.OperationLogDeleteRequest
	2,  // 17: proto.OperationLogService.Create:output_type -> proto.OperationLogCreateResponse
	4,  // 18: proto.OperationLogService.Get:output_type -> proto.OperationLogGetResponse
	6,  // 19: proto.OperationLogService.List:output_type -> proto.OperationLogListResponse
	8,  // 20: proto.OperationLogService.Update:output_type -> proto.OperationLogUpdateResponse
	10, // 21: proto.OperationLogService.Delete:output_type -> proto.OperationLogDeleteResponse
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_operationLog_proto_init() }
func file_operationLog_proto_init() {
	if File_operationLog_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_operationLog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLog); i {
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
		file_operationLog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogCreateRequest); i {
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
		file_operationLog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogCreateResponse); i {
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
		file_operationLog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogGetRequest); i {
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
		file_operationLog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogGetResponse); i {
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
		file_operationLog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogListRequest); i {
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
		file_operationLog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogListResponse); i {
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
		file_operationLog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogUpdateRequest); i {
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
		file_operationLog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogUpdateResponse); i {
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
		file_operationLog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogDeleteRequest); i {
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
		file_operationLog_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationLogDeleteResponse); i {
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
			RawDescriptor: file_operationLog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_operationLog_proto_goTypes,
		DependencyIndexes: file_operationLog_proto_depIdxs,
		MessageInfos:      file_operationLog_proto_msgTypes,
	}.Build()
	File_operationLog_proto = out.File
	file_operationLog_proto_rawDesc = nil
	file_operationLog_proto_goTypes = nil
	file_operationLog_proto_depIdxs = nil
}
