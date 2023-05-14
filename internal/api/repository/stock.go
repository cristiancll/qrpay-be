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

const (
	createStockTableQuery = `CREATE TABLE IF NOT EXISTS stocks(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		sku_id INT NOT NULL REFERENCES skus(id),
    		quantity INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
)

func (r *stock) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createStockTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
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
