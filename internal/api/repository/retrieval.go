package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RetrievalRepository interface {
	Migrater
	TCRUDer[model.Retrieval]
}

type retrievalRepository struct {
	db *pgxpool.Pool
}

func (r *retrievalRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *retrievalRepository) TCreate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	//TODO implement me
	panic("implement me")
}

func (r *retrievalRepository) TUpdate(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	//TODO implement me
	panic("implement me")
}

func (r *retrievalRepository) TDelete(ctx context.Context, tx pgx.Tx, retrieval *model.Retrieval) error {
	//TODO implement me
	panic("implement me")
}

func (r *retrievalRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Retrieval, error) {
	//TODO implement me
	panic("implement me")
}

func (r *retrievalRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Retrieval, error) {
	//TODO implement me
	panic("implement me")
}

func (r *retrievalRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Retrieval, error) {
	//TODO implement me
	panic("implement me")
}

func NewRetrievalRepository(db *pgxpool.Pool) RetrievalRepository {
	return &retrievalRepository{db: db}
}
