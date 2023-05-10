package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/security"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthService interface {
	Login(ctx context.Context, username string, password string) (*model.User, error)
}

type authService struct {
	pool     *pgxpool.Pool
	repo     repository.AuthRepository
	userRepo repository.UserRepository
}

func NewAuthService(pool *pgxpool.Pool, r repository.AuthRepository, userRepo repository.UserRepository) AuthService {
	return &authService{
		pool:     pool,
		repo:     r,
		userRepo: userRepo,
	}
}

func (s *authService) Login(ctx context.Context, username string, password string) (*model.User, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)
	user, err := s.userRepo.GetUserByPhone(ctx, tx, username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}
	auth, err := s.repo.TGetById(ctx, tx, user.ID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "auth not found: %v", err)
	}
	if !auth.Verified {
		return nil, status.Error(codes.PermissionDenied, "user not verified")
	}
	if auth.Disabled {
		return nil, status.Error(codes.PermissionDenied, "user disabled")
	}
	if auth.Locked {
		return nil, status.Error(codes.PermissionDenied, "user locked")
	}
	err = security.CheckPassword(auth.Password, password)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "invalid password")
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}
	return user, nil
}
