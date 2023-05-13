package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepository interface {
	Migrater
	TCRUDer[model.Category]
}

type categoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) TCreate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) TUpdate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) TDelete(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *categoryRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Category, error) {
	//TODO implement me
	panic("implement me")
}
