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

type Stock interface {
	Create(ctx context.Context, skuUUID string, quantity int64) (*model.Stock, error)
	Update(ctx context.Context, uuid string, quantity int64) (*model.Stock, error)
	List(ctx context.Context) ([]*model.Stock, error)
}

type stock struct {
	pool    *pgxpool.Pool
	repo    repository.Stock
	skuRepo repository.SKU
}

func NewStock(pool *pgxpool.Pool, r repository.Stock, skuRepo repository.SKU) Stock {
	return &stock{
		pool:    pool,
		repo:    r,
		skuRepo: skuRepo,
	}
}

func (s *stock) Create(ctx context.Context, skuUUID string, quantity int64) (*model.Stock, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	sku, err := s.skuRepo.TGetByUUID(ctx, tx, skuUUID)
	if err != nil {
		return nil, err
	}

	err = s.repo.TCountBySKU(ctx, tx, sku.ID)
	if err != nil {
		return nil, err
	}

	stock := &model.Stock{
		SKU:      *sku,
		Quantity: quantity,
	}
	err = s.repo.TCreate(ctx, tx, stock)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return stock, nil
}

func (s *stock) Update(ctx context.Context, uuid string, quantity int64) (*model.Stock, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	stock, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, err
	}

	sku, err := s.skuRepo.TGetById(ctx, tx, stock.SKU.ID)
	if err != nil {
		return nil, err
	}
	stock.Quantity = quantity
	stock.SKU = *sku

	err = s.repo.TUpdate(ctx, tx, stock)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return stock, nil
}

func (s *stock) List(ctx context.Context) ([]*model.Stock, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	stocks, err := s.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return stocks, nil
}
