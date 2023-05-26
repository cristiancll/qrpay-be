package repository

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type User interface {
	Migrater
	TCRUDer[model.User]
	CountByPhone(ctx context.Context, tx pgx.Tx, phone string) (int64, error)
	GetUserByPhone(ctx context.Context, tx pgx.Tx, phone string) (*model.User, error)
	GetVerifiedList(ctx context.Context) ([]string, error)
}

type user struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) User {
	return &user{db: db}
}

func (r *user) GetVerifiedList(ctx context.Context) ([]string, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer tx.Rollback(ctx)

	query := "SELECT u.phone FROM users u INNER JOIN auths a ON u.id = a.user_id WHERE a.verified = true"
	rows, err := tx.Query(ctx, query)
	if err == pgx.ErrNoRows {
		return nil, status.Error(codes.NotFound, errors.NO_USERS_FOUND)
	} else if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	var phones []string
	for rows.Next() {
		var phone string
		err = rows.Scan(&phone)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		phones = append(phones, phone)
	}
	err = tx.Commit(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return phones, nil
}

func (r *user) GetUserByPhone(ctx context.Context, tx pgx.Tx, phone string) (*model.User, error) {
	query := "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users WHERE phone = $1"
	return tGet[model.User](ctx, tx, query, phone)
}

func (r *user) CountByPhone(ctx context.Context, tx pgx.Tx, phone string) (int64, error) {
	query := "SELECT count(*) FROM users WHERE phone = $1"
	return tCount(ctx, tx, query, phone)
}

func (r *user) TCreate(ctx context.Context, tx pgx.Tx, user *model.User) error {
	user.UUID = uuid.New().String()
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	query := "INSERT INTO users (uuid, name, role, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	id, err := tCreate(ctx, tx, query, user.UUID, user.Name, user.Role, user.Phone, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

func (r *user) TUpdate(ctx context.Context, tx pgx.Tx, user *model.User) error {
	user.UpdatedAt = time.Now().UTC()
	query := "UPDATE users SET name = $2, role = $3, phone = $4, updated_at = $5 WHERE id = $1"
	return tUpdate(ctx, tx, query, user.ID, user.Name, user.Role, user.Phone, user.UpdatedAt)
}

func (r *user) TDelete(ctx context.Context, tx pgx.Tx, user *model.User) error {
	query := "DELETE FROM users WHERE id = $1"
	return tDelete(ctx, tx, query, user.ID)
}

func (r *user) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.User, error) {
	query := "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users WHERE id = $1"
	return tGet[model.User](ctx, tx, query, id)
}

func (r *user) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.User, error) {
	query := "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users WHERE uuid = $1"
	return tGet[model.User](ctx, tx, query, uuid)
}

func (r *user) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.User, error) {
	query := "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users"
	return tGetAll[model.User](ctx, tx, query)
}

func (r *user) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			uuid VARCHAR(255) NOT NULL, 
			name VARCHAR(255) NOT NULL, 
			role INT NOT NULL, 
			phone VARCHAR(255) NOT NULL UNIQUE, 
			created_at TIMESTAMP NOT NULL, 
			updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}
