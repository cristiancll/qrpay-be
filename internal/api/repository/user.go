package repository

import (
	"context"
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserRepository interface {
	PublicRepository[model.User]
	CountByEmailOrPassword(ctx context.Context, tx pgx.Tx, email string, phone string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

const (
	createUserTableQuery        = "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, uuid VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, role INT NOT NULL, email VARCHAR(255) NOT NULL, phone VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL)"
	createUserQuery             = "INSERT INTO users (uuid, name, role, email, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, now(), now()) RETURNING id, created_at, updated_at"
	updateUserQuery             = "UPDATE users SET name = $2, role = $3, email = $4, phone = $5, updated_at = now() WHERE id = $1"
	deleteUserQuery             = "DELETE FROM users WHERE id = $1"
	getUserByIDQuery            = "SELECT id, uuid, name, role, email, phone, created_at, updated_at FROM users WHERE id = $1"
	getUserByUUIDQuery          = "SELECT id, uuid, name, role, email, phone, created_at, updated_at FROM users WHERE uuid = $1"
	getAllUsersQuery            = "SELECT id, uuid, name, role, email, phone, created_at, updated_at FROM users"
	countByEmailOrPasswordQuery = "SELECT count(*) FROM users WHERE email = $1 OR phone = $2"
)

func (r *userRepository) CountByEmailOrPassword(ctx context.Context, tx pgx.Tx, email string, phone string) error {
	var count int
	row := tx.QueryRow(ctx, countByEmailOrPasswordQuery, email, phone)
	err := row.Scan(&count)
	if err != nil {
		return fmt.Errorf("error getting user count: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("user already exists")
	}
	return nil
}

func (r *userRepository) Create(ctx context.Context, tx pgx.Tx, user *model.User) error {
	var (
		id        int64
		createdAt time.Time
		updatedAt time.Time
	)
	user.UUID = uuid.New().String()
	row := tx.QueryRow(ctx, createUserQuery, user.UUID, user.Name, user.Role, user.Email, user.Phone)
	err := row.Scan(&id, &createdAt, &updatedAt)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}

	user.ID = id
	user.CreatedAt = createdAt
	user.UpdatedAt = updatedAt
	return nil
}

func (r *userRepository) Update(ctx context.Context, tx pgx.Tx, user *model.User) error {
	var (
		updatedAt time.Time
	)
	row := tx.QueryRow(ctx, updateUserQuery, user.ID, user.Name, user.Role, user.Email, user.Phone)
	err := row.Scan(&updatedAt)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	user.UpdatedAt = updatedAt
	return nil
}

func (r *userRepository) Delete(ctx context.Context, tx pgx.Tx, user *model.User) error {
	_, err := tx.Exec(ctx, deleteUserQuery, user.ID)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (r *userRepository) GetById(ctx context.Context, tx pgx.Tx, id int64) (*model.User, error) {
	var user *model.User

	row := tx.QueryRow(ctx, getUserByIDQuery, id)
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Role, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func (r *userRepository) GetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.User, error) {
	var user *model.User

	row := tx.QueryRow(ctx, getUserByUUIDQuery, uuid)
	err := row.Scan(&user.ID, &user.UUID, &user.Name, &user.Role, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func (r *userRepository) GetAll(ctx context.Context, tx pgx.Tx) ([]model.User, error) {
	var users []model.User

	rows, err := tx.Query(ctx, getAllUsersQuery)
	if err != nil {
		return nil, fmt.Errorf("error getting users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Role, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error getting user: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createUserTableQuery)
	if err != nil {
		return fmt.Errorf("error migrating user table: %w", err)
	}
	return nil
}
