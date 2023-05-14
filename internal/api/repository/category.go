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

const (
	createCategoryTable = `CREATE TABLE IF NOT EXISTS categories (
								id SERIAL PRIMARY KEY, 
								uuid VARCHAR(255) NOT NULL, 
								name VARCHAR(255) NOT NULL, 
								created_at TIMESTAMP NOT NULL, 
								updated_at TIMESTAMP NOT NULL);`
)

func (r *category) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createCategoryTable)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *category) TCreate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *category) TUpdate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *category) TDelete(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *category) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *category) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *category) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Category, error) {
	//TODO implement me
	panic("implement me")
}
