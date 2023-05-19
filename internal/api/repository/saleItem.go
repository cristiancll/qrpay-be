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

type SaleItem interface {
	Migrater
	TCreater[model.SaleItem]
	TGetAllByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.SaleItem, error)
	TGetAllAvailableByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.SaleItem, error)
	TGetAllByUUIDs(ctx context.Context, tx pgx.Tx, saleItemUUIDs []string) ([]*model.SaleItem, error)
}

type saleItem struct {
	db *pgxpool.Pool
}

func NewSaleItem(db *pgxpool.Pool) SaleItem {
	return &saleItem{db: db}
}

const (
	createSaleItemTableQuery = `CREATE TABLE IF NOT EXISTS sale_items(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		sale_id INT NOT NULL REFERENCES sales(id),
    		sku_id INT NOT NULL REFERENCES skus(id),
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	createSaleItemQuery = `INSERT INTO sale_items (uuid, sale_id, sku_id, created_at, updated_at) 
										VALUES ($1, $2, $3, $4, $5) RETURNING id`

	getAllByUserQuery = `SELECT si.id, si.uuid, sk.uuid, i.uuid, si.created_at, si.updated_at FROM sale_items AS si
									JOIN sales AS s ON si.sale_id = s.id
                                    JOIN items AS i ON si.sku_id = i.id
									JOIN skus AS sk ON si.sku_id = sk.id                                 
								  WHERE s.user_id = $1`

	getAllAvailableByUserQuery = `SELECT si.id, si.uuid, sk.uuid, i.uuid, si.created_at, si.updated_at FROM sale_items AS si
     									JOIN sales AS s ON s.id = si.sale_id
     									JOIN skus AS sk ON sk.id = si.sku_id
     									JOIN items AS i ON i.id = sk.item_id
								   LEFT JOIN retrievals r on r.sale_item_id = si.id
								   WHERE s.user_id = $1 AND r.id IS NULL`
	getAllSaleItemsByUUIDs = `SELECT si.id, si.uuid, si.sale_id, s.id, s.uuid, s.name, si.created_at, si.updated_at FROM sale_items AS si
								   JOIN skus s on s.id = si.sku_id
								   WHERE si.uuid = ANY($1)`
)

func (r *saleItem) Migrate(ctx context.Context) error {
	_, err := r.db.Exec(ctx, createSaleItemTableQuery)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *saleItem) TCreate(ctx context.Context, tx pgx.Tx, item *model.SaleItem) error {
	item.UUID = uuid.New().String()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	err := tx.QueryRow(ctx, createSaleItemQuery, item.UUID, item.Sale.ID, item.SKU.ID, item.CreatedAt, item.UpdatedAt).Scan(&item.ID)
	if err != nil {
		return status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	return nil
}

func (r *saleItem) TGetAllByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.SaleItem, error) {
	rows, err := tx.Query(ctx, getAllByUserQuery, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	return getSaleItemsFromRows(rows)
}

func (r *saleItem) TGetAllAvailableByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.SaleItem, error) {
	rows, err := tx.Query(ctx, getAllAvailableByUserQuery, user.ID)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	return getSaleItemsFromRows(rows)
}

func (r *saleItem) TGetAllByUUIDs(ctx context.Context, tx pgx.Tx, saleItemUUIDs []string) ([]*model.SaleItem, error) {
	rows, err := tx.Query(ctx, getAllSaleItemsByUUIDs, saleItemUUIDs)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
	}
	defer rows.Close()
	var saleItems []*model.SaleItem
	for rows.Next() {
		var saleItem model.SaleItem
		err := rows.Scan(&saleItem.ID, &saleItem.UUID, &saleItem.Sale.ID, &saleItem.SKU.ID, &saleItem.SKU.UUID, &saleItem.SKU.Name, &saleItem.CreatedAt, &saleItem.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		saleItems = append(saleItems, &saleItem)
	}
	return saleItems, nil
}

func getSaleItemsFromRows(rows pgx.Rows) ([]*model.SaleItem, error) {
	var saleItems []*model.SaleItem
	for rows.Next() {
		var saleItem model.SaleItem
		err := rows.Scan(&saleItem.ID, &saleItem.UUID, &saleItem.SKU.UUID, &saleItem.SKU.Item.UUID, &saleItem.CreatedAt, &saleItem.UpdatedAt)
		if err != nil {
			return nil, status.Error(codes.Internal, errors.DATABASE_ERROR)
		}
		saleItems = append(saleItems, &saleItem)
	}
	return saleItems, nil
}
