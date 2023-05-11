package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type AuthRepository interface {
	Migrater
	TCRUDer[model.Auth]
	VerifyUser(context.Context, pgx.Tx, *model.User) error
}

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{db: db}
}

const (
	createAuthTableQuery = "CREATE TABLE IF NOT EXISTS auths (id SERIAL PRIMARY KEY, user_id BIGINT NOT NULL REFERENCES users(id), password VARCHAR(255) NOT NULL, verified BOOLEAN NOT NULL DEFAULT TRUE, disabled BOOLEAN NOT NULL DEFAULT FALSE, reset_token VARCHAR(255), last_login TIMESTAMP, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL)"
	createAuthQuery      = "INSERT INTO auths (user_id, password, verified, disabled, reset_token, last_login, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, now(), now()) RETURNING id, created_at, updated_at"
	updateAuthQuery      = "UPDATE auths SET password = $2, verified = $3, disabled = $4, reset_token = $5, last_login = $6, updated_at = now() WHERE id = $1 RETURNING updated_at"
	deleteAuthQuery      = "DELETE FROM auths WHERE id = $1"
	getAuthByIDQuery     = "SELECT id, user_id, password, verified, disabled, reset_token, last_login, created_at, updated_at FROM auths WHERE id = $1"
	verifyUserQuery      = "UPDATE auths SET verified = TRUE WHERE user_id = $1"
)

func (r *authRepository) VerifyUser(ctx context.Context, tx pgx.Tx, user *model.User) error {
	_, err := tx.Exec(ctx, verifyUserQuery, user.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *authRepository) TCreate(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	var (
		id        int64
		createdAt time.Time
		updatedAt time.Time
	)
	row := tx.QueryRow(ctx, createAuthQuery, auth.UserID, auth.Password, auth.Verified, auth.Disabled, auth.ResetToken, auth.LastLogin)
	err := row.Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	auth.ID = id
	auth.CreatedAt = createdAt
	auth.UpdatedAt = updatedAt
	return nil
}

func (r *authRepository) TUpdate(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	var (
		updatedAt time.Time
	)
	row := tx.QueryRow(ctx, updateAuthQuery, auth.ID, auth.Password, auth.Verified, auth.Disabled, auth.ResetToken, auth.LastLogin)
	err := row.Scan(&updatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	auth.UpdatedAt = updatedAt
	return nil
}

func (r *authRepository) TDelete(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	_, err := tx.Exec(ctx, deleteAuthQuery, auth.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *authRepository) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Auth, error) {
	auth := &model.Auth{}
	row := tx.QueryRow(ctx, getAuthByIDQuery, id)
	err := row.Scan(&auth.ID, &auth.UserID, &auth.Password, &auth.Verified, &auth.Disabled, &auth.ResetToken, &auth.LastLogin, &auth.CreatedAt, &auth.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return auth, nil
}

func (r *authRepository) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Auth, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (r *authRepository) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Auth, error) {
	return nil, status.Error(codes.Unimplemented, "method not implemented")
}

func (r *authRepository) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createAuthTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}
