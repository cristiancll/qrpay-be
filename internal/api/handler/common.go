package handler

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"github.com/google/uuid"
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

func getUUIDFromContext(ctx context.Context) (string, error) {
	ctxUUID := ctx.Value("UUID")
	if ctxUUID == nil {
		return "", errs.New(errors.New(errMsg.FailedUUIDContext), errCode.Internal)
	}

	stringUUID, ok := ctxUUID.(string)
	if !ok {
		return "", errs.New(errors.New(errMsg.FailedStringConversion), errCode.Internal, stringUUID)
	}

	if err := checkValidUUID(stringUUID); err != nil {
		return "", errs.Wrap(err, errMsg.UUIDInvalid, stringUUID)
	}

	return stringUUID, nil
}

func getRoleFromContext(ctx context.Context) (roles.Role, error) {
	ctxRole := ctx.Value("Role")
	if ctxRole == nil {
		return 0, errs.New(errors.New(errMsg.FailedRoleContext), errCode.Internal)
	}

	stringRole, ok := ctxRole.(string)
	if !ok {
		return 0, errs.New(errors.New(errMsg.FailedStringConversion), errCode.Internal, stringRole)
	}

	intRole, err := strconv.ParseInt(stringRole, 10, 64)
	if err != nil {
		return 0, errs.New(errors.New(errMsg.NotAnInteger), errCode.Internal, stringRole)
	}

	return roles.Role(intRole), nil
}

func checkAdminAuthorization(ctx context.Context) error {
	role, err := getRoleFromContext(ctx)
	if err != nil {
		return errs.Wrap(err, errMsg.RoleNotFound)
	}
	if !role.IsAdmin() {
		return errs.New(errors.New(errMsg.UserNotAdmin), errCode.Unauthorized, role)
	}
	return nil
}

func checkStaffAuthorization(ctx context.Context) error {
	role, err := getRoleFromContext(ctx)
	if err != nil {
		return errs.Wrap(err, errMsg.RoleNotFound)
	}
	if !role.IsAdmin() && !role.IsBilling() && !role.IsManager() && !role.IsSeller() {
		return errs.New(errors.New(errMsg.UserNotStaff), errCode.Unauthorized, role)
	}
	return nil
}

func checkValidUUID(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errs.New(errors.New(errMsg.UUIDInvalid), errCode.InvalidArgument, id)
	}
	return nil
}
