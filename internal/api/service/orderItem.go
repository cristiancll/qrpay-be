package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderItem interface {
}

type orderItem struct {
	pool *pgxpool.Pool
	repo repository.OrderItem
}

func NewOrderItem(pool *pgxpool.Pool, r repository.OrderItem) OrderItem {
	return &orderItem{
		pool: pool,
		repo: r,
	}
}
