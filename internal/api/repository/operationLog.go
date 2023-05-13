package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *operationLog) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
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
