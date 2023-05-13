package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SKURepository interface {
	Migrater
	TCRUDer[model.SKU]
}

type skuRepository struct {
	db *pgxpool.Pool
}

func (s skuRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s skuRepository) TCreate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	//TODO implement me
	panic("implement me")
}

func (s skuRepository) TUpdate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	//TODO implement me
	panic("implement me")
}

func (s skuRepository) TDelete(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	//TODO implement me
	panic("implement me")
}

func (s skuRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.SKU, error) {
	//TODO implement me
	panic("implement me")
}

func (s skuRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s2 string) (*model.SKU, error) {
	//TODO implement me
	panic("implement me")
}

func (s skuRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.SKU, error) {
	//TODO implement me
	panic("implement me")
}

func NewSKURepository(db *pgxpool.Pool) SKURepository {
	return &skuRepository{db: db}
}
