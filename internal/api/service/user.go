package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/common"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"github.com/cristiancll/qrpay-be/internal/security"
	"github.com/cristiancll/qrpay-be/internal/wpp"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type UserService interface {
	Create(ctx context.Context, user *model.User, password string) error
	Get(ctx context.Context, uuid string) (*model.User, error)
	List(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, uuid string) error
	AdminCreated(ctx context.Context, user *model.User) error
}

type userService struct {
	pool     *pgxpool.Pool
	repo     repository.UserRepository
	authRepo repository.AuthRepository
	wpp      wpp.WhatsAppClient
}

func NewUserService(pool *pgxpool.Pool, wpp wpp.WhatsAppClient, r repository.UserRepository, authRepo repository.AuthRepository) UserService {
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
		return status.Errorf(codes.Internal, "failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx)

	user.Email = strings.ToLower(user.Email)
	user.Phone = common.RemoveNonNumeric(user.Phone)

	err = s.repo.CountByEmailOrPassword(ctx, tx, user.Email, user.Phone)
	if err != nil {
		return status.Error(codes.AlreadyExists, err.Error())
	}

	user.Role = roles.Client
	err = s.repo.TCreate(ctx, tx, user)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	password, err = security.HashPassword(password)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}
	auth := &model.Auth{
		UserID:   user.ID,
		Password: password,
		Verified: true, // TODO: send email to verify
	}
	err = s.authRepo.TCreate(ctx, tx, auth)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to create auth: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		//return err
		return status.Errorf(codes.Internal, "failed to commit transaction: %v", err)
	}

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

func (s *userService) Update(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (s *userService) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (s *userService) AdminCreated(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}
