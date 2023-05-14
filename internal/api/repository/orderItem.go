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

const (
	createOrderItemTableQuery = `CREATE TABLE IF NOT EXISTS order_item(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		order_id INT NOT NULL REFERENCES "order" (id),
    		sku_id INT NOT NULL REFERENCES "sku" (id),
    		quantity INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
)

func (r *orderItem) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createOrderItemTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
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
