package handler

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/errCode"
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
		return "", errs.New(errors.New(""), errCode.Internal)
	}

	stringUUID, ok := ctxUUID.(string)
	if !ok {
		return "", errs.New(errors.New(""), errCode.Internal)
	}

	if err := checkValidUUID(stringUUID); err != nil {
		return "", errs.Wrap(err, "")
	}

	return stringUUID, nil
}

func getRoleFromContext(ctx context.Context) (roles.Role, error) {
	ctxRole := ctx.Value("Role")
	if ctxRole == nil {
		return 0, errs.New(errors.New(""), errCode.Internal)
	}

	stringRole, ok := ctxRole.(string)
	if !ok {
		return 0, errs.New(errors.New(""), errCode.Internal)
	}

	intRole, err := strconv.ParseInt(stringRole, 10, 64)
	if err != nil {
		return 0, errs.New(errors.New(""), errCode.Internal)
	}

	return roles.Role(intRole), nil
}

func checkAdminAuthorization(ctx context.Context) error {
	role, err := getRoleFromContext(ctx)
	if err != nil {
		return errs.Wrap(err, "")
	}
	if !role.IsAdmin() {
		return errs.New(errors.New(""), errCode.Unauthorized)
	}
	return nil
}

func checkStaffAuthorization(ctx context.Context) error {
	role, err := getRoleFromContext(ctx)
	if err != nil {
		return errs.Wrap(err, "")
	}
	if !role.IsAdmin() && !role.IsBilling() && !role.IsManager() && !role.IsSeller() {
		return errs.New(errors.New(""), errCode.Unauthorized)
	}
	return nil
}

func checkValidUUID(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errs.New(errors.New(""), errCode.InvalidArgument)
	}
	return nil
}
