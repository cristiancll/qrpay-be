package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
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
	createSaleItemTableQuery = `CREATE TABLE IF NOT EXISTS sale_items(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		sale_id INT NOT NULL REFERENCES sales(id),
    		sku_id INT NOT NULL REFERENCES skus(id),
    		quantity INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	createSaleItemQuery = "INSERT INTO sale_items (uuid, sale_id, sku_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
)

func (r *saleItem) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createSaleItemTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *saleItem) TCreate(ctx context.Context, tx pgx.Tx, item *model.SaleItem) error {
	item.UUID = uuid.New().String()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createSaleItemQuery, item.UUID, item.Sale.ID, item.SKU.ID, item.Quantity, item.CreatedAt, item.UpdatedAt).Scan(&item.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
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
