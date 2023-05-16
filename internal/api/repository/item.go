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

const (
	createItemTableQuery = `CREATE TABLE IF NOT EXISTS items(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		category_id INT NOT NULL REFERENCES categories(id),
    		name VARCHAR(255) NOT NULL UNIQUE,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	createItemQuery    = "INSERT INTO items (uuid, category_id, name, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	updateItemQuery    = "UPDATE items SET category_id = $2, name = $3, updated_at = $4 WHERE id = $1"
	deleteItemQuery    = "DELETE FROM items WHERE id = $1"
	getItemByIDQuery   = "SELECT id, uuid, category_id, name, created_at, updated_at FROM items WHERE id = $1"
	getItemByUUIDQuery = "SELECT id, uuid, category_id, name, created_at, updated_at FROM items WHERE uuid = $1"
	getAllItemsQuery   = "SELECT i.id, i.uuid, c.uuid, c.name, c.created_at, c.updated_at, i.name, i.created_at, i.updated_at FROM items AS i JOIN categories AS c ON (i.category_id = c.id)"
)

func (r *item) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createItemTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *item) TCreate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	item.UUID = uuid.New().String()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createItemQuery, item.UUID, item.Category.ID, item.Name, item.CreatedAt, item.UpdatedAt).Scan(&item.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *item) TUpdate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	item.UpdatedAt = time.Now().UTC()
	_, err := tx.Exec(ctx, updateItemQuery, item.ID, item.Category.ID, item.Name, item.UpdatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *item) TDelete(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	_, err := tx.Exec(ctx, deleteItemQuery, item.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *item) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Item, error) {
	item := &model.Item{}
	err := tx.QueryRow(ctx, getItemByIDQuery, id).Scan(&item.ID, &item.UUID, &item.Category.ID, &item.Name, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return item, nil
}

func (r *item) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Item, error) {
	item := &model.Item{}
	err := tx.QueryRow(ctx, getItemByUUIDQuery, uuid).Scan(&item.ID, &item.UUID, &item.Category.ID, &item.Name, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return item, nil
}

func (r *item) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Item, error) {
	items := []*model.Item{}
	rows, err := tx.Query(ctx, getAllItemsQuery)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	for rows.Next() {
		item := &model.Item{}
		err = rows.Scan(&item.ID, &item.UUID, &item.Category.UUID, &item.Category.Name, &item.Category.CreatedAt, &item.Category.UpdatedAt, &item.Name, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		items = append(items, item)
	}
	return items, nil
}
