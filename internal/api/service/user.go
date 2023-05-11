package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/common"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"github.com/cristiancll/qrpay-be/internal/security"
	"github.com/cristiancll/qrpay-be/internal/wpp"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Create(ctx context.Context, user *model.User, password string) error
	Get(ctx context.Context, uuid string) (*model.User, error)
	List(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user *model.User, password string) error
	Delete(ctx context.Context, uuid string) error
	AdminCreated(ctx context.Context, user *model.User) error
}

type userService struct {
	pool     *pgxpool.Pool
	repo     repository.UserRepository
	authRepo repository.AuthRepository
	wpp      wpp.WhatsAppSystem
}

func NewUserService(pool *pgxpool.Pool, wpp wpp.WhatsAppSystem, r repository.UserRepository, authRepo repository.AuthRepository) UserService {
	return &userService{
		pool:     pool,
		repo:     r,
		authRepo: authRepo,
		wpp:      wpp,
	}
}

func (s *userService) Create(ctx context.Context, user *model.User, password string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	user.Phone = common.FormatPhone(user.Phone)

	err = s.repo.CountByPhone(ctx, tx, user.Phone)
	if err != nil {
		return status.Error(codes.AlreadyExists, errors.USER_ALREADY_EXISTS)
	}

	user.Role = roles.Client
	err = s.repo.TCreate(ctx, tx, user)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	passwordHash, err := security.HashPassword(password)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	auth := &model.Auth{
		UserID:   user.ID,
		Password: passwordHash,
		Verified: false,
	}
	err = s.authRepo.TCreate(ctx, tx, auth)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	go s.wpp.SendImage(user, user.WelcomeMessage())
	return nil
}

func (s *userService) Get(ctx context.Context, uuid string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) List(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *userService) Update(ctx context.Context, user *model.User, password string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	user.Phone = common.FormatPhone(user.Phone)

	existing, err := s.repo.TGetByUUID(ctx, tx, user.UUID)
	if err != nil {
		return err
	}
	if existing.UUID != user.UUID {
		return status.Error(codes.PermissionDenied, errors.UNAUTHORIZED)
	}

	// Check if password is correct
	auth, err := s.authRepo.TGetById(ctx, tx, existing.ID)
	if err != nil {
		return err
	}
	err = security.CheckPassword(auth.Password, password)
	if err != nil {
		return status.Error(codes.PermissionDenied, errors.INVALID_PASSWORD)
	}

	existing.Name = user.Name
	existing.Phone = user.Phone
	err = s.repo.TUpdate(ctx, tx, existing)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return nil
}

func (s *userService) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (s *userService) AdminCreated(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}
