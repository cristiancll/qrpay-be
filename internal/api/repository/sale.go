package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type Sale interface {
	Migrater
	TCRUDer[model.Sale]
}

type sale struct {
	db *pgxpool.Pool
}

func NewSale(db *pgxpool.Pool) Sale {
	return &sale{db: db}
}

const (
	createSaleTableQuery = `CREATE TABLE IF NOT EXISTS sales (
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		user_id INT NOT NULL REFERENCES users(id),
    		seller_id INT NOT NULL REFERENCES users(id),
    		total INT NOT NULL,
    		paid BOOLEAN NOT NULL DEFAULT FALSE,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	createSaleQuery = "INSERT INTO sales (uuid, user_id, seller_id, total, paid, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
)

func (r *sale) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createSaleTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *sale) TCreate(ctx context.Context, tx pgx.Tx, sale *model.Sale) error {
	sale.UUID = uuid.New().String()
	sale.CreatedAt = time.Now().UTC()
	sale.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createSaleQuery, sale.UUID, sale.User.ID, sale.Seller.ID, sale.Total, sale.Paid, sale.CreatedAt, sale.UpdatedAt).Scan(&sale.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *sale) TUpdate(ctx context.Context, tx pgx.Tx, sale *model.Sale) error {
	//TODO implement me
	panic("implement me")
}

func (r *sale) TDelete(ctx context.Context, tx pgx.Tx, sale *model.Sale) error {
	//TODO implement me
	panic("implement me")
}

func (r *sale) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (r *sale) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Sale, error) {
	//TODO implement me
	panic("implement me")
}

func (r *sale) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Sale, error) {
	//TODO implement me
	panic("implement me")
}
