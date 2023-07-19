package repository

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Category interface {
	Migrater
	TCRUDer[model.Category]
}

type category struct {
	db *pgxpool.Pool
}

func NewCategory(db *pgxpool.Pool) Category {
	return &category{db: db}
}

func (r *category) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS categories (
				id SERIAL PRIMARY KEY, 
				uuid VARCHAR(255) NOT NULL, 
				name VARCHAR(255) NOT NULL UNIQUE, 
				created_at TIMESTAMP NOT NULL, 
				updated_at TIMESTAMP NOT NULL);`
	return migrate(ctx, r.db, query)
}

func (r *category) TCreate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	category.UUID = uuid.New().String()
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()
	query := "INSERT INTO categories (uuid, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	id, err := tCreate(ctx, tx, query, category.UUID, category.Name, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return errs.Wrap(err, errMsg.FailedCreateCategory, category)
	}
	category.ID = id
	return nil
}

func (r *category) TUpdate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	category.UpdatedAt = time.Now().UTC()
	query := "UPDATE categories SET name = $2, updated_at = $3 WHERE id = $1"
	return tUpdate(ctx, tx, query, category.ID, category.Name, category.UpdatedAt)
}

func (r *category) TDelete(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	query := "DELETE FROM categories WHERE id = $1"
	return tDelete(ctx, tx, query, category.ID)
}

func (r *category) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Category, error) {
	query := "SELECT id, uuid, name, created_at, updated_at FROM categories WHERE id = $1"
	return tGet[model.Category](ctx, tx, query, id)
}

func (r *category) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Category, error) {
	query := "SELECT id, uuid, name, created_at, updated_at FROM categories WHERE uuid = $1"
	return tGet[model.Category](ctx, tx, query, uuid)
}

func (r *category) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Category, error) {
	query := "SELECT id, uuid, name, created_at, updated_at FROM categories"
	return tGetAll[model.Category](ctx, tx, query)
}
