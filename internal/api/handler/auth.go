package handler

import (
	"context"
	"encoding/json"
	"github.com/cristiancll/qrpay-be/configs"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/cristiancll/qrpay-be/internal/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthHandler interface {
	Login(ctx context.Context, req *proto.AuthLoginRequest) (*proto.AuthLoginResponse, error)
	Logout(ctx context.Context, req *proto.AuthVoid) (*proto.AuthVoid, error)
	Heartbeat(ctx context.Context, req *proto.AuthVoid) (*proto.AuthHeartbeatResponse, error)
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
		return nil, status.Error(codes.InvalidArgument, errors.PHONE_REQUIRED)
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errors.PASSWORD_REQUIRED)
	}
	user, auth, err := h.service.Login(ctx, req.Phone, req.Password)
	if err != nil {
		return nil, err
	}

	privateKey := configs.Get().Keys.JWT.PrivateKey
	subj, err := json.Marshal(security.SubjectClaims{
		UUID: user.UUID,
		Role: string(user.Role),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	token, err := security.GenerateJWTToken(string(subj[:]), privateKey)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.AUTH_ERROR)
	}

	err = security.UpdateJWTCookie(ctx, token)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.AUTH_ERROR)
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
		Auth: &proto.Auth{
			Verified: auth.Verified,
			Disabled: auth.Disabled,
		},
	}
	return res, nil
}

func (h *authHandler) Logout(ctx context.Context, req *proto.AuthVoid) (*proto.AuthVoid, error) {
	err := security.DeleteJWTCookie(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.AUTH_ERROR)
	}
	res := &proto.AuthVoid{}
	return res, nil
}

func (h *authHandler) Heartbeat(ctx context.Context, req *proto.AuthVoid) (*proto.AuthHeartbeatResponse, error) {
	user, auth, err := h.service.Heartbeat(ctx)
	if err != nil {
		return nil, err
	}
	res := &proto.AuthHeartbeatResponse{
		User: &proto.User{
			Uuid:      user.UUID,
			Name:      user.Name,
			Role:      int64(user.Role),
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
		Auth: &proto.Auth{
			Verified: auth.Verified,
			Disabled: auth.Disabled,
		},
	}
	return res, nil
}
