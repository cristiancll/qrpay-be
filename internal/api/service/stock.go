package service

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Stock interface {
	Create(ctx context.Context, skuUUID string, quantity int64) (*model.Stock, error)
	Update(ctx context.Context, uuid string, quantity int64) (*model.Stock, error)
	Delete(ctx context.Context, uuid string) error
	List(ctx context.Context) ([]*model.Stock, error)
}

type stock struct {
	pool      *pgxpool.Pool
	repo      repository.Stock
	skuRepo   repository.SKU
	opLogRepo repository.OperationLog
}

func NewStock(pool *pgxpool.Pool, r repository.Stock, skuRepo repository.SKU, opLogRepo repository.OperationLog) Stock {
	return &stock{
		pool:      pool,
		repo:      r,
		skuRepo:   skuRepo,
		opLogRepo: opLogRepo,
	}
}

func (s *stock) Create(ctx context.Context, skuUUID string, quantity int64) (*model.Stock, error) {
	return Transaction[*model.Stock](ctx, s.pool, func(tx pgx.Tx) (*model.Stock, error) {
		sku, err := s.skuRepo.TGetByUUID(ctx, tx, skuUUID)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetSKU, skuUUID)
		}

		count, err := s.repo.TCountBySKU(ctx, tx, sku.ID)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedCountStock, sku.ID)
		}
		if count > 0 {
			return nil, errs.New(errors.New(errMsg.StockAlreadyExists), errCode.AlreadyExists, skuUUID, count)
		}

		stock := &model.Stock{
			SKU:      *sku,
			Quantity: quantity,
		}
		err = s.repo.TCreate(ctx, tx, stock)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedCreateStock, stock)
		}

		return stock, nil
	})
}

func (s *stock) Update(ctx context.Context, uuid string, quantity int64) (*model.Stock, error) {
	return Transaction[*model.Stock](ctx, s.pool, func(tx pgx.Tx) (*model.Stock, error) {
		stock, err := s.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetStock, uuid)
		}

		sku, err := s.skuRepo.TGetById(ctx, tx, stock.SKU.ID)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetSKU, stock.SKU.ID)
		}
		stock.Quantity = quantity
		stock.SKU = *sku

		err = s.repo.TUpdate(ctx, tx, stock)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedUpdateStock, stock)
		}

		return stock, nil
	})
}

func (s *stock) Delete(ctx context.Context, uuid string) error {
	return TransactionVoid(ctx, s.pool, func(tx pgx.Tx) error {
		stock, err := s.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetStock, uuid)
		}
		err = s.repo.TDelete(ctx, tx, stock)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedDeleteStock, stock)
		}
		return nil
	})
}

func (s *stock) List(ctx context.Context) ([]*model.Stock, error) {
	return TransactionReturnList[*model.Stock](ctx, s.pool, func(tx pgx.Tx) ([]*model.Stock, error) {
		stocks, err := s.repo.TGetAll(ctx, tx)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetAllStock)
		}
		return stocks, nil
	})
}
