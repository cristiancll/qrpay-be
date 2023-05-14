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

type Retrieval interface {
	Migrater
	TCRUDer[model.Retrieval]
}

func NewRetrieval(db *pgxpool.Pool) Retrieval {
	return &retrieval{db: db}
}

type retrieval struct {
	db *pgxpool.Pool
}

const (
	createRetrievalTableQuery = `CREATE TABLE IF NOT EXISTS retrievals(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		user_id INT NOT NULL REFERENCES users(id),
    		seller_id INT NOT NULL REFERENCES users(id),
    		sale_item_id INT NOT NULL REFERENCES sale_items(id),
    		quantity INT NOT NULL,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
)

func (r *retrieval) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createRetrievalTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *retrieval) TCreate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	//TODO implement me
	panic("implement me")
}

func (r *retrieval) TUpdate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	//TODO implement me
	panic("implement me")
}

func (r *retrieval) TDelete(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	//TODO implement me
	panic("implement me")
}

func (r *retrieval) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Retrieval, error) {
	//TODO implement me
	panic("implement me")
}

func (r *retrieval) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Retrieval, error) {
	//TODO implement me
	panic("implement me")
}

func (r *retrieval) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	//TODO implement me
	panic("implement me")
}
