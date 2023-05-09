package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func LoggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	p, _ := peer.FromContext(ctx)
	fmt.Printf("[U] %s | %s\n", p.Addr, info.FullMethod)
	return handler(ctx, req)
}

func LoggingStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	p, _ := peer.FromContext(stream.Context())
	fmt.Printf("[S] %s | %s\n", p.Addr, info.FullMethod)
	return handler(srv, stream)
}
