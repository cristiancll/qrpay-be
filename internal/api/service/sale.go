package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Sale interface {
}

type sale struct {
	pool *pgxpool.Pool
	repo repository.Sale
}

func NewSale(pool *pgxpool.Pool, r repository.Sale) Sale {
	return &sale{
		pool: pool,
		repo: r,
	}
}
