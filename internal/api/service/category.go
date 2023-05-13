package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Category interface {
}

type category struct {
	pool *pgxpool.Pool
	repo repository.Category
}

func NewCategory(pool *pgxpool.Pool, r repository.Category) Category {
	return &category{
		pool: pool,
		repo: r,
	}
}
