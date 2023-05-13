package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SKU interface {
	Migrater
	TCRUDer[model.SKU]
}

type sku struct {
	db *pgxpool.Pool
}

func NewSKU(db *pgxpool.Pool) SKU {
	return &sku{db: db}
}

func (s sku) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s sku) TCreate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	//TODO implement me
	panic("implement me")
}

func (s sku) TUpdate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	//TODO implement me
	panic("implement me")
}

func (s sku) TDelete(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	//TODO implement me
	panic("implement me")
}

func (s sku) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.SKU, error) {
	//TODO implement me
	panic("implement me")
}

func (s sku) TGetByUUID(ctx context.Context, tx pgx.Tx, s2 string) (*model.SKU, error) {
	//TODO implement me
	panic("implement me")
}

func (s sku) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.SKU, error) {
	//TODO implement me
	panic("implement me")
}
