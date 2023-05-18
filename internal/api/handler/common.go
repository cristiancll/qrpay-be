package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type Creater[E any, REQ any, RES any] interface {
	Create(context.Context, *REQ) (*RES, error)
}
type Getter[E any, REQ any, RES any] interface {
	Get(context.Context, *REQ) (*RES, error)
}
type Lister[E any, REQ any, RES any] interface {
	List(context.Context, *REQ) (*RES, error)
}
type Updater[E any, REQ any, RES any] interface {
	Update(context.Context, *REQ) (*RES, error)
}
type Deleter[E any, REQ any, RES any] interface {
	Delete(context.Context, *REQ) (*RES, error)
}

func getRoleFromContext(ctx context.Context) (roles.Role, error) {
	ctxRole := ctx.Value("Role")
	if ctxRole == nil {
		return 0, status.Error(codes.Unauthenticated, errors.UNAUTHORIZED)
	}

	stringRole, ok := ctxRole.(string)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, errors.UNAUTHORIZED)
	}

	intRole, err := strconv.ParseInt(stringRole, 10, 64)
	if err != nil {
		return 0, status.Error(codes.Unauthenticated, errors.UNAUTHORIZED)
	}

	return roles.Role(intRole), nil
}

func checkAdminAuthorization(ctx context.Context) error {
	role, err := getRoleFromContext(ctx)
	if err != nil {
		return err
	}
	if !role.IsAdmin() {
		return status.Error(codes.PermissionDenied, errors.PERMISSION_DENIED)
	}
	return nil
}

func checkStaffAuthorization(ctx context.Context) error {
	role, err := getRoleFromContext(ctx)
	if err != nil {
		return err
	}
	if !role.IsAdmin() && !role.IsBilling() && !role.IsManager() && !role.IsSeller() {
		return status.Error(codes.PermissionDenied, errors.PERMISSION_DENIED)
	}
	return nil
}

func checkValidUUID(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return status.Error(codes.InvalidArgument, errors.INVALID_UUID)
	}
	return nil
}
