package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	//TODO implement me
	panic("implement me")
}

func (r *item) TCreate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r *item) TUpdate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r *item) TDelete(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r *item) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (r *item) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (r *item) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Item, error) {
	//TODO implement me
	panic("implement me")
}
