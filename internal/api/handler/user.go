package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User interface {
	Creater[model.User, proto.UserCreateRequest, proto.UserCreateResponse]
	Getter[model.User, proto.UserGetRequest, proto.UserGetResponse]
	Lister[model.User, proto.UserListRequest, proto.UserListResponse]
	Updater[model.User, proto.UserUpdateRequest, proto.UserUpdateResponse]
	Deleter[model.User, proto.UserDeleteRequest, proto.UserDeleteResponse]
	AdminCreated(ctx context.Context, req *proto.UserAdminCreatedRequest) (*proto.UserAdminCreatedResponse, error)
	proto.UserServiceServer
}

type user struct {
	service service.User
	proto.UnimplementedUserServiceServer
}

func NewUser(s service.User) User {
	return &user{service: s}
}

func (h *user) Create(ctx context.Context, req *proto.UserCreateRequest) (*proto.UserCreateResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errors.NAME_REQUIRED)
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errors.PASSWORD_REQUIRED)
	}
	if req.Phone == "" {
		return nil, status.Error(codes.InvalidArgument, errors.PHONE_REQUIRED)
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

func (h *user) Get(ctx context.Context, req *proto.UserGetRequest) (*proto.UserGetResponse, error) {
	res := &proto.UserGetResponse{}

	return res, nil
}

func (h *user) List(ctx context.Context, req *proto.UserListRequest) (*proto.UserListResponse, error) {
	res := &proto.UserListResponse{}

	return res, nil
}

func (h *user) Update(ctx context.Context, req *proto.UserUpdateRequest) (*proto.UserUpdateResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errors.NAME_REQUIRED)
	}
	if req.Phone == "" {
		return nil, status.Error(codes.InvalidArgument, errors.PHONE_REQUIRED)
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, errors.PASSWORD_REQUIRED)
	}
	UUID := ctx.Value("UUID").(string)
	user := &model.User{
		UUID:  UUID,
		Name:  req.Name,
		Phone: req.Phone,
	}
	err := h.service.Update(ctx, user, req.Password)
	if err != nil {
		return nil, err
	}
	res := &proto.UserUpdateResponse{
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

func (h *user) Delete(ctx context.Context, req *proto.UserDeleteRequest) (*proto.UserDeleteResponse, error) {
	res := &proto.UserDeleteResponse{}

	return res, nil
}

func (h *user) AdminCreated(ctx context.Context, req *proto.UserAdminCreatedRequest) (*proto.UserAdminCreatedResponse, error) {
	res := &proto.UserAdminCreatedResponse{}

	return res, nil
}
