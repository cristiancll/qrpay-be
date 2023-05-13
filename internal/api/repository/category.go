package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	//TODO implement me
	panic("implement me")
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
