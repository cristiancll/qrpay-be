package middleware

import (
	"context"
	"fmt"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ErrorUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err == nil {
		return resp, err
	}
	customErr, ok := err.(*errs.Error)
	if !ok {
		return resp, err
	}
	errorCode := errCode.ToGRPCCode(customErr.Code)
	message := customErr.Message
	fmt.Print(customErr.Error())
	return resp, status.Error(errorCode, message)
}

func ErrorStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	err := handler(srv, stream)
	if err == nil {
		return err
	}
	customErr, ok := err.(*errs.Error)
	if !ok {
		return err
	}
	errorCode := errCode.ToGRPCCode(customErr.Code)
	message := customErr.Message
	fmt.Printf("Error: %v\n", customErr.Error())
	return status.Error(errorCode, message)
}
