package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderItem interface {
	Migrater
	TCRUDer[model.OrderItem]
}

type orderItem struct {
	db *pgxpool.Pool
}

func NewOrderItem(db *pgxpool.Pool) OrderItem {
	return &orderItem{db: db}
}

func (r *orderItem) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItem) TCreate(ctx context.Context, tx pgx.Tx, item *model.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItem) TUpdate(ctx context.Context, tx pgx.Tx, item *model.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItem) TDelete(ctx context.Context, tx pgx.Tx, item *model.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItem) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderItem) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderItem) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}
