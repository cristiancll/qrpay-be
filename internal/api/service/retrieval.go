package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Retrieval interface {
}

type retrieval struct {
	pool *pgxpool.Pool
	repo repository.Retrieval
}

func NewRetrieval(pool *pgxpool.Pool, r repository.Retrieval) Retrieval {
	return &retrieval{
		pool: pool,
		repo: r,
	}
}
