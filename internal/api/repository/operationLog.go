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

type OperationLog interface {
	Migrater
	TCRUDer[model.OperationLog]
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
    		createdAt TIMESTAMP NOT NULL,
    		updatedAt TIMESTAMP NOT NULL)`
)

func (r *operationLog) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createOperationLogTable)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *operationLog) TCreate(ctx context.Context, tx pgx.Tx, log *model.OperationLog) error {
	//TODO implement me
	panic("implement me")
}

func (r *operationLog) TUpdate(ctx context.Context, tx pgx.Tx, log *model.OperationLog) error {
	//TODO implement me
	panic("implement me")
}

func (r *operationLog) TDelete(ctx context.Context, tx pgx.Tx, log *model.OperationLog) error {
	//TODO implement me
	panic("implement me")
}

func (r *operationLog) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.OperationLog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *operationLog) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.OperationLog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *operationLog) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.OperationLog, error) {
	//TODO implement me
	panic("implement me")
}
