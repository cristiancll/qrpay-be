package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemRepository interface {
	Migrater
	TCRUDer[model.Item]
}

type itemRepository struct {
	db *pgxpool.Pool
}

func NewItemRepository(db *pgxpool.Pool) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Migrate(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *itemRepository) TCreate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r *itemRepository) TUpdate(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r *itemRepository) TDelete(ctx context.Context, tx pgx.Tx, item *model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (r *itemRepository) TGetById(ctx context.Context, tx pgx.Tx, i int64) (*model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (r *itemRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, s string) (*model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (r *itemRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Item, error) {
	//TODO implement me
	panic("implement me")
}
