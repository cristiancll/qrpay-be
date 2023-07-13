package repository

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type OperationLog interface {
	Migrater
	Create(ctx context.Context, log *model.OperationLog) error
}

type operationLog struct {
	db *pgxpool.Pool
}

func NewOperationLog(db *pgxpool.Pool) OperationLog {
	return &operationLog{db: db}
}

func (r *operationLog) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS operation_logs(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		userId INT NOT NULL REFERENCES users(id),
    		sellerId INT NOT NULL REFERENCES users(id),
    		operation VARCHAR(255) NOT NULL,
    		operationId INT NOT NULL,
    		metadata JSONB,
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}

func (r *operationLog) Create(ctx context.Context, log *model.OperationLog) error {
	log.UUID = uuid.New().String()
	log.CreatedAt = time.Now().UTC()
	log.UpdatedAt = time.Now().UTC()
	log.Metadata = "{}"
	query := `INSERT INTO operation_logs(uuid, userId, sellerId, operation, operationId, metadata, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	id, err := create(ctx, r.db, query, log.UUID, log.User.ID, log.Seller.ID, log.Operation, log.OperationId, log.Metadata, log.CreatedAt, log.UpdatedAt)
	if err != nil {
		return errs.Wrap(err, "")
	}
	log.ID = id
	return nil
}
