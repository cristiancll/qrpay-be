package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/configs"
	"github.com/cristiancll/qrpay-be/internal/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthHandler interface {
	Login(ctx context.Context, req *proto.AuthLoginRequest) (*proto.AuthLoginResponse, error)
	Logout(ctx context.Context, req *proto.AuthVoid) (*proto.AuthVoid, error)
	proto.AuthServiceServer
}

type authHandler struct {
	service service.AuthService
	proto.UnimplementedAuthServiceServer
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &authHandler{
		service: service,
	}
}

func (h *authHandler) Login(ctx context.Context, req *proto.AuthLoginRequest) (*proto.AuthLoginResponse, error) {
	if req.Phone == "" {
		return nil, status.Error(codes.InvalidArgument, "User is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Password is required")
	}
	user, err := h.service.Login(ctx, req.Phone, req.Password)
	if err != nil {
		return nil, err
	}

	privateKey := configs.Get().Keys.JWT.PrivateKey
	token, err := security.GenerateJWTToken(user.UUID, privateKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error generating token: %v", err)
	}

	err = security.UpdateJWTCookie(ctx, token)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error updating cookie: %v", err)
	}
	res := &proto.AuthLoginResponse{
		User: &proto.User{
			Uuid:      user.UUID,
			Name:      user.Name,
			Role:      int64(user.Role),
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}
	return res, nil
}

func (h *authHandler) Logout(ctx context.Context, req *proto.AuthVoid) (*proto.AuthVoid, error) {
	err := security.DeleteJWTCookie(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error deleting cookie: %v", err)
	}
	res := &proto.AuthVoid{}
	return res, nil
}
