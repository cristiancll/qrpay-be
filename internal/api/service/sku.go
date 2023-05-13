package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SKU interface {
}

type sku struct {
	pool *pgxpool.Pool
	repo repository.SKU
}

func NewSKU(pool *pgxpool.Pool, r repository.SKU) SKU {
	return &sku{
		pool: pool,
		repo: r,
	}
}
