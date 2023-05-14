// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: operationLog.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OperationLogServiceClient is the client API for OperationLogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationLogServiceClient interface {
	Create(ctx context.Context, in *OperationLogCreateRequest, opts ...grpc.CallOption) (*OperationLogCreateResponse, error)
	Get(ctx context.Context, in *OperationLogGetRequest, opts ...grpc.CallOption) (*OperationLogGetResponse, error)
	List(ctx context.Context, in *OperationLogListRequest, opts ...grpc.CallOption) (*OperationLogListResponse, error)
	Update(ctx context.Context, in *OperationLogUpdateRequest, opts ...grpc.CallOption) (*OperationLogUpdateResponse, error)
	Delete(ctx context.Context, in *OperationLogDeleteRequest, opts ...grpc.CallOption) (*OperationLogDeleteResponse, error)
}

type operationLogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationLogServiceClient(cc grpc.ClientConnInterface) OperationLogServiceClient {
	return &operationLogServiceClient{cc}
}

func (c *operationLogServiceClient) Create(ctx context.Context, in *OperationLogCreateRequest, opts ...grpc.CallOption) (*OperationLogCreateResponse, error) {
	out := new(OperationLogCreateResponse)
	err := c.cc.Invoke(ctx, "/proto.OperationLogService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationLogServiceClient) Get(ctx context.Context, in *OperationLogGetRequest, opts ...grpc.CallOption) (*OperationLogGetResponse, error) {
	out := new(OperationLogGetResponse)
	err := c.cc.Invoke(ctx, "/proto.OperationLogService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationLogServiceClient) List(ctx context.Context, in *OperationLogListRequest, opts ...grpc.CallOption) (*OperationLogListResponse, error) {
	out := new(OperationLogListResponse)
	err := c.cc.Invoke(ctx, "/proto.OperationLogService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationLogServiceClient) Update(ctx context.Context, in *OperationLogUpdateRequest, opts ...grpc.CallOption) (*OperationLogUpdateResponse, error) {
	out := new(OperationLogUpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.OperationLogService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationLogServiceClient) Delete(ctx context.Context, in *OperationLogDeleteRequest, opts ...grpc.CallOption) (*OperationLogDeleteResponse, error) {
	out := new(OperationLogDeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.OperationLogService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationLogServiceServer is the server API for OperationLogService service.
// All implementations must embed UnimplementedOperationLogServiceServer
// for forward compatibility
type OperationLogServiceServer interface {
	Create(context.Context, *OperationLogCreateRequest) (*OperationLogCreateResponse, error)
	Get(context.Context, *OperationLogGetRequest) (*OperationLogGetResponse, error)
	List(context.Context, *OperationLogListRequest) (*OperationLogListResponse, error)
	Update(context.Context, *OperationLogUpdateRequest) (*OperationLogUpdateResponse, error)
	Delete(context.Context, *OperationLogDeleteRequest) (*OperationLogDeleteResponse, error)
	mustEmbedUnimplementedOperationLogServiceServer()
}

// UnimplementedOperationLogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationLogServiceServer struct {
}

func (UnimplementedOperationLogServiceServer) Create(context.Context, *OperationLogCreateRequest) (*OperationLogCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOperationLogServiceServer) Get(context.Context, *OperationLogGetRequest) (*OperationLogGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOperationLogServiceServer) List(context.Context, *OperationLogListRequest) (*OperationLogListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOperationLogServiceServer) Update(context.Context, *OperationLogUpdateRequest) (*OperationLogUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOperationLogServiceServer) Delete(context.Context, *OperationLogDeleteRequest) (*OperationLogDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOperationLogServiceServer) mustEmbedUnimplementedOperationLogServiceServer() {}

// UnsafeOperationLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationLogServiceServer will
// result in compilation errors.
type UnsafeOperationLogServiceServer interface {
	mustEmbedUnimplementedOperationLogServiceServer()
}

func RegisterOperationLogServiceServer(s grpc.ServiceRegistrar, srv OperationLogServiceServer) {
	s.RegisterService(&OperationLogService_ServiceDesc, srv)
}

func _OperationLogService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationLogCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OperationLogService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServiceServer).Create(ctx, req.(*OperationLogCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationLogService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationLogGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OperationLogService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServiceServer).Get(ctx, req.(*OperationLogGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationLogService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationLogListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OperationLogService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServiceServer).List(ctx, req.(*OperationLogListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationLogService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationLogUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OperationLogService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServiceServer).Update(ctx, req.(*OperationLogUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationLogService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationLogDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OperationLogService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServiceServer).Delete(ctx, req.(*OperationLogDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationLogService_ServiceDesc is the grpc.ServiceDesc for OperationLogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationLogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OperationLogService",
	HandlerType: (*OperationLogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OperationLogService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _OperationLogService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OperationLogService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OperationLogService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OperationLogService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operationLog.proto",
}