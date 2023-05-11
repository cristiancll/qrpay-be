package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/cristiancll/qrpay-be/internal/wpp"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WhatsAppService interface {
	Get(ctx context.Context, uuid string) (*model.WhatsApp, error)
	GetAll(ctx context.Context) ([]*model.WhatsApp, error)
}

func NewWhatsAppService(pool *pgxpool.Pool, wpp wpp.WhatsAppSystem, repo repository.WhatsAppRepository) WhatsAppService {
	return &whatsAppService{
		pool: pool,
		repo: repo,
		wpp:  wpp,
	}
}

type whatsAppService struct {
	pool *pgxpool.Pool
	repo repository.WhatsAppRepository
	wpp  wpp.WhatsAppSystem
}

func (s *whatsAppService) Get(ctx context.Context, uuid string) (*model.WhatsApp, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	whats, err := s.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return whats, nil
}

func (s *whatsAppService) GetAll(ctx context.Context) ([]*model.WhatsApp, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	whats, err := s.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return whats, nil
}
