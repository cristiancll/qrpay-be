package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Order interface {
	Migrater
	TCRUDer[model.Order]
}

type order struct {
	db *pgxpool.Pool
}

func NewOrder(db *pgxpool.Pool) Order {
	return &order{db: db}
}

func (r *order) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *order) TCreate(ctx context.Context, tx pgx.Tx, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r *order) TUpdate(ctx context.Context, tx pgx.Tx, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r *order) TDelete(ctx context.Context, tx pgx.Tx, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r *order) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *order) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *order) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Order, error) {
	//TODO implement me
	panic("implement me")
}
