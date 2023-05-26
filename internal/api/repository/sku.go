package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type SKU interface {
	Migrater
	TCRUDer[model.SKU]
	TGetAllByUUIDs(ctx context.Context, tx pgx.Tx, uuids []string) ([]*model.SKU, error)
}

type sku struct {
	db *pgxpool.Pool
}

func NewSKU(db *pgxpool.Pool) SKU {
	return &sku{db: db}
}

func (s sku) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS skus (
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		item_id INT NOT NULL REFERENCES items(id),
    		name VARCHAR(255) NOT NULL UNIQUE,
    		description VARCHAR(255),
    		price INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, s.db, query)
}

func (s *sku) TCreate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	sku.UUID = uuid.New().String()
	sku.CreatedAt = time.Now().UTC()
	sku.UpdatedAt = time.Now().UTC()
	query := "INSERT INTO skus (uuid, item_id, name, description, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	id, err := tCreate(ctx, tx, query, sku.UUID, sku.Item.ID, sku.Name, sku.Description, sku.Price, sku.CreatedAt, sku.UpdatedAt)
	if err != nil {
		return err
	}
	sku.ID = id
	return nil
}

func (s *sku) TUpdate(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	sku.UpdatedAt = time.Now().UTC()
	query := "UPDATE skus SET item_id = $2, name = $3, description = $4, price = $5, updated_at = $6 WHERE id = $1"
	return tUpdate(ctx, tx, query, sku.ID, sku.Item.ID, sku.Name, sku.Description, sku.Price, sku.UpdatedAt)
}

func (s *sku) TDelete(ctx context.Context, tx pgx.Tx, sku *model.SKU) error {
	query := "DELETE FROM skus WHERE id = $1"
	return tDelete(ctx, tx, query, sku.ID)
}

func (s *sku) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.SKU, error) {
	query := "SELECT id, uuid, item_id, name, description, price, created_at, updated_at FROM skus WHERE id = $1"
	return tGet[model.SKU](ctx, tx, query, id)
}

func (s *sku) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.SKU, error) {
	query := "SELECT id, uuid, item_id, name, description, price, created_at, updated_at FROM skus WHERE uuid = $1"
	return tGet[model.SKU](ctx, tx, query, uuid)
}

func (s *sku) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.SKU, error) {
	query := "SELECT s.id, s.uuid, i.id, i.uuid, i.category_id, i.name, i.created_at, i.updated_at, s.name, s.description, s.price, s.created_at, s.updated_at FROM skus AS s JOIN items AS i ON s.item_id = i.id"
	return tGetAll[model.SKU](ctx, tx, query)
}

func (s *sku) TGetAllByUUIDs(ctx context.Context, tx pgx.Tx, uuids []string) ([]*model.SKU, error) {
	query := "SELECT id, uuid, item_id, name, description, price, created_at, updated_at FROM skus WHERE uuid = ANY($1)"
	return tGetAll[model.SKU](ctx, tx, query, uuids)
}
