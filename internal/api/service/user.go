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
	"github.com/cristiancll/qrpay-be/internal/roles"
	"github.com/cristiancll/qrpay-be/internal/security"
	"github.com/jackc/pgx/v5"
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
	return Transaction[*model.User](ctx, s.pool, func(tx pgx.Tx) (*model.User, error) {
		phone = common.SanitizePhone(phone)

		count, err := s.repo.CountByPhone(ctx, tx, phone)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedCountUser)
		}
		if count > 0 {
			return nil, errs.New(errors.New(errMsg.UserAlreadyExists), errCode.AlreadyExists)
		}

		user := &model.User{
			Name:  name,
			Phone: phone,
			Role:  roles.Client,
		}
		err = s.repo.TCreate(ctx, tx, user)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedCreateUser, user)
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
			return nil, errs.Wrap(err, errMsg.FailedCreateAuth, user.ID)
		}

		//go s.wpp.SendImage(user, user.WelcomeMessage()) // TODO:
		return user, nil
	})
}

func (s *user) Get(ctx context.Context, uuid string) (*model.User, error) {
	return Transaction[*model.User](ctx, s.pool, func(tx pgx.Tx) (*model.User, error) {
		user, err := s.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetUser, uuid)
		}
		return user, nil
	})
}

func (s *user) List(ctx context.Context) ([]*model.User, error) {
	return TransactionReturnList[*model.User](ctx, s.pool, func(tx pgx.Tx) ([]*model.User, error) {
		users, err := s.repo.TGetAll(ctx, tx)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetAllUser)
		}
		return users, nil
	})
}

func (s *user) Update(ctx context.Context, user *model.User, password string) error {
	return TransactionVoid(ctx, s.pool, func(tx pgx.Tx) error {
		user.Phone = common.SanitizePhone(user.Phone)

		existing, err := s.repo.TGetByUUID(ctx, tx, user.UUID)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetUser, user.UUID)
		}

		// Check if password is correct
		auth, err := s.authRepo.TGetById(ctx, tx, existing.ID)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetAuth, existing.ID)
		}
		err = security.CheckPassword(auth.Password, password)
		if err != nil {
			return errs.Wrap(err, errMsg.IncorrectPassword)
		}

		existing.Name = user.Name
		existing.Phone = user.Phone
		err = s.repo.TUpdate(ctx, tx, existing)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedUpdateUser, existing)
		}

		return nil
	})
}

func (s *user) Delete(ctx context.Context, uuid string) error {
	//TODO implement me
	panic("implement me")
}

func (s *user) AdminCreated(ctx context.Context, name string, phone string, sellerUUID string) (*model.User, error) {
	return Transaction[*model.User](ctx, s.pool, func(tx pgx.Tx) (*model.User, error) {
		seller, err := s.repo.TGetByUUID(ctx, tx, sellerUUID)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetUser, sellerUUID)
		}

		phone = common.SanitizePhone(phone)

		count, err := s.repo.CountByPhone(ctx, tx, phone)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedCountUser)
		}
		if count > 0 {
			return nil, errs.New(errors.New(errMsg.UserAlreadyExists), errCode.AlreadyExists)
		}

		user := &model.User{
			Name:  name,
			Phone: phone,
			Role:  roles.Client,
		}
		err = s.repo.TCreate(ctx, tx, user)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedCreateUser, user)
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
			return nil, errs.Wrap(err, errMsg.FailedCreateAuth, user.ID)
		}

		opLog := &model.OperationLog{
			User:        *user,
			Seller:      *seller,
			Operation:   "User",
			OperationId: user.ID,
		}
		_ = s.opLogRepo.Create(context.Background(), opLog) // We ignore the error here since it's not critical

		//go s.wpp.SendImage(user, user.WelcomeMessage()) // TODO:
		return user, nil
	})
}

func (s *user) UpdateRole(ctx context.Context, uuid string, role roles.Role, enabled bool) (*model.User, error) {
	return Transaction[*model.User](ctx, s.pool, func(tx pgx.Tx) (*model.User, error) {
		user, err := s.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetUser, uuid)
		}
		user.Role = user.Role.ToggleRole(role, enabled)

		err = s.repo.TUpdate(ctx, tx, user)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedUpdateUser, user)
		}

		return user, nil
	})
}
