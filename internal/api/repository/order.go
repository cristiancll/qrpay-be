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

const (
	createOrderTableQuery = `CREATE TABLE IF NOT EXISTS "order" (
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		user_id INT NOT NULL REFERENCES "user" (id),
    		seller_id INT NOT NULL REFERENCES "user" (id),
    		amount INT NOT NULL,
    		paid BOOLEAN NOT NULL DEFAULT FALSE,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
)

func (r *order) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createOrderTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
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
