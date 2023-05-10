package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler interface {
	Create(ctx context.Context, req *proto.UserCreateRequest) (*proto.UserCreateResponse, error)
	Get(ctx context.Context, req *proto.UserGetRequest) (*proto.UserGetResponse, error)
	List(ctx context.Context, req *proto.UserListRequest) (*proto.UserListResponse, error)
	Update(ctx context.Context, req *proto.UserUpdateRequest) (*proto.UserUpdateResponse, error)
	Delete(ctx context.Context, req *proto.UserDeleteRequest) (*proto.UserDeleteResponse, error)
	AdminCreated(ctx context.Context, req *proto.UserAdminCreatedRequest) (*proto.UserAdminCreatedResponse, error)
	proto.UserServiceServer
}

type userHandler struct {
	service service.UserService
	proto.UnimplementedUserServiceServer
}

func NewUserHandler(s service.UserService) UserHandler {
	return &userHandler{service: s}
}

func (h *userHandler) Create(ctx context.Context, req *proto.UserCreateRequest) (*proto.UserCreateResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "Password is required")
	}
	if req.Phone == "" {
		return nil, status.Error(codes.InvalidArgument, "Phone is required")
	}
	user := &model.User{
		Name:  req.Name,
		Phone: req.Phone,
	}
	err := h.service.Create(ctx, user, req.Password)
	if err != nil {
		return nil, err
	}
	res := &proto.UserCreateResponse{
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

func (h *userHandler) Get(ctx context.Context, req *proto.UserGetRequest) (*proto.UserGetResponse, error) {
	res := &proto.UserGetResponse{}

	return res, nil
}

func (h *userHandler) List(ctx context.Context, req *proto.UserListRequest) (*proto.UserListResponse, error) {
	res := &proto.UserListResponse{}

	return res, nil
}

func (h *userHandler) Update(ctx context.Context, req *proto.UserUpdateRequest) (*proto.UserUpdateResponse, error) {
	res := &proto.UserUpdateResponse{}

	return res, nil
}

func (h *userHandler) Delete(ctx context.Context, req *proto.UserDeleteRequest) (*proto.UserDeleteResponse, error) {
	res := &proto.UserDeleteResponse{}

	return res, nil
}

func (h *userHandler) AdminCreated(ctx context.Context, req *proto.UserAdminCreatedRequest) (*proto.UserAdminCreatedResponse, error) {
	res := &proto.UserAdminCreatedResponse{}

	return res, nil
}
