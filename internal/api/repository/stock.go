package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Stock interface {
	Migrater
	TCRUDer[model.Stock]
}

type stock struct {
	db *pgxpool.Pool
}

func NewStock(db *pgxpool.Pool) Stock {
	return &stock{db: db}
}

func (r *stock) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *stock) TCreate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	//TODO implement me
	panic("implement me")
}

func (r *stock) TUpdate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	//TODO implement me
	panic("implement me")
}

func (r *stock) TDelete(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	//TODO implement me
	panic("implement me")
}

func (r *stock) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Stock, error) {
	//TODO implement me
	panic("implement me")
}

func (r *stock) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Stock, error) {
	//TODO implement me
	panic("implement me")
}

func (r *stock) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Stock, error) {
	//TODO implement me
	panic("implement me")
}
