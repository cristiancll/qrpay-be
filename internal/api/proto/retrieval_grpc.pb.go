// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: retrieval.proto

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

// RetrievalServiceClient is the client API for RetrievalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RetrievalServiceClient interface {
	Create(ctx context.Context, in *RetrievalCreateRequest, opts ...grpc.CallOption) (*RetrievalCreateResponse, error)
	Get(ctx context.Context, in *RetrievalGetRequest, opts ...grpc.CallOption) (*RetrievalGetResponse, error)
	List(ctx context.Context, in *RetrievalListRequest, opts ...grpc.CallOption) (*RetrievalListResponse, error)
	Update(ctx context.Context, in *RetrievalUpdateRequest, opts ...grpc.CallOption) (*RetrievalUpdateResponse, error)
	Delete(ctx context.Context, in *RetrievalDeleteRequest, opts ...grpc.CallOption) (*RetrievalDeleteResponse, error)
}

type retrievalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRetrievalServiceClient(cc grpc.ClientConnInterface) RetrievalServiceClient {
	return &retrievalServiceClient{cc}
}

func (c *retrievalServiceClient) Create(ctx context.Context, in *RetrievalCreateRequest, opts ...grpc.CallOption) (*RetrievalCreateResponse, error) {
	out := new(RetrievalCreateResponse)
	err := c.cc.Invoke(ctx, "/proto.RetrievalService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalServiceClient) Get(ctx context.Context, in *RetrievalGetRequest, opts ...grpc.CallOption) (*RetrievalGetResponse, error) {
	out := new(RetrievalGetResponse)
	err := c.cc.Invoke(ctx, "/proto.RetrievalService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalServiceClient) List(ctx context.Context, in *RetrievalListRequest, opts ...grpc.CallOption) (*RetrievalListResponse, error) {
	out := new(RetrievalListResponse)
	err := c.cc.Invoke(ctx, "/proto.RetrievalService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalServiceClient) Update(ctx context.Context, in *RetrievalUpdateRequest, opts ...grpc.CallOption) (*RetrievalUpdateResponse, error) {
	out := new(RetrievalUpdateResponse)
	err := c.cc.Invoke(ctx, "/proto.RetrievalService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *retrievalServiceClient) Delete(ctx context.Context, in *RetrievalDeleteRequest, opts ...grpc.CallOption) (*RetrievalDeleteResponse, error) {
	out := new(RetrievalDeleteResponse)
	err := c.cc.Invoke(ctx, "/proto.RetrievalService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RetrievalServiceServer is the server API for RetrievalService service.
// All implementations must embed UnimplementedRetrievalServiceServer
// for forward compatibility
type RetrievalServiceServer interface {
	Create(context.Context, *RetrievalCreateRequest) (*RetrievalCreateResponse, error)
	Get(context.Context, *RetrievalGetRequest) (*RetrievalGetResponse, error)
	List(context.Context, *RetrievalListRequest) (*RetrievalListResponse, error)
	Update(context.Context, *RetrievalUpdateRequest) (*RetrievalUpdateResponse, error)
	Delete(context.Context, *RetrievalDeleteRequest) (*RetrievalDeleteResponse, error)
	mustEmbedUnimplementedRetrievalServiceServer()
}

// UnimplementedRetrievalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRetrievalServiceServer struct {
}

func (UnimplementedRetrievalServiceServer) Create(context.Context, *RetrievalCreateRequest) (*RetrievalCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRetrievalServiceServer) Get(context.Context, *RetrievalGetRequest) (*RetrievalGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRetrievalServiceServer) List(context.Context, *RetrievalListRequest) (*RetrievalListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRetrievalServiceServer) Update(context.Context, *RetrievalUpdateRequest) (*RetrievalUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRetrievalServiceServer) Delete(context.Context, *RetrievalDeleteRequest) (*RetrievalDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRetrievalServiceServer) mustEmbedUnimplementedRetrievalServiceServer() {}

// UnsafeRetrievalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RetrievalServiceServer will
// result in compilation errors.
type UnsafeRetrievalServiceServer interface {
	mustEmbedUnimplementedRetrievalServiceServer()
}

func RegisterRetrievalServiceServer(s grpc.ServiceRegistrar, srv RetrievalServiceServer) {
	s.RegisterService(&RetrievalService_ServiceDesc, srv)
}

func _RetrievalService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrievalCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RetrievalService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServiceServer).Create(ctx, req.(*RetrievalCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RetrievalService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrievalGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RetrievalService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServiceServer).Get(ctx, req.(*RetrievalGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RetrievalService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrievalListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RetrievalService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServiceServer).List(ctx, req.(*RetrievalListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RetrievalService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrievalUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RetrievalService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServiceServer).Update(ctx, req.(*RetrievalUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RetrievalService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrievalDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RetrievalServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RetrievalService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RetrievalServiceServer).Delete(ctx, req.(*RetrievalDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RetrievalService_ServiceDesc is the grpc.ServiceDesc for RetrievalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RetrievalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RetrievalService",
	HandlerType: (*RetrievalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RetrievalService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RetrievalService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RetrievalService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RetrievalService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RetrievalService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "retrieval.proto",
}
