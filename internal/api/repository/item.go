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

const (
	createItemTableQuery = `CREATE TABLE IF NOT EXISTS item(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		category_id INT NOT NULL REFERENCES category(id),
    		name VARCHAR(255) NOT NULL,
    		create_at TIMESTAMP NOT NULL,
    		update_at TIMESTAMP NOT NULL)`
)

func (r *item) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createItemTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
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
