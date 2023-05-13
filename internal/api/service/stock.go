package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StockService interface {
}

type stockService struct {
	pool *pgxpool.Pool
	repo repository.StockRepository
}

func NewStockService(pool *pgxpool.Pool, r repository.StockRepository) StockService {
	return &stockService{
		pool: pool,
		repo: r,
	}
}
