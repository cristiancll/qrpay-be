package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func RateLimiterUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return handler(ctx, req)
}

func RateLimiterStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, stream)
}
