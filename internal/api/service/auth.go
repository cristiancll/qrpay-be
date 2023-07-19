package service

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/common"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/cristiancll/qrpay-be/internal/security"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Auth interface {
	Login(ctx context.Context, username string, password string) (*model.User, *model.Auth, error)
	Heartbeat(ctx context.Context) (*model.User, *model.Auth, error)
}

type auth struct {
	pool      *pgxpool.Pool
	repo      repository.Auth
	opLogRepo repository.OperationLog
	userRepo  repository.User
}

func NewAuth(pool *pgxpool.Pool, r repository.Auth, userRepo repository.User, opLogRepo repository.OperationLog) Auth {
	return &auth{
		pool:      pool,
		repo:      r,
		userRepo:  userRepo,
		opLogRepo: opLogRepo,
	}
}

func (s *auth) Login(ctx context.Context, phone string, password string) (*model.User, *model.Auth, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)
	sanitizedPhone := common.SanitizePhone(phone)
	user, err := s.userRepo.GetUserByPhone(ctx, tx, sanitizedPhone)
	if err != nil {
		return nil, nil, errs.Wrap(err, errMsg.FailedGetUser, sanitizedPhone)
	}
	auth, err := s.repo.TGetById(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, errs.Wrap(err, errMsg.FailedGetAuth, user.ID)
	}
	if !auth.Verified {
		return nil, nil, errs.New(errors.New(errMsg.UserNotVerified), errCode.AccessDenied)
	}
	if auth.Disabled {
		return nil, nil, errs.New(errors.New(errMsg.UserDisabled), errCode.AccessDenied)
	}
	err = security.CheckPassword(auth.Password, password)
	if err != nil {
		return nil, nil, errs.Wrap(err, errMsg.IncorrectPassword)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, nil, errs.New(err, errCode.Internal)
	}
	return user, auth, nil
}

func (s *auth) Heartbeat(ctx context.Context) (*model.User, *model.Auth, error) {
	UUID := ctx.Value("UUID").(string)
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)
	user, err := s.userRepo.TGetByUUID(ctx, tx, UUID)
	if err != nil {
		return nil, nil, errs.Wrap(err, errMsg.FailedGetUser, UUID)
	}
	auth, err := s.repo.TGetById(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, errs.Wrap(err, errMsg.FailedGetAuth, user.ID)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, nil, errs.New(err, errCode.Internal)
	}
	return user, auth, nil
}
