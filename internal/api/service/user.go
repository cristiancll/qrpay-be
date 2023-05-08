package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/common"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService interface {
	Create(ctx context.Context, user *model.User, password string) error
	Get(ctx context.Context, uuid string) (*model.User, error)
	List(ctx context.Context) ([]model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, uuid string) error
	Login(ctx context.Context, user string, password string) (*model.User, error)
	Logout(ctx context.Context, token string) error
	AdminCreated(ctx context.Context, user *model.User) error
}

type userService struct {
	pool     *pgxpool.Pool
	repo     repository.UserRepository
	authRepo repository.AuthRepository
}

func NewUserService(pool *pgxpool.Pool, r repository.UserRepository, authRepo repository.AuthRepository) UserService {
	return &userService{
		pool:     pool,
		repo:     r,
		authRepo: authRepo,
	}
}

func (u *userService) Create(ctx context.Context, user *model.User, password string) error {
	tx, err := u.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = u.repo.CountByEmailOrPassword(ctx, tx, user.Email, user.Phone)
	if err != nil {
		return err
	}

	user.Role = roles.Client
	err = u.repo.Create(ctx, tx, user)
	if err != nil {
		return err
	}
	password, err = common.HashPassword(password)
	if err != nil {
		return err
	}
	auth := &model.Auth{
		UserID:   user.ID,
		Password: password,
		Verified: true, // TODO: send email to verify
	}
	err = u.authRepo.Create(ctx, tx, auth)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (u *userService) Get(ctx context.Context, uuid string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) List(ctx context.Context) ([]model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Update(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Login(ctx context.Context, user string, password string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userService) Logout(ctx context.Context, token string) error {
	//TODO implement me
	panic("implement me")
}

func (u *userService) AdminCreated(ctx context.Context, user *model.User) error {
	//TODO implement me
	panic("implement me")
}
