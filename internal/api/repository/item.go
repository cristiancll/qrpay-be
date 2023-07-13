package repository

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Item interface {
	Migrater
	TCRUDer[model.Item]
}

type item struct {
	db *pgxpool.Pool
}

func NewItem(db *pgxpool.Pool) Item {
	return &item{db: db}
}

func (r *item) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS items(
				id SERIAL PRIMARY KEY,
				uuid VARCHAR(255) NOT NULL,
				category_id INT NOT NULL REFERENCES categories(id),
				name VARCHAR(255) NOT NULL UNIQUE,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}

func (r *item) TCreate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	item.UUID = uuid.New().String()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	query := "INSERT INTO items (uuid, category_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	id, err := tCreate(ctx, tx, query, item.UUID, item.Category.ID, item.Name, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		return errs.Wrap(err, "")
	}
	item.ID = id
	return nil
}

func (r *item) TUpdate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	item.UpdatedAt = time.Now().UTC()
	query := "UPDATE items SET category_id = $2, name = $3, updated_at = $4 WHERE id = $1"
	return tUpdate(ctx, tx, query, item.ID, item.Category.ID, item.Name, item.UpdatedAt)
}

func (r *item) TDelete(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	query := "DELETE FROM items WHERE id = $1"
	return tDelete(ctx, tx, query, item.ID)
}

func (r *item) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Item, error) {
	query := "SELECT id, uuid, category_id, name, created_at, updated_at FROM items WHERE id = $1"
	return tGet[model.Item](ctx, tx, query, id)
}

func (r *item) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Item, error) {
	query := "SELECT id, uuid, category_id, name, created_at, updated_at FROM items WHERE uuid = $1"
	return tGet[model.Item](ctx, tx, query, uuid)
}

func (r *item) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Item, error) {
	query := "SELECT i.id, i.uuid, c.uuid, c.name, c.created_at, c.updated_at, i.name, i.created_at, i.updated_at FROM items AS i JOIN categories AS c ON (i.category_id = c.id)"
	return tGetAll[model.Item](ctx, tx, query)
}
