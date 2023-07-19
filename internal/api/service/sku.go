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

type SKU interface {
	Create(ctx context.Context, itemUUID string, name string, description string, price int64) (*model.SKU, error)
	Update(ctx context.Context, uuid string, itemUUID string, name string, description string, price int64) (*model.SKU, error)
	Delete(ctx context.Context, uuid string) error
	List(ctx context.Context) ([]*model.SKU, error)
}

type sku struct {
	pool      *pgxpool.Pool
	repo      repository.SKU
	itemRepo  repository.Item
	opLogRepo repository.OperationLog
}

func NewSKU(pool *pgxpool.Pool, repo repository.SKU, itemRepo repository.Item, opLogRepo repository.OperationLog) SKU {
	return &sku{
		pool:      pool,
		repo:      repo,
		itemRepo:  itemRepo,
		opLogRepo: opLogRepo,
	}
}

func (s *sku) Create(ctx context.Context, itemUUID string, name string, description string, price int64) (*model.SKU, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	item, err := s.itemRepo.TGetByUUID(ctx, tx, itemUUID)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetItem, itemUUID)
	}
	sku := &model.SKU{
		Name:        name,
		Description: description,
		Price:       price,
		Item:        *item,
	}
	err = s.repo.TCreate(ctx, tx, sku)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedCreateSKU, sku)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return sku, nil
}

func (s *sku) Update(ctx context.Context, uuid string, itemUUID string, name string, description string, price int64) (*model.SKU, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	sku, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetSKU, uuid)
	}
	item, err := s.itemRepo.TGetByUUID(ctx, tx, itemUUID)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetItem, itemUUID)
	}
	sku.Name = name
	sku.Description = description
	sku.Price = price
	sku.Item = *item

	err = s.repo.TUpdate(ctx, tx, sku)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedUpdateSKU, sku)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return sku, nil
}

func (s *sku) Delete(ctx context.Context, uuid string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	sku, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return errs.Wrap(err, errMsg.FailedGetSKU, uuid)
	}
	err = s.repo.TDelete(ctx, tx, sku)
	if err != nil {
		return errs.Wrap(err, errMsg.FailedDeleteSKU, sku)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil
}

func (s *sku) List(ctx context.Context) ([]*model.SKU, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	skus, err := s.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetAllSKU)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, errs.New(err, errCode.Internal)
	}
	return skus, nil
}
