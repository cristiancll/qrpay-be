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

type Category interface {
	Migrater
	TCRUDer[model.Category]
}

type category struct {
	db *pgxpool.Pool
}

func NewCategory(db *pgxpool.Pool) Category {
	return &category{db: db}
}

const (
	createCategoryTable = `CREATE TABLE IF NOT EXISTS categories (
								id SERIAL PRIMARY KEY, 
								uuid VARCHAR(255) NOT NULL, 
								name VARCHAR(255) NOT NULL, 
								created_at TIMESTAMP NOT NULL, 
								updated_at TIMESTAMP NOT NULL);`
	createCategoryQuery    = "INSERT INTO categories (uuid, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"
	updateCategoryQuery    = "UPDATE categories SET name = $2, updated_at = $3 WHERE id = $1 RETURNING updated_at"
	deleteCategoryQuery    = "DELETE FROM categories WHERE id = $1"
	getCategoryByIDQuery   = "SELECT id, uuid, name, created_at, updated_at FROM categories WHERE id = $1"
	getCategoryByUUIDQuery = "SELECT id, uuid, name, created_at, updated_at FROM categories WHERE uuid = $1"
	getAllCategoriesQuery  = "SELECT id, uuid, name, created_at, updated_at FROM categories"
)

func (r *category) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createCategoryTable)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *category) TCreate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	category.UUID = uuid.New().String()
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createCategoryQuery, category.UUID, category.Name, category.CreatedAt, category.UpdatedAt).Scan(&category.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *category) TUpdate(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	category.UpdatedAt = time.Now().UTC()
	_, err := tx.Exec(ctx, updateCategoryQuery, category.ID, category.Name, category.UpdatedAt)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil

}

func (r *category) TDelete(ctx context.Context, tx pgx.Tx, category *model.Category) error {
	_, err := tx.Exec(ctx, deleteCategoryQuery, category.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *category) TGetById(ctx context.Context, tx pgx.Tx, id int64) (*model.Category, error) {
	category := &model.Category{}
	err := tx.QueryRow(ctx, getCategoryByIDQuery, id).Scan(&category.ID, &category.UUID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return category, nil
}

func (r *category) TGetByUUID(ctx context.Context, tx pgx.Tx, uuid string) (*model.Category, error) {
	category := &model.Category{}
	err := tx.QueryRow(ctx, getCategoryByUUIDQuery, uuid).Scan(&category.ID, &category.UUID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return category, nil
}

func (r *category) TGetAll(ctx context.Context, tx pgx.Tx) ([]*model.Category, error) {
	rows, err := tx.Query(ctx, getAllCategoriesQuery)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()

	categories := make([]*model.Category, 0)
	for rows.Next() {
		category := &model.Category{}
		err := rows.Scan(&category.ID, &category.UUID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		categories = append(categories, category)
	}
	return categories, nil

}
