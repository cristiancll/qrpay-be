package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/cristiancll/qrpay-be/internal/roles"
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

func checkAdminAuthorization(ctx context.Context) error {
	ctxRole := ctx.Value("Role")
	if ctxRole == nil {
		return status.Error(codes.Unauthenticated, errors.UNAUTHORIZED)
	}

	stringRole, ok := ctxRole.(string)
	if !ok {
		return status.Error(codes.Unauthenticated, errors.UNAUTHORIZED)
	}

	intRole, err := strconv.ParseInt(stringRole, 10, 64)
	if err != nil {
		return status.Error(codes.Unauthenticated, errors.UNAUTHORIZED)
	}

	role := roles.Role(intRole)

	if !role.IsAdmin() {
		return status.Error(codes.PermissionDenied, errors.PERMISSION_DENIED)
	}

	return nil
}
