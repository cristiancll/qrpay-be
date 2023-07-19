package service

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Transaction[E any](ctx context.Context, pool *pgxpool.Pool, function func(pgx.Tx) (E, error)) (E, error) {
	var entity E
	tx, err := pool.Begin(ctx)
	if err != nil {
		return entity, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	entity, err = function(tx)
	if err != nil {
		return entity, errs.Wrap(err, errMsg.FailedExecuteTransaction)
	}
	err = tx.Commit(ctx)

	if err != nil {
		return entity, errs.New(err, errCode.Internal)
	}
	return entity, nil
}

func TransactionReturnList[E any](ctx context.Context, pool *pgxpool.Pool, function func(pgx.Tx) ([]E, error)) ([]E, error) {
	var entities []E
	tx, err := pool.Begin(ctx)
	if err != nil {
		return entities, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)

	entities, err = function(tx)
	if err != nil {
		return entities, errs.Wrap(err, errMsg.FailedExecuteTransaction)
	}
	err = tx.Commit(ctx)

	if err != nil {
		return entities, errs.New(err, errCode.Internal)
	}
	return entities, nil
}

func TransactionVoid(ctx context.Context, pool *pgxpool.Pool, function func(pgx.Tx) error) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)
	err = function(tx)
	if err != nil {
		return errs.Wrap(err, errMsg.FailedExecuteTransaction)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return errs.New(err, errCode.Internal)
	}
	return nil
}

func TransactionTuple[E any, F any](ctx context.Context, pool *pgxpool.Pool, function func(pgx.Tx) (E, F, error)) (E, F, error) {
	var entity E
	var entityTwo F
	tx, err := pool.Begin(ctx)
	if err != nil {
		return entity, entityTwo, errs.New(err, errCode.Internal)
	}
	defer tx.Rollback(ctx)
	entity, entityTwo, err = function(tx)
	if err != nil {
		return entity, entityTwo, errs.Wrap(err, errMsg.FailedExecuteTransaction)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return entity, entityTwo, errs.New(err, errCode.Internal)
	}
	return entity, entityTwo, nil
}
