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
	return Transaction[*model.Item](ctx, i.pool, func(tx pgx.Tx) (*model.Item, error) {
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

		return item, nil
	})
}

func (i *item) Update(ctx context.Context, uuid string, name string, categoryUUID string) (*model.Item, error) {
	return Transaction[*model.Item](ctx, i.pool, func(tx pgx.Tx) (*model.Item, error) {
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

		return item, nil
	})
}

func (i *item) Delete(ctx context.Context, uuid string) error {
	return TransactionVoid(ctx, i.pool, func(tx pgx.Tx) error {
		item, err := i.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetItem, uuid)
		}
		err = i.repo.TDelete(ctx, tx, item)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedDeleteItem, uuid)
		}
		return nil
	})
}

func (i *item) List(ctx context.Context) ([]*model.Item, error) {
	return TransactionReturnList[*model.Item](ctx, i.pool, func(tx pgx.Tx) ([]*model.Item, error) {
		items, err := i.repo.TGetAll(ctx, tx)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetAllItem)
		}
		return items, nil
	})
}
