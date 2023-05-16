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
	createSKUQuery    = "INSERT INTO skus (uuid, item_id, name, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	updateSKUQuery    = "UPDATE skus SET item_id = $2, name = $3, description = $4, price = $5, updated_at = $6 WHERE id = $1"
	deleteSKUQuery    = "DELETE FROM skus WHERE id = $1"
	getSKUByIDQuery   = "SELECT id, uuid, item_id, name, description, price, created_at, updated_at FROM skus WHERE id = $1"
	getSKUByUUIDQuery = "SELECT id, uuid, item_id, name, description, price, created_at, updated_at FROM skus WHERE uuid = $1"
	getAllSKUsQuery   = "SELECT s.id, s.uuid, i.id, i.uuid, i.category_id, i.name, i.created_at, i.updated_at, s.name, s.description, s.price, s.created_at, s.updated_at FROM skus AS s JOIN items AS i ON s.item_id = i.id"
)

func (s sku) Migrate(ctx context.Context) error {
	_, err := s.db.Exec(ctx, createSKUTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return err
}

func (s sku) TCreate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	sku.UUID = uuid.New().String()
	sku.CreatedAt = time.Now().UTC()
	sku.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createSKUQuery, sku.UUID, sku.Item.ID, sku.Name, sku.Description, sku.Price, sku.CreatedAt, sku.UpdatedAt).Scan(&sku.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (s sku) TUpdate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	sku.UpdatedAt = time.Now().UTC()
	_, err := tx.Exec(ctx, updateSKUQuery, sku.ID, sku.Item.ID, sku.Name, sku.Description, sku.Price, sku.UpdatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (s sku) TDelete(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	_, err := tx.Exec(ctx, deleteSKUQuery, sku.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (s sku) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.SKU, error) {
	sku := &model.SKU{}
	err := tx.QueryRow(ctx, getSKUByIDQuery, id).Scan(&sku.ID, &sku.UUID, &sku.Item.ID, &sku.Name, &sku.Description, &sku.Price, &sku.CreatedAt, &sku.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return sku, nil

}

func (s sku) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.SKU, error) {
	sku := &model.SKU{}
	err := tx.QueryRow(ctx, getSKUByUUIDQuery, uuid).Scan(&sku.ID, &sku.UUID, &sku.Item.ID, &sku.Name, &sku.Description, &sku.Price, &sku.CreatedAt, &sku.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return sku, nil
}

func (s sku) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.SKU, error) {
	skus := make([]*model.SKU, 0)
	rows, err := tx.Query(ctx, getAllSKUsQuery)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	for rows.Next() {
		sku := &model.SKU{}
		err := rows.Scan(&sku.ID, &sku.UUID, &sku.Item.ID, &sku.Item.UUID, &sku.Item.Category.ID, &sku.Item.Name, &sku.Item.CreatedAt, &sku.Item.UpdatedAt, &sku.Name, &sku.Description, &sku.Price, &sku.CreatedAt, &sku.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		skus = append(skus, sku)
	}
	return skus, nil
}
