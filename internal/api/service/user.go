package service

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/common"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"github.com/cristiancll/qrpay-be/internal/security"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User interface {
	Create(ctx context.Context, name string, phone string, password string) (*model.User, error)
	Get(ctx context.Context, uuid string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, user *model.User, password string) error
	Delete(ctx context.Context, uuid string) error
	AdminCreated(ctx context.Context, name string, phone string, sellerUUID string) (*model.User, error)
	UpdateRole(ctx context.Context, uuid string, role roles.Role, enabled bool) (*model.User, error)
}

type user struct {
	pool      *pgxpool.Pool
	repo      repository.User
	authRepo  repository.Auth
	opLogRepo repository.OperationLog
}

func NewUser(pool *pgxpool.Pool, r repository.User, authRepo repository.Auth, opLogRepo repository.OperationLog) User {
	return &user{
		pool:      pool,
		repo:      r,
		authRepo:  authRepo,
		opLogRepo: opLogRepo,
	}
}

func (s *user) Create(ctx context.Context, name string, phone string, password string) (*model.User, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	phone = common.SanitizePhone(phone)

	count, err := s.repo.CountByPhone(ctx, tx, phone)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if count > 0 {
		return nil, errs.New(errors.New(""), errCode.AlreadyExists)
	}

	user := &model.User{
		Name:  name,
		Phone: phone,
		Role:  roles.Client,
	}
	err = s.repo.TCreate(ctx, tx, user)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	passwordHash, err := security.HashPassword(password)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	auth := &model.Auth{
		UserID:   user.ID,
		Password: passwordHash,
	}
	err = s.authRepo.TCreate(ctx, tx, auth)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	//go s.wpp.SendImage(user, user.WelcomeMessage()) // TODO:
	return user, nil
}

func (s *user) Get(ctx context.Context, uuid string) (*model.User, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	user, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return user, nil
}

func (s *user) List(ctx context.Context) ([]*model.User, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	users, err := s.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return users, nil
}

func (s *user) Update(ctx context.Context, user *model.User, password string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	user.Phone = common.SanitizePhone(user.Phone)

	existing, err := s.repo.TGetByUUID(ctx, tx, user.UUID)
	if err != nil {
		return errs.Wrap(err, "")
	}
	if existing.UUID != user.UUID {
		return errs.New(errors.New(""), errCode.AccessDenied)
	}

	// Check if password is correct
	auth, err := s.authRepo.TGetById(ctx, tx, existing.ID)
	if err != nil {
		return errs.Wrap(err, "")
	}
	err = security.CheckPassword(auth.Password, password)
	if err != nil {
		return errs.Wrap(err, "")
	}

	existing.Name = user.Name
	existing.Phone = user.Phone
	err = s.repo.TUpdate(ctx, tx, existing)
	if err != nil {
		return errs.Wrap(err, "")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil
}

func (s *user) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (s *user) AdminCreated(ctx context.Context, name string, phone string, sellerUUID string) (*model.User, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	seller, err := s.repo.TGetByUUID(ctx, tx, sellerUUID)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	phone = common.SanitizePhone(phone)

	count, err := s.repo.CountByPhone(ctx, tx, phone)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if count > 0 {
		return nil, errs.New(errors.New(""), errCode.AlreadyExists)
	}

	user := &model.User{
		Name:  name,
		Phone: phone,
		Role:  roles.Client,
	}
	err = s.repo.TCreate(ctx, tx, user)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	// This password will be changed by the user upon first login
	passwordHash, err := security.HashPassword(security.RandomPassword())
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	auth := &model.Auth{
		UserID:   user.ID,
		Password: passwordHash,
	}
	err = s.authRepo.TCreate(ctx, tx, auth)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}

	opLog := &model.OperationLog{
		User:        *user,
		Seller:      *seller,
		Operation:   "User",
		OperationId: user.ID,
	}
	_ = s.opLogRepo.Create(context.Background(), opLog)

	//go s.wpp.SendImage(user, user.WelcomeMessage()) // TODO:
	return user, nil
}

func (s *user) UpdateRole(ctx context.Context, uuid string, role roles.Role, enabled bool) (*model.User, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	user, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	user.Role = user.Role.ToggleRole(role, enabled)

	err = s.repo.TUpdate(ctx, tx, user)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return user, nil
}
