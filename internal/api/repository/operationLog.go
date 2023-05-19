package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type OperationLog interface {
	Migrater
	Creater[model.OperationLog]
}

type operationLog struct {
	db *pgxpool.Pool
}

func NewOperationLog(db *pgxpool.Pool) OperationLog {
	return &operationLog{db: db}
}

const (
	createOperationLogTable = `CREATE TABLE IF NOT EXISTS operation_logs(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		userId INT NOT NULL REFERENCES users(id),
    		sellerId INT NOT NULL REFERENCES users(id),
    		operation VARCHAR(255) NOT NULL,
    		operationId INT NOT NULL,
    		metadata JSONB,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	createOperationLogQuery = `INSERT INTO operation_logs(uuid, userId, sellerId, operation, operationId, metadata, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
)

func (r *operationLog) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createOperationLogTable)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *operationLog) Create(ctx context.Context, log *model.OperationLog) error {
	log.UUID = uuid.New().String()
	log.CreatedAt = time.Now().UTC()
	log.UpdatedAt = time.Now().UTC()
	log.Metadata = "{}"

	row := r.db.QueryRow(ctx, createOperationLogQuery, log.UUID, log.User.ID, log.Seller.ID, log.Operation, log.OperationId, log.Metadata, log.CreatedAt, log.UpdatedAt)
	err := row.Scan(&log.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}
