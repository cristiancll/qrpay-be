package repository

import (
	"context"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r *saleItem) Migrate(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS sale_items(
    		id SERIAL PRIMARY KEY,
    		uuid VARCHAR(255) NOT NULL,
    		sale_id INT NOT NULL REFERENCES sales(id),
    		sku_id INT NOT NULL REFERENCES skus(id),
    		created_at TIMESTAMP NOT NULL,
    		updated_at TIMESTAMP NOT NULL)`
	return migrate(ctx, r.db, query)
}

func (r *saleItem) TCreate(ctx context.Context, tx pgx.Tx, item *model.SaleItem) error {
	item.UUID = uuid.New().String()
	item.CreatedAt = time.Now().UTC()
	item.UpdatedAt = time.Now().UTC()
	query := `INSERT INTO sale_items (uuid, sale_id, sku_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	id, err := tCreate(ctx, tx, query, item.UUID, item.Sale.ID, item.SKU.ID, item.CreatedAt, item.UpdatedAt)
	if err != nil {
		return errs.Wrap(err, "")
	}
	item.ID = id
	return nil
}

func (r *saleItem) TGetAllByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.SaleItem, error) {
	query := `SELECT si.id, si.uuid, sk.uuid, i.uuid, si.created_at, si.updated_at FROM sale_items AS si
									JOIN sales AS s ON si.sale_id = s.id
                                    JOIN items AS i ON si.sku_id = i.id
									JOIN skus AS sk ON si.sku_id = sk.id                                 
								  WHERE s.user_id = $1`
	return tGetAll[model.SaleItem](ctx, tx, query, user.ID)
}

func (r *saleItem) TGetAllAvailableByUser(ctx context.Context, tx pgx.Tx, user *model.User) ([]*model.SaleItem, error) {
	query := `SELECT si.id, si.uuid, sk.uuid, i.uuid, si.created_at, si.updated_at FROM sale_items AS si
     									JOIN sales AS s ON s.id = si.sale_id
     									JOIN skus AS sk ON sk.id = si.sku_id
     									JOIN items AS i ON i.id = sk.item_id
								   LEFT JOIN retrievals r on r.sale_item_id = si.id
								   WHERE s.user_id = $1 AND r.id IS NULL`
	return tGetAll[model.SaleItem](ctx, tx, query, user.ID)
}

func (r *saleItem) TGetAllByUUIDs(ctx context.Context, tx pgx.Tx, saleItemUUIDs []string) ([]*model.SaleItem, error) {
	query := `SELECT si.id, si.uuid, si.sale_id, s.id, s.uuid, s.name, si.created_at, si.updated_at FROM sale_items AS si
								   JOIN skus s on s.id = si.sku_id
								   WHERE si.uuid = ANY($1)`

	return tGetAll[model.SaleItem](ctx, tx, query, saleItemUUIDs)
}
