package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockRepository interface {
	Migrater
	TCRUDer[model.Stock]
}

type stockRepository struct {
	db *pgxpool.Pool
}

func (r *stockRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *stockRepository) TCreate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	//TODO implement me
	panic("implement me")
}

func (r *stockRepository) TUpdate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	//TODO implement me
	panic("implement me")
}

func (r *stockRepository) TDelete(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	//TODO implement me
	panic("implement me")
}

func (r *stockRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Stock, error) {
	//TODO implement me
	panic("implement me")
}

func (r *stockRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Stock, error) {
	//TODO implement me
	panic("implement me")
}

func (r *stockRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Stock, error) {
	//TODO implement me
	panic("implement me")
}

func NewStockRepository(db *pgxpool.Pool) StockRepository {
	return &stockRepository{db: db}
}
