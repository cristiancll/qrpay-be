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
	CountByPhone(ctx context.Context, tx pgx.Tx, phone string) error
	GetUserByPhone(ctx context.Context, tx pgx.Tx, phone string) (*model.User, error)
	GetVerifiedList(ctx context.Context) ([]string, error)
}

type user struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) User {
	return &user{db: db}
}

const (
	createUserTableQuery = `CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY, 
			uuid VARCHAR(255) NOT NULL, 
			name VARCHAR(255) NOT NULL, 
			role INT NOT NULL, 
			phone VARCHAR(255) NOT NULL, 
			created_at TIMESTAMP NOT NULL, 
			updated_at TIMESTAMP NOT NULL)`
	createUserQuery          = "INSERT INTO users (uuid, name, role, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, now(), now()) RETURNING id, created_at, updated_at"
	updateUserQuery          = "UPDATE users SET name = $2, role = $3, phone = $4, updated_at = now() WHERE id = $1 RETURNING updated_at"
	deleteUserQuery          = "DELETE FROM users WHERE id = $1"
	getUserByIDQuery         = "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users WHERE id = $1"
	getUserByUUIDQuery       = "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users WHERE uuid = $1"
	getAllUsersQuery         = "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users"
	getUserByPhoneQuery      = "SELECT id, uuid, name, role, phone, created_at, updated_at FROM users WHERE phone = $1"
	countByPhoneQuery        = "SELECT count(*) FROM users WHERE phone = $1"
	getVerifiedUserListQuery = "SELECT u.phone FROM users u INNER JOIN auths a ON u.id = a.user_id WHERE a.verified = true"
)

func (r *user) GetVerifiedList(ctx context.Context) ([]string, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer tx.Rollback(ctx)

	rows, err := tx.Query(ctx, getVerifiedUserListQuery)
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

func (r *user) GetUserByPhone(ctx context.Context, tx pgx.Tx, username string) (*model.User, error) {
	user := &model.User{}
	row := tx.QueryRow(ctx, getUserByPhoneQuery, username)
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Role, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, status.Error(codes.NotFound, errors.USER_NOT_FOUND)
	} else if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return user, nil
}

func (r *user) CountByPhone(ctx context.Context, tx pgx.Tx, phone string) error {
	count := 0
	row := tx.QueryRow(ctx, countByPhoneQuery, phone)
	err := row.Scan(&count)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	if count > 0 {
		return status.Error(codes.Internal, errors.USER_ALREADY_EXISTS)
	}
	return nil
}

func (r *user) TCreate(ctx context.Context, tx pgx.Tx, user *model.User) error {
	var (
		id        int64
		createdAt time.Time
		updatedAt time.Time
	)
	user.UUID = uuid.New().String()
	row := tx.QueryRow(ctx, createUserQuery, user.UUID, user.Name, user.Role, user.Phone)
	err := row.Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}

	user.ID = id
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return nil
}

func (r *user) TUpdate(ctx context.Context, tx pgx.Tx, user *model.User) error {
	var (
		updatedAt time.Time
	)
	row := tx.QueryRow(ctx, updateUserQuery, user.ID, user.Name, user.Role, user.Phone)
	err := row.Scan(&updatedAt)
	if err == pgx.ErrNoRows {
		return status.Error(codes.NotFound, errors.USER_NOT_FOUND)
	} else if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	user.UpdatedAt = updatedAt
	return nil
}

func (r *user) TDelete(ctx context.Context, tx pgx.Tx, user *model.User) error {
	_, err := tx.Exec(ctx, deleteUserQuery, user.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *user) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.User, error) {
	user := &model.User{}

	row := tx.QueryRow(ctx, getUserByIDQuery, id)
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Role, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, status.Error(codes.NotFound, errors.USER_NOT_FOUND)
	} else if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}

	return user, nil
}

func (r *user) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.User, error) {
	user := &model.User{}
	row := tx.QueryRow(ctx, getUserByUUIDQuery, uuid)
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Role, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err == pgx.ErrNoRows {
		return nil, status.Error(codes.NotFound, errors.USER_NOT_FOUND)
	} else if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return user, nil
}

func (r *user) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.User, error) {
	var users []*model.User

	rows, err := tx.Query(ctx, getAllUsersQuery)
	if err == pgx.ErrNoRows {
		return nil, status.Error(codes.NotFound, errors.NO_USERS_FOUND)
	} else if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.UUID, &u.Name, &u.Role, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *user) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createUserTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}
