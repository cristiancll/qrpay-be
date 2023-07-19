package service

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Item interface {
	Create(ctx context.Context, name string, categoryUUID string) (*model.Item, error)
	Update(ctx context.Context, uuid string, name string, categoryUUID string) (*model.Item, error)
	Delete(ctx context.Context, uuid string) error
	List(ctx context.Context) ([]*model.Item, error)
}

type item struct {
	pool         *pgxpool.Pool
	repo         repository.Item
	categoryRepo repository.Category
	opLogRepo    repository.OperationLog
}

func NewItem(pool *pgxpool.Pool, r repository.Item, categoryRepo repository.Category, opLogRepo repository.OperationLog) Item {
	return &item{
		pool:         pool,
		repo:         r,
		categoryRepo: categoryRepo,
		opLogRepo:    opLogRepo,
	}
}

func (i *item) Create(ctx context.Context, name string, categoryUUID string) (*model.Item, error) {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	category, err := i.categoryRepo.TGetByUUID(ctx, tx, categoryUUID)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetCategory, categoryUUID)
	}
	item := &model.Item{
		Name:     name,
		Category: *category,
	}
	err = i.repo.TCreate(ctx, tx, item)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedCreateItem, categoryUUID, name)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return item, nil
}

func (i *item) Update(ctx context.Context, uuid string, name string, categoryUUID string) (*model.Item, error) {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	item, err := i.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetItem, uuid)
	}
	category, err := i.categoryRepo.TGetByUUID(ctx, tx, categoryUUID)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetCategory, categoryUUID)
	}

	item.Name = name
	item.Category = *category

	err = i.repo.TUpdate(ctx, tx, item)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedUpdateItem, uuid, name)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return item, nil
}

func (i *item) Delete(ctx context.Context, uuid string) error {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	item, err := i.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return errs.Wrap(err, errMsg.FailedGetItem, uuid)
	}
	err = i.repo.TDelete(ctx, tx, item)
	if err != nil {
		return errs.Wrap(err, errMsg.FailedDeleteItem, uuid)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil
}

func (i *item) List(ctx context.Context) ([]*model.Item, error) {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	items, err := i.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetAllItem)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return items, nil
}
