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

type SKU interface {
	Create(ctx context.Context, itemUUID string, name string, description string, price int64) (*model.SKU, error)
	Update(ctx context.Context, uuid string, itemUUID string, name string, description string, price int64) (*model.SKU, error)
	Delete(ctx context.Context, uuid string) error
	List(ctx context.Context) ([]*model.SKU, error)
}

type sku struct {
	pool     *pgxpool.Pool
	repo     repository.SKU
	itemRepo repository.Item
}

func NewSKU(pool *pgxpool.Pool, repo repository.SKU, itemRepo repository.Item) SKU {
	return &sku{
		pool:     pool,
		repo:     repo,
		itemRepo: itemRepo,
	}
}

func (s *sku) Create(ctx context.Context, itemUUID string, name string, description string, price int64) (*model.SKU, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	item, err := s.itemRepo.TGetByUUID(ctx, tx, itemUUID)
	if err != nil {
		return nil, err
	}
	sku := &model.SKU{
		Name:        name,
		Description: description,
		Price:       price,
		Item:        *item,
	}
	err = s.repo.TCreate(ctx, tx, sku)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return sku, nil
}

func (s *sku) Update(ctx context.Context, uuid string, itemUUID string, name string, description string, price int64) (*model.SKU, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	sku, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, err
	}
	item, err := s.itemRepo.TGetByUUID(ctx, tx, itemUUID)
	if err != nil {
		return nil, err
	}
	sku.Name = name
	sku.Description = description
	sku.Price = price
	sku.Item = *item

	err = s.repo.TUpdate(ctx, tx, sku)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return sku, nil
}

func (s *sku) Delete(ctx context.Context, uuid string) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	sku, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return err
	}
	err = s.repo.TDelete(ctx, tx, sku)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return nil
}

func (s *sku) List(ctx context.Context) ([]*model.SKU, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	skus, err := s.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return skus, nil
}
