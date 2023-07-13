package repository

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Sale interface {
	Migrater
	TCreater[model.Sale]
}

type sale struct {
	db *pgxpool.Pool
}

func NewSale(db *pgxpool.Pool) Sale {
	return &sale{db: db}
}

func (r *sale) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS sales (
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		user_id INT NOT NULL REFERENCES users(id),
    		seller_id INT NOT NULL REFERENCES users(id),
    		total INT NOT NULL,
    		paid BOOLEAN NOT NULL DEFAULT FALSE,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}

func (r *sale) TCreate(ctx context.Context, tx pgx.Tx, sale *model.Sale) error {
	sale.UUID = uuid.New().String()
	sale.CreatedAt = time.Now().UTC()
	sale.UpdatedAt = time.Now().UTC()
	query := "INSERT INTO sales (uuid, user_id, seller_id, total, paid, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	id, err := tCreate(ctx, tx, query, sale.UUID, sale.User.ID, sale.Seller.ID, sale.Total, sale.Paid, sale.CreatedAt, sale.UpdatedAt)
	if err != nil {
		return errs.Wrap(err, "")
	}
	sale.ID = id
	return nil
}
