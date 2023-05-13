package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Order interface {
}

type order struct {
	pool *pgxpool.Pool
	repo repository.Order
}

func NewOrder(pool *pgxpool.Pool, r repository.Order) Order {
	return &order{
		pool: pool,
		repo: r,
	}
}
