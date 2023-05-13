package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OperationLog interface {
}

type operationLog struct {
	pool *pgxpool.Pool
	repo repository.OperationLog
}

func NewOperationLog(pool *pgxpool.Pool, r repository.OperationLog) OperationLog {
	return &operationLog{
		pool: pool,
		repo: r,
	}
}
