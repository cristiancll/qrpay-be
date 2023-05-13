package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderItemService interface {
}

type orderItemService struct {
	pool *pgxpool.Pool
	repo repository.OrderItemRepository
}

func NewOrderItemService(pool *pgxpool.Pool, r repository.OrderItemRepository) OrderItemService {
	return &orderItemService{
		pool: pool,
		repo: r,
	}
}
