package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OperationLogRepository interface {
	Migrater
	TCRUDer[model.OperationLog]
}

type operationLogRepository struct {
	db *pgxpool.Pool
}

func NewOperationLogRepository(db *pgxpool.Pool) OperationLogRepository {
	return &operationLogRepository{db: db}
}

func (r *operationLogRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *operationLogRepository) TCreate(ctx context.Context, tx pgx.Tx, log *model.OperationLog) error {
	//TODO implement me
	panic("implement me")
}

func (r *operationLogRepository) TUpdate(ctx context.Context, tx pgx.Tx, log *model.OperationLog) error {
	//TODO implement me
	panic("implement me")
}

func (r *operationLogRepository) TDelete(ctx context.Context, tx pgx.Tx, log *model.OperationLog) error {
	//TODO implement me
	panic("implement me")
}

func (r *operationLogRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.OperationLog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *operationLogRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.OperationLog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *operationLogRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.OperationLog, error) {
	//TODO implement me
	panic("implement me")
}
