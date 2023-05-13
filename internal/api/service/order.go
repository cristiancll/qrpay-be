package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderService interface {
}

type orderService struct {
	pool *pgxpool.Pool
	repo repository.OrderRepository
}

func NewOrderService(pool *pgxpool.Pool, r repository.OrderRepository) OrderService {
	return &orderService{
		pool: pool,
		repo: r,
	}
}
