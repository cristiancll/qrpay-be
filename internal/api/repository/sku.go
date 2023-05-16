package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

const (
	createSKUTableQuery = `CREATE TABLE IF NOT EXISTS skus (
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		item_id INT NOT NULL REFERENCES items(id),
    		name VARCHAR(255) NOT NULL UNIQUE,
    		description VARCHAR(255),
    		price INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
)

func (s sku) Migrate(ctx context.Context) error {
	_, err := s.db.Exec(ctx, createSKUTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return err
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
