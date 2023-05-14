package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SaleItem interface {
}

type saleItem struct {
	pool *pgxpool.Pool
	repo repository.SaleItem
}

func NewSaleItem(pool *pgxpool.Pool, r repository.SaleItem) SaleItem {
	return &saleItem{
		pool: pool,
		repo: r,
	}
}
