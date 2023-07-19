package service

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Category interface {
	Create(ctx context.Context, name string) (*model.Category, error)
	Update(ctx context.Context, uuid string, name string) (*model.Category, error)
	Delete(ctx context.Context, uuid string) error
	List(ctx context.Context) ([]*model.Category, error)
}

type category struct {
	pool      *pgxpool.Pool
	repo      repository.Category
	opLogRepo repository.OperationLog
}

func NewCategory(pool *pgxpool.Pool, r repository.Category, opLogRepo repository.OperationLog) Category {
	return &category{
		pool:      pool,
		repo:      r,
		opLogRepo: opLogRepo,
	}
}

func (c category) Create(ctx context.Context, name string) (*model.Category, error) {
	return Transaction[*model.Category](ctx, c.pool, func(tx pgx.Tx) (*model.Category, error) {
		category := &model.Category{
			Name: name,
		}
		err := c.repo.TCreate(ctx, tx, category)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedCreateCategory, name)
		}
		return category, nil
	})
}

func (c category) Update(ctx context.Context, uuid string, name string) (*model.Category, error) {
	return Transaction[*model.Category](ctx, c.pool, func(tx pgx.Tx) (*model.Category, error) {
		existing, err := c.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetCategory, uuid)
		}
		existing.Name = name
		err = c.repo.TUpdate(ctx, tx, existing)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedUpdateCategory, uuid, name)
		}
		return existing, nil
	})
}

func (c category) Delete(ctx context.Context, uuid string) error {
	return TransactionVoid(ctx, c.pool, func(tx pgx.Tx) error {
		existing, err := c.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetCategory, uuid)
		}
		err = c.repo.TDelete(ctx, tx, existing)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedDeleteCategory, uuid)
		}
		return nil
	})
}

func (c category) List(ctx context.Context) ([]*model.Category, error) {
	return TransactionReturnList[*model.Category](ctx, c.pool, func(tx pgx.Tx) ([]*model.Category, error) {
		categories, err := c.repo.TGetAll(ctx, tx)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetAllCategory)
		}
		return categories, nil
	})
}
