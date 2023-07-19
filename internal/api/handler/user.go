package handler

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User interface {
	Creater[model.User, proto.UserCreateRequest, proto.UserCreateResponse]
	Getter[model.User, proto.UserGetRequest, proto.UserGetResponse]
	Lister[model.User, proto.UserListRequest, proto.UserListResponse]
	Updater[model.User, proto.UserUpdateRequest, proto.UserUpdateResponse]
	Deleter[model.User, proto.UserDeleteRequest, proto.UserDeleteResponse]
	AdminCreated(ctx context.Context, req *proto.UserAdminCreatedRequest) (*proto.UserAdminCreatedResponse, error)
	UpdateRole(ctx context.Context, req *proto.UserUpdateRoleRequest) (*proto.UserUpdateRoleResponse, error)
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
		return nil, errs.New(errors.New(errMsg.NameRequired), errCode.InvalidArgument)
	}
	if req.Password == "" {
		return nil, errs.New(errors.New(errMsg.PasswordRequired), errCode.InvalidArgument)
	}
	if req.Phone == "" {
		return nil, errs.New(errors.New(errMsg.PhoneRequired), errCode.InvalidArgument)
	}
	user, err := h.service.Create(ctx, req.Name, req.Phone, req.Password)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedCreateUser, req.Name, req.Phone) // Don't log password
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
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}
	err = checkValidUUID(req.Uuid)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.UUIDInvalid)
	}
	user, err := h.service.Get(ctx, req.Uuid)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetUser, req.Uuid)
	}

	res := &proto.UserGetResponse{
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

func (h *user) List(ctx context.Context, req *proto.UserListRequest) (*proto.UserListResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}

	users, err := h.service.List(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetAllUser)
	}
	res := &proto.UserListResponse{}
	for _, user := range users {
		res.Users = append(res.Users, &proto.User{
			Uuid:      user.UUID,
			Name:      user.Name,
			Role:      int64(user.Role),
			Phone:     user.Phone,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}
	return res, nil
}

func (h *user) Update(ctx context.Context, req *proto.UserUpdateRequest) (*proto.UserUpdateResponse, error) {
	if req.Name == "" {
		return nil, errs.New(errors.New(errMsg.NameRequired), errCode.InvalidArgument)
	}
	if req.Phone == "" {
		return nil, errs.New(errors.New(errMsg.PhoneRequired), errCode.InvalidArgument)
	}
	if req.Password == "" {
		return nil, errs.New(errors.New(errMsg.PasswordRequired), errCode.InvalidArgument)
	}
	UUID, err := getUUIDFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedUUIDContext)
	}
	user := &model.User{
		UUID:  UUID,
		Name:  req.Name,
		Phone: req.Phone,
	}
	err = h.service.Update(ctx, user, req.Password)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedUpdateUser) // Don't log password
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
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}
	sellerUUID, err := getUUIDFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedUUIDContext)
	}
	if req.Name == "" {
		return nil, errs.New(errors.New(errMsg.NameRequired), errCode.InvalidArgument)
	}
	if req.Phone == "" {
		return nil, errs.New(errors.New(errMsg.PhoneRequired), errCode.InvalidArgument)
	}
	user, err := h.service.AdminCreated(ctx, req.Name, req.Phone, sellerUUID)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedCreateUser, req.Name, req.Phone, sellerUUID)
	}
	res := &proto.UserAdminCreatedResponse{
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

func (h *user) UpdateRole(ctx context.Context, req *proto.UserUpdateRoleRequest) (*proto.UserUpdateRoleResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}

	if req.Uuid == "" {
		return nil, errs.New(errors.New(errMsg.UUIDRequired), errCode.InvalidArgument)
	}
	if req.Role == 0 || req.Role < 0 {
		return nil, errs.New(errors.New(errMsg.RoleRequired), errCode.InvalidArgument)
	}
	role := roles.Role(req.Role)

	user, err := h.service.UpdateRole(ctx, req.Uuid, role, req.Enabled)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedUpdateUser, req.Uuid, role, req.Enabled)
	}

	res := &proto.UserUpdateRoleResponse{
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
