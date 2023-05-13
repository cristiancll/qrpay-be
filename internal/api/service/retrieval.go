package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RetrievalService interface {
}

type retrievalService struct {
	pool *pgxpool.Pool
	repo repository.RetrievalRepository
}

func NewRetrievalService(pool *pgxpool.Pool, r repository.RetrievalRepository) RetrievalService {
	return &retrievalService{
		pool: pool,
		repo: r,
	}
}
