package service

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	opLogRepo    repository.OperationLog
}

func NewRetrieval(pool *pgxpool.Pool, r repository.Retrieval, userRepo repository.User, saleItemRepo repository.SaleItem, opLogRepo repository.OperationLog) Retrieval {
	return &retrieval{
		pool:         pool,
		repo:         r,
		userRepo:     userRepo,
		saleItemRepo: saleItemRepo,
		opLogRepo:    opLogRepo,
	}
}

func (r *retrieval) Create(ctx context.Context, userUUID string, sellerUUID string, saleItemUUIDs []string) error {
	return TransactionVoid(ctx, r.pool, func(tx pgx.Tx) error {
		// Validates user
		user, err := r.userRepo.TGetByUUID(ctx, tx, userUUID)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetUser, userUUID)
		}

		// Validates seller
		seller, err := r.userRepo.TGetByUUID(ctx, tx, sellerUUID)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetUser, sellerUUID)
		}

		// Validates sale items
		saleItems, err := r.saleItemRepo.TGetAllByUUIDs(ctx, tx, saleItemUUIDs)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetAllSaleItem, saleItemUUIDs)
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
				return errs.Wrap(err, errMsg.FailedCreateRetrieval, user, seller, saleItem)
			}
			opLog := &model.OperationLog{
				User:        *user,
				Seller:      *seller,
				Operation:   "Retrieval",
				OperationId: retrieval.ID,
			}
			_ = r.opLogRepo.Create(context.Background(), opLog)
		}

		//go r.wpp.SendText(user, user.NewRetrieval(saleItems)) // TODO
		return nil
	})
}

func (r *retrieval) Update(ctx context.Context, uuid string, delivered bool) (*model.Retrieval, error) {
	return Transaction[*model.Retrieval, error](ctx, r.pool, func(tx pgx.Tx) (*model.Retrieval, error) {
		retrieval, err := r.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetRetrieval, uuid)
		}
		retrieval.Delivered = delivered
		err = r.repo.TUpdate(ctx, tx, retrieval)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedUpdateRetrieval, retrieval)
		}

		return retrieval, nil
	})
}

func (r *retrieval) Delete(ctx context.Context, uuid string) error {
	return TransactionVoid(ctx, r.pool, func(tx pgx.Tx) error {
		retrieval, err := r.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedGetRetrieval, uuid)
		}

		err = r.repo.TDelete(ctx, tx, retrieval)
		if err != nil {
			return errs.Wrap(err, errMsg.FailedDeleteRetrieval, retrieval)
		}
		return nil
	})
}

func (r *retrieval) Get(ctx context.Context, uuid string) (*model.Retrieval, error) {
	return Transaction[*model.Retrieval, error](ctx, r.pool, func(tx pgx.Tx) (*model.Retrieval, error) {
		retrieval, err := r.repo.TGetByUUID(ctx, tx, uuid)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetRetrieval, uuid)
		}

		return retrieval, nil
	})
}

func (r *retrieval) List(ctx context.Context) ([]*model.Retrieval, error) {
	return TransactionReturnList[*model.Retrieval](ctx, r.pool, func(tx pgx.Tx) ([]*model.Retrieval, error) {
		retrievals, err := r.repo.TGetAll(ctx, tx)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetAllRetrieval)
		}
		return retrievals, nil
	})
}

func (r *retrieval) ListByUser(ctx context.Context, userUUID string) ([]*model.Retrieval, error) {
	return TransactionReturnList[*model.Retrieval](ctx, r.pool, func(tx pgx.Tx) ([]*model.Retrieval, error) {
		user, err := r.userRepo.TGetByUUID(ctx, tx, userUUID)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetUser, userUUID)
		}

		retrievals, err := r.repo.TGetAllByUser(ctx, tx, user)
		if err != nil {
			return nil, errs.Wrap(err, errMsg.FailedGetAllRetrieval, user)
		}
		return retrievals, nil
	})
}
