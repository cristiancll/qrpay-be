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

type Retrieval interface {
	Migrater
	TCRUDer[model.Retrieval]
	TGetAllPendingDelivery(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error)
	TGetAllDelivered(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error)
	TGetAllByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.Retrieval, error)
}

func NewRetrieval(db *pgxpool.Pool) Retrieval {
	return &retrieval{db: db}
}

type retrieval struct {
	db *pgxpool.Pool
}

func (r *retrieval) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS retrievals(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		user_id INT NOT NULL REFERENCES users(id),
    		seller_id INT NOT NULL REFERENCES users(id),
    		sale_item_id INT NOT NULL REFERENCES sale_items(id),
    		delivered BOOLEAN DEFAULT TRUE,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}

func (r *retrieval) TCreate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	retrieval.UUID = uuid.New().String()
	retrieval.CreatedAt = time.Now().UTC()
	retrieval.UpdatedAt = time.Now().UTC()
	query := `INSERT INTO retrievals(uuid, user_id, seller_id, sale_item_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	id, err := tCreate(ctx, tx, query, retrieval.UUID, retrieval.User.ID, retrieval.Seller.ID, retrieval.SaleItem.ID, retrieval.CreatedAt, retrieval.UpdatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	retrieval.ID = id
	return nil
}

func (r *retrieval) TUpdate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	query := `UPDATE retrievals SET delivered = $2, updated_at = $3 WHERE id = $1`
	retrieval.UpdatedAt = time.Now().UTC()
	return tUpdate(ctx, tx, query, retrieval.ID, retrieval.Delivered, retrieval.UpdatedAt)
}

func (r *retrieval) TDelete(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	query := `DELETE FROM retrievals WHERE id = $1`
	return tDelete(ctx, tx, query, retrieval.ID)
}

func (r *retrieval) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Retrieval, error) {
	query := `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE id = $1`
	return tGet[model.Retrieval](ctx, tx, query, id)
}

func (r *retrieval) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Retrieval, error) {
	query := `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE uuid = $1`
	return tGet[model.Retrieval](ctx, tx, query, uuid)
}

func (r *retrieval) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	query := `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals`
	return tGetAll[model.Retrieval](ctx, tx, query)
}

func (r *retrieval) TGetAllPendingDelivery(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	query := `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE delivered = FALSE`
	return tGetAll[model.Retrieval](ctx, tx, query)
}

func (r *retrieval) TGetAllDelivered(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	query := `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE delivered = TRUE`
	return tGetAll[model.Retrieval](ctx, tx, query)
}

func (r *retrieval) TGetAllByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.Retrieval, error) {
	query := `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE user_id = $1`
	return tGetAll[model.Retrieval](ctx, tx, query, user.ID)
}
