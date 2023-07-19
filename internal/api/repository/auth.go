package repository

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Auth interface {
	Migrater
	TCreater[model.Auth]
	TUpdater[model.Auth]
	TDeleter[model.Auth]
	TGetterById[model.Auth]
	VerifyUser(context.Context, pgx.Tx, *model.User) error
}

type auth struct {
	db *pgxpool.Pool
}

func NewAuth(db *pgxpool.Pool) Auth {
	return &auth{db: db}
}

func (r *auth) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS auths (
				id SERIAL PRIMARY KEY, 
				user_id INT NOT NULL REFERENCES users(id), 
				password VARCHAR(255) NOT NULL, 
				verified BOOLEAN NOT NULL DEFAULT FALSE, 
				disabled BOOLEAN NOT NULL DEFAULT FALSE, 
				reset_token VARCHAR(255), 
				last_login TIMESTAMP, 
				created_at TIMESTAMP NOT NULL, 
				updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}

func (r *auth) VerifyUser(ctx context.Context, tx pgx.Tx, user *model.User) error {
	user.UpdatedAt = time.Now().UTC()
	query := "UPDATE auths SET verified = TRUE, updated_at = $2 WHERE user_id = $1"
	return tUpdate(ctx, tx, query, user.ID, user.UpdatedAt)
}

func (r *auth) TCreate(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	auth.CreatedAt = time.Now().UTC()
	auth.UpdatedAt = time.Now().UTC()
	query := "INSERT INTO auths (user_id, password, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	id, err := tCreate(ctx, tx, query, auth.UserID, auth.Password, auth.CreatedAt, auth.UpdatedAt)
	if err != nil {
		return errs.Wrap(err, errMsg.FailedCreateAuth, auth.UserID)
	}
	auth.ID = id
	return nil
}

func (r *auth) TUpdate(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	auth.UpdatedAt = time.Now().UTC()
	query := "UPDATE auths SET password = $2, verified = $3, disabled = $4, reset_token = $5, last_login = $6, updated_at = $7 WHERE id = $1"
	return tUpdate(ctx, tx, query, auth.ID, auth.Password, auth.Verified, auth.Disabled, auth.ResetToken, auth.LastLogin, auth.UpdatedAt)
}

func (r *auth) TDelete(ctx context.Context, tx pgx.Tx, auth *model.Auth) error {
	query := "DELETE FROM auths WHERE id = $1"
	return tDelete(ctx, tx, query, auth.ID)
}

func (r *auth) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Auth, error) {
	query := "SELECT id, user_id, password, verified, disabled, reset_token, last_login, created_at, updated_at FROM auths WHERE id = $1"
	return tGet[model.Auth](ctx, tx, query, id)
}
