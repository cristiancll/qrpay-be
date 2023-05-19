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

const (
	createRetrievalTableQuery = `CREATE TABLE IF NOT EXISTS retrievals(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		user_id INT NOT NULL REFERENCES users(id),
    		seller_id INT NOT NULL REFERENCES users(id),
    		sale_item_id INT NOT NULL REFERENCES sale_items(id),
    		delivered BOOLEAN DEFAULT TRUE,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	createRetrievalQuery       = `INSERT INTO retrievals(uuid, user_id, seller_id, sale_item_id, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6) RETURNING id`
	updateRetrievalQuery       = `UPDATE retrievals SET delivered = $2, updated_at = $3 WHERE id = $1`
	deleteRetrievalQuery       = `DELETE FROM retrievals WHERE id = $1`
	getRetrievalByIdQuery      = `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE id = $1`
	getRetrievalByUUIDQuery    = `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE uuid = $1`
	getAllRetrievalsQuery      = `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals`
	getAllPendingDeliveryQuery = `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE delivered = FALSE`
	getAllDeliveredQuery       = `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE delivered = TRUE`
	getAllByUserIdQuery        = `SELECT (id, uuid, user_id, seller_id, sale_item_id, delivered, created_at, updated_at) FROM retrievals WHERE user_id = $1`
)

func (r *retrieval) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createRetrievalTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *retrieval) TCreate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	retrieval.UUID = uuid.New().String()
	retrieval.CreatedAt = time.Now().UTC()
	retrieval.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createRetrievalQuery, retrieval.UUID, retrieval.User.ID, retrieval.Seller.ID, retrieval.SaleItem.ID, retrieval.CreatedAt, retrieval.UpdatedAt).Scan(&retrieval.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *retrieval) TUpdate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	retrieval.UpdatedAt = time.Now().UTC()
	_, err := tx.Exec(ctx, updateRetrievalQuery, retrieval.ID, retrieval.Delivered, retrieval.UpdatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *retrieval) TDelete(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	_, err := tx.Exec(ctx, deleteRetrievalQuery, retrieval.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *retrieval) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Retrieval, error) {
	row := tx.QueryRow(ctx, getRetrievalByIdQuery, id)
	return getRetrievalFromRow(row)
}

func (r *retrieval) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Retrieval, error) {
	row := tx.QueryRow(ctx, getRetrievalByUUIDQuery, uuid)
	return getRetrievalFromRow(row)
}

func (r *retrieval) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	rows, err := tx.Query(ctx, getAllRetrievalsQuery)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	return getRetrievalsFromRows(rows)
}

func (r *retrieval) TGetAllPendingDelivery(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	rows, err := tx.Query(ctx, getAllPendingDeliveryQuery, false)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	return getRetrievalsFromRows(rows)
}

func (r *retrieval) TGetAllDelivered(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	rows, err := tx.Query(ctx, getAllDeliveredQuery, true)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	return getRetrievalsFromRows(rows)
}

func (r *retrieval) TGetAllByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.Retrieval, error) {
	rows, err := tx.Query(ctx, getAllByUserIdQuery, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	return getRetrievalsFromRows(rows)
}

func getRetrievalFromRow(row pgx.Row) (*model.Retrieval, error) {
	retrieval := &model.Retrieval{}
	err := row.Scan(&retrieval.ID, &retrieval.UUID, &retrieval.User.ID, &retrieval.Seller.ID, &retrieval.SaleItem.ID, &retrieval.Delivered, &retrieval.CreatedAt, &retrieval.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return retrieval, nil
}

func getRetrievalsFromRows(rows pgx.Rows) ([]*model.Retrieval, error) {
	var retrievals []*model.Retrieval
	for rows.Next() {
		retrieval := &model.Retrieval{}
		err := rows.Scan(&retrieval.ID, &retrieval.UUID, &retrieval.User.ID, &retrieval.Seller.ID, &retrieval.SaleItem.ID, &retrieval.Delivered, &retrieval.CreatedAt, &retrieval.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		retrievals = append(retrievals, retrieval)
	}
	return retrievals, nil
}
