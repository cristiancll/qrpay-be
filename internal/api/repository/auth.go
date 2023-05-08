package repository

import (
	"context"
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type AuthRepository interface {
	PrivateRepository[model.Auth]
}

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) AuthRepository {
	return &authRepository{db: db}
}

const (
	createAuthTableQuery = "CREATE TABLE IF NOT EXISTS auths (id SERIAL PRIMARY KEY, user_id BIGINT NOT NULL REFERENCES users(id), password VARCHAR(255) NOT NULL, verified BOOLEAN NOT NULL DEFAULT TRUE, disabled BOOLEAN NOT NULL DEFAULT FALSE, locked BOOLEAN NOT NULL DEFAULT FALSE, activation_token VARCHAR(255), reset_token VARCHAR(255), last_login TIMESTAMP, reset_expiration TIMESTAMP, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL)"
	createAuthQuery      = "INSERT INTO auths (user_id, password, verified, disabled, locked, activation_token, reset_token, last_login, reset_expiration, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now()) RETURNING id, created_at, updated_at"
	updateAuthQuery      = "UPDATE auths SET password = $2, verified = $3, disabled = $4, locked = $5, activation_token = $6, reset_token = $7, last_login = $8, reset_expiration = $9, updated_at = now() WHERE id = $1 RETURNING updated_at"
	deleteAuthQuery      = "DELETE FROM auths WHERE id = $1"
	getAuthByIDQuery     = "SELECT id, user_id, password, verified, disabled, locked, activation_token, reset_token, last_login, reset_expiration, created_at, updated_at FROM auths WHERE id = $1"
)

func (r *authRepository) Create(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	var (
		id        int64
		createdAt time.Time
		updatedAt time.Time
	)
	row := tx.QueryRow(ctx, createAuthQuery, auth.UserID, auth.Password, auth.Verified, auth.Disabled, auth.Locked, auth.ActivationToken, auth.ResetToken, auth.LastLogin, auth.ResetExpiration)
	err := row.Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return fmt.Errorf("error creating auth: %w", err)
	}
	auth.ID = id
	auth.CreatedAt = createdAt
	auth.UpdatedAt = updatedAt
	return nil
}

func (r *authRepository) Update(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	var (
		updatedAt time.Time
	)
	row := tx.QueryRow(ctx, updateAuthQuery, auth.ID, auth.Password, auth.Verified, auth.Disabled, auth.Locked, auth.ActivationToken, auth.ResetToken, auth.LastLogin, auth.ResetExpiration)
	err := row.Scan(&updatedAt)
	if err != nil {
		return fmt.Errorf("error updating auth: %w", err)
	}
	auth.UpdatedAt = updatedAt
	return nil
}

func (r *authRepository) Delete(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	_, err := tx.Exec(ctx, deleteAuthQuery, auth.ID)
	if err != nil {
		return fmt.Errorf("error deleting auth: %w", err)
	}
	return nil
}

func (r *authRepository) GetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Auth, error) {
	var auth *model.Auth
	row := tx.QueryRow(ctx, getAuthByIDQuery, id)
	err := row.Scan(&auth.ID, &auth.UserID, &auth.Password, &auth.Verified, &auth.Disabled, &auth.Locked, &auth.ActivationToken, &auth.ResetToken, &auth.LastLogin, &auth.ResetExpiration, &auth.CreatedAt, &auth.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error getting auth by id: %w", err)
	}
	return auth, nil
}

func (r *authRepository) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createAuthTableQuery)
	if err != nil {
		return fmt.Errorf("error migrating auth: %w", err)
	}
	return nil
}
