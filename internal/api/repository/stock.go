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

type Stock interface {
	Migrater
	TCreater[model.Stock]
	TUpdater[model.Stock]
	TGetterAll[model.Stock]
	TGetterByUUID[model.Stock]
	TDeleter[model.Stock]
	TCountBySKU(ctx context.Context, tx pgx.Tx, skuID int64) (int64, error)
	TDecreaseStock(ctx context.Context, tx pgx.Tx, sku *model.SKU, quantity int64) error
}

type stock struct {
	db *pgxpool.Pool
}

func NewStock(db *pgxpool.Pool) Stock {
	return &stock{db: db}
}

func (r *stock) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS stocks(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		sku_id INT NOT NULL REFERENCES skus(id),
    		quantity INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}

func (r *stock) TCreate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	stock.UUID = uuid.New().String()
	stock.CreatedAt = time.Now().UTC()
	stock.UpdatedAt = time.Now().UTC()
	query := "INSERT INTO stocks (uuid, sku_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	id, err := tCreate(ctx, tx, query, stock.UUID, stock.SKU.ID, stock.Quantity, stock.CreatedAt, stock.UpdatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	stock.ID = id
	return nil
}

func (r *stock) TUpdate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	stock.UpdatedAt = time.Now().UTC()
	query := "UPDATE stocks SET quantity = $2, updated_at = $3 WHERE id = $1"
	return tUpdate(ctx, tx, query, stock.ID, stock.Quantity, stock.UpdatedAt)
}

func (r *stock) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Stock, error) {
	query := "SELECT id, uuid, sku_id, quantity, created_at, updated_at FROM stocks WHERE uuid = $1"
	return tGet[model.Stock](ctx, tx, query, uuid)
}

func (r *stock) TCountBySKU(ctx context.Context, tx pgx.Tx, skuID int64) (int64, error) {
	query := "SELECT COUNT(*) FROM stocks WHERE sku_id = $1"
	return tCount(ctx, tx, query, skuID)
}

func (r *stock) TDelete(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	query := "DELETE FROM stocks WHERE id = $1"
	return tDelete(ctx, tx, query, stock.ID)
}

func (r *stock) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Stock, error) {
	query := "SELECT s.id, s.uuid, sk.uuid, sk.name, sk.created_at, sk.updated_at, s.quantity, s.created_at, s.updated_at FROM stocks AS s JOIN skus AS sk ON (s.sku_id = sk.id)"
	return tGetAll[model.Stock](ctx, tx, query)
}

func (r *stock) TDecreaseStock(ctx context.Context, tx pgx.Tx, sku *model.SKU, quantity int64) error {
	query := "UPDATE stocks SET quantity = quantity - $2, updated_at = $3 WHERE sku_id = $1"
	return tUpdate(ctx, tx, query, sku.ID, quantity, time.Now().UTC())
}
