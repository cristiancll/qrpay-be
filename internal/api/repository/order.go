package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository interface {
	Migrater
	TCRUDer[model.Order]
}

type orderRepository struct {
	db *pgxpool.Pool
}

func (r *orderRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) TCreate(ctx context.Context, tx pgx.Tx, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) TUpdate(ctx context.Context, tx pgx.Tx, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) TDelete(ctx context.Context, tx pgx.Tx, order *model.Order) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Order, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db *pgxpool.Pool) OrderRepository {
	return &orderRepository{db: db}
}
