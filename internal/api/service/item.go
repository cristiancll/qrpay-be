package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ItemService interface {
}

type itemService struct {
	pool *pgxpool.Pool
	repo repository.ItemRepository
}

func NewItemService(pool *pgxpool.Pool, r repository.ItemRepository) ItemService {
	return &itemService{
		pool: pool,
		repo: r,
	}
}
