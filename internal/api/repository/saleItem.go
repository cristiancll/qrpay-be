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

type SaleItem interface {
	Migrater
	TCRUDer[model.SaleItem]
}

type saleItem struct {
	db *pgxpool.Pool
}

func NewSaleItem(db *pgxpool.Pool) SaleItem {
	return &saleItem{db: db}
}

const (
	createSaleItemTableQuery = `CREATE TABLE IF NOT EXISTS sale_item(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		sale_id INT NOT NULL REFERENCES "sale" (id),
    		sku_id INT NOT NULL REFERENCES "sku" (id),
    		quantity INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
)

func (r *saleItem) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createSaleItemTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *saleItem) TCreate(ctx context.Context, tx pgx.Tx, item *model.SaleItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *saleItem) TUpdate(ctx context.Context, tx pgx.Tx, item *model.SaleItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *saleItem) TDelete(ctx context.Context, tx pgx.Tx, item *model.SaleItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *saleItem) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.SaleItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *saleItem) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.SaleItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *saleItem) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.SaleItem, error) {
	//TODO implement me
	panic("implement me")
}
