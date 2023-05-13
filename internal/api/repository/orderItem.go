package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderItemRepository interface {
	Migrater
	TCRUDer[model.OrderItem]
}

type orderItemRepository struct {
	db *pgxpool.Pool
}

func (r *orderItemRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItemRepository) TCreate(ctx context.Context, tx pgx.Tx, item *model.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItemRepository) TUpdate(ctx context.Context, tx pgx.Tx, item *model.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItemRepository) TDelete(ctx context.Context, tx pgx.Tx, item *model.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (r *orderItemRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderItemRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (r *orderItemRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func NewOrderItemRepository(db *pgxpool.Pool) OrderItemRepository {
	return &orderItemRepository{db: db}
}
