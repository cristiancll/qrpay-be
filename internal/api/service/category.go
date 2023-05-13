package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryService interface {
}

type categoryService struct {
	pool *pgxpool.Pool
	repo repository.CategoryRepository
}

func NewCategoryService(pool *pgxpool.Pool, r repository.CategoryRepository) CategoryService {
	return &categoryService{
		pool: pool,
		repo: r,
	}
}
