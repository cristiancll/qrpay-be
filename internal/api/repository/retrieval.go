package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *retrieval) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
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
