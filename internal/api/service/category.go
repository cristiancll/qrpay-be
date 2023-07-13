package service

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errCode"
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
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	category := &model.Category{
		Name: name,
	}
	err = c.repo.TCreate(ctx, tx, category)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return category, nil
}

func (c category) Update(ctx context.Context, uuid string, name string) (*model.Category, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	existing, err := c.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	existing.Name = name
	err = c.repo.TUpdate(ctx, tx, existing)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return existing, nil
}

func (c category) Delete(ctx context.Context, uuid string) error {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)
	existing, err := c.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return errs.Wrap(err, "")
	}
	err = c.repo.TDelete(ctx, tx, existing)
	if err != nil {
		return errs.Wrap(err, "")
	}
	err = tx.Commit(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil

}

func (c category) List(ctx context.Context) ([]*model.Category, error) {
	tx, err := c.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	categories, err := c.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return categories, nil
}
