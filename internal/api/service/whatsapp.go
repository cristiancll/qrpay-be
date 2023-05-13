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

type WhatsApp interface {
	Get(ctx context.Context, uuid string) (*model.WhatsApp, error)
	GetAll(ctx context.Context) ([]*model.WhatsApp, error)
}

type whatsApp struct {
	pool *pgxpool.Pool
	repo repository.WhatsApp
	wpp  wpp.WhatsAppSystem
}

func NewWhatsApp(pool *pgxpool.Pool, wpp wpp.WhatsAppSystem, repo repository.WhatsApp) WhatsApp {
	return &whatsApp{
		pool: pool,
		repo: repo,
		wpp:  wpp,
	}
}

func (s *whatsApp) Get(ctx context.Context, uuid string) (*model.WhatsApp, error) {
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

func (s *whatsApp) GetAll(ctx context.Context) ([]*model.WhatsApp, error) {
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
