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

type Retrieval interface {
	Create(ctx context.Context, userUUID string, sellerUUID string, saleItemUUIDs []string) error
	Update(ctx context.Context, uuid string, delivered bool) (*model.Retrieval, error)
	Delete(ctx context.Context, uuid string) error
	Get(ctx context.Context, uuid string) (*model.Retrieval, error)
	List(ctx context.Context) ([]*model.Retrieval, error)
	ListByUser(ctx context.Context, userUUID string) ([]*model.Retrieval, error)
}

type retrieval struct {
	pool         *pgxpool.Pool
	repo         repository.Retrieval
	userRepo     repository.User
	saleItemRepo repository.SaleItem
	wpp          wpp.WhatsAppSystem
}

func NewRetrieval(pool *pgxpool.Pool, wpp wpp.WhatsAppSystem, r repository.Retrieval, userRepo repository.User, saleItemRepo repository.SaleItem) Retrieval {
	return &retrieval{
		pool:         pool,
		repo:         r,
		userRepo:     userRepo,
		saleItemRepo: saleItemRepo,
		wpp:          wpp,
	}
}

func (r *retrieval) Create(ctx context.Context, userUUID string, sellerUUID string, saleItemUUIDs []string) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	// Validates user
	user, err := r.userRepo.TGetByUUID(ctx, tx, userUUID)
	if err != nil {
		return err
	}

	// Validates seller
	seller, err := r.userRepo.TGetByUUID(ctx, tx, sellerUUID)
	if err != nil {
		return err
	}

	// Validates sale items
	saleItems, err := r.saleItemRepo.TGetAllByUUIDs(ctx, tx, saleItemUUIDs)
	if err != nil {
		return err
	}

	// Creates retrievals
	for _, saleItem := range saleItems {
		retrieval := &model.Retrieval{
			User:     *user,
			Seller:   *seller,
			SaleItem: *saleItem,
		}
		err = r.repo.TCreate(ctx, tx, retrieval)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	go r.wpp.SendText(user, user.NewRetrieval(saleItems))
	return nil
}

func (r *retrieval) Update(ctx context.Context, uuid string, delivered bool) (*model.Retrieval, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	retrieval, err := r.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, err
	}
	retrieval.Delivered = delivered
	err = r.repo.TUpdate(ctx, tx, retrieval)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return retrieval, nil
}

func (r *retrieval) Delete(ctx context.Context, uuid string) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	retrieval, err := r.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return err
	}

	err = r.repo.TDelete(ctx, tx, retrieval)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return nil
}

func (r *retrieval) Get(ctx context.Context, uuid string) (*model.Retrieval, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	retrieval, err := r.repo.TGetByUUID(ctx, tx, uuid)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return retrieval, nil
}

func (r *retrieval) List(ctx context.Context) ([]*model.Retrieval, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	retrievals, err := r.repo.TGetAll(ctx, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return retrievals, nil
}

func (r *retrieval) ListByUser(ctx context.Context, userUUID string) ([]*model.Retrieval, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	defer tx.Rollback(ctx)

	user, err := r.userRepo.TGetByUUID(ctx, tx, userUUID)
	if err != nil {
		return nil, err
	}

	retrievals, err := r.repo.TGetAllByUser(ctx, tx, user)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.INTERNAL_ERROR)
	}
	return retrievals, nil
}
