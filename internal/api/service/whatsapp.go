package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WhatsApp interface {
}

type whatsApp struct {
	pool      *pgxpool.Pool
	opLogRepo repository.OperationLog
}

func NewWhatsApp(pool *pgxpool.Pool, opLogRepo repository.OperationLog) WhatsApp {
	return &whatsApp{
		pool:      pool,
		opLogRepo: opLogRepo,
	}
}
