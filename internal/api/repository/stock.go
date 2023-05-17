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
	TCountBySKU(ctx context.Context, tx pgx.Tx, skuID int64) error
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
	createStockQuery    = "INSERT INTO stocks (uuid, sku_id, quantity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	updateStockQuery    = "UPDATE stocks SET quantity = $2, updated_at = $3 WHERE id = $1"
	getStockByUUIDQuery = "SELECT id, uuid, sku_id, quantity, created_at, updated_at FROM stocks WHERE uuid = $1"
	getAllStocksQuery   = "SELECT s.id, s.uuid, sk.uuid, sk.name, sk.created_at, sk.updated_at, s.quantity, s.created_at, s.updated_at FROM stocks AS s JOIN skus AS sk ON (s.sku_id = sk.id)"
	countBySKUQuery     = "SELECT COUNT(*) FROM stocks WHERE sku_id = $1"
)

func (r *stock) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createStockTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *stock) TCreate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	stock.UUID = uuid.New().String()
	stock.CreatedAt = time.Now().UTC()
	stock.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createStockQuery, stock.UUID, stock.SKU.ID, stock.Quantity, stock.CreatedAt, stock.UpdatedAt).Scan(&stock.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *stock) TUpdate(ctx context.Context, tx pgx.Tx, stock *model.Stock) error {
	stock.UpdatedAt = time.Now().UTC()
	_, err := tx.Exec(ctx, updateStockQuery, stock.ID, stock.Quantity, stock.UpdatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *stock) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Stock, error) {
	stock := &model.Stock{}
	err := tx.QueryRow(ctx, getStockByUUIDQuery, uuid).Scan(&stock.ID, &stock.UUID, &stock.SKU.ID, &stock.Quantity, &stock.CreatedAt, &stock.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return stock, nil
}

func (r *stock) TCountBySKU(ctx context.Context, tx pgx.Tx, skuID int64) error {
	count := 0
	row := tx.QueryRow(ctx, countBySKUQuery, skuID)
	err := row.Scan(&count)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	if count > 0 {
		return status.Error(codes.Internal, errors.STOCK_ALREADY_EXISTS)
	}
	return nil
}

func (r *stock) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Stock, error) {
	stocks := make([]*model.Stock, 0)
	rows, err := tx.Query(ctx, getAllStocksQuery)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	for rows.Next() {
		stock := &model.Stock{}
		err := rows.Scan(&stock.ID, &stock.UUID, &stock.SKU.UUID, &stock.SKU.Name, &stock.SKU.CreatedAt, &stock.SKU.UpdatedAt, &stock.Quantity, &stock.CreatedAt, &stock.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}
