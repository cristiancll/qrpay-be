package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SKUService interface {
}

type skuService struct {
	pool *pgxpool.Pool
	repo repository.SKURepository
}

func NewSKUService(pool *pgxpool.Pool, r repository.SKURepository) SKUService {
	return &skuService{
		pool: pool,
		repo: r,
	}
}
