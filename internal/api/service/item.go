package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	category, err := i.categoryRepo.TGetByUUID(ctx, tx, categoryUUID)
	if err != nil {
		return nil, err
	}
	item := &model.Item{
		Name:     name,
		Category: *category,
	}
	err = i.repo.TCreate(ctx, tx, item)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return item, nil
}

func (i *item) Update(ctx context.Context, uuid string, name string, categoryUUID string) (*model.Item, error) {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	item, err := i.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, err
	}
	category, err := i.categoryRepo.TGetByUUID(ctx, tx, categoryUUID)
	if err != nil {
		return nil, err
	}

	item.Name = name
	item.Category = *category

	err = i.repo.TUpdate(ctx, tx, item)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return item, nil
}

func (i *item) Delete(ctx context.Context, uuid string) error {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	item, err := i.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return err
	}
	err = i.repo.TDelete(ctx, tx, item)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return nil
}

func (i *item) List(ctx context.Context) ([]*model.Item, error) {
	tx, err := i.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	items, err := i.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return items, nil
}
