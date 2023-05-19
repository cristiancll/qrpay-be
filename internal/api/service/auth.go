package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/common"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/cristiancll/qrpay-be/internal/security"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)
	sanitizedPhone := common.SanitizePhone(phone)
	user, err := s.userRepo.GetUserByPhone(ctx, tx, sanitizedPhone)
	if err != nil {
		return nil, nil, status.Error(codes.PermissionDenied, errors.INVALID_CREDENTIALS)
	}
	auth, err := s.repo.TGetById(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, status.Error(codes.PermissionDenied, errors.INVALID_CREDENTIALS)
	}
	if !auth.Verified {
		return nil, nil, status.Error(codes.PermissionDenied, errors.VERIFICATION_ERROR)
	}
	if auth.Disabled {
		return nil, nil, status.Error(codes.PermissionDenied, errors.DISABLED_USER)
	}
	err = security.CheckPassword(auth.Password, password)
	if err != nil {
		return nil, nil, status.Error(codes.PermissionDenied, errors.INVALID_CREDENTIALS)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return user, auth, nil
}

func (s *auth) Heartbeat(ctx context.Context) (*model.User, *model.Auth, error) {
	UUID := ctx.Value("UUID").(string)
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)
	user, err := s.userRepo.TGetByUUID(ctx, tx, UUID)
	if err != nil {
		return nil, nil, err
	}
	auth, err := s.repo.TGetById(ctx, tx, user.ID)
	if err != nil {
		return nil, nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return user, auth, nil
}
