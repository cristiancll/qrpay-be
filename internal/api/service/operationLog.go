package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OperationLogService interface {
}

type operationLogService struct {
	pool *pgxpool.Pool
	repo repository.OperationLogRepository
}

func NewOperationLogService(pool *pgxpool.Pool, r repository.OperationLogRepository) OperationLogService {
	return &operationLogService{
		pool: pool,
		repo: r,
	}
}
