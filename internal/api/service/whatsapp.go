package service

import (
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/wpp"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WhatsAppService interface {
}

type whatsAppService struct {
	pool *pgxpool.Pool
	repo repository.WhatsAppRepository
	wpp  wpp.WhatsAppClient
}

func NewWhatsAppService(pool *pgxpool.Pool, wpp wpp.WhatsAppClient, repo repository.WhatsAppRepository) WhatsAppService {
	return &whatsAppService{
		pool: pool,
		repo: repo,
		wpp:  wpp,
	}
}
