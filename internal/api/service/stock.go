package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Stock interface {
}

type stock struct {
	pool *pgxpool.Pool
	repo repository.Stock
}

func NewStock(pool *pgxpool.Pool, r repository.Stock) Stock {
	return &stock{
		pool: pool,
		repo: r,
	}
}
