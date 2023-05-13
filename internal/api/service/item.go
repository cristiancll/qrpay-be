package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Item interface {
}

type item struct {
	pool *pgxpool.Pool
	repo repository.Item
}

func NewItem(pool *pgxpool.Pool, r repository.Item) Item {
	return &item{
		pool: pool,
		repo: r,
	}
}
