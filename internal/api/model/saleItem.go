package model

import "time"

type SaleItem struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	SaleID    int64     `db:"sale_id"`
	SKUID     int64     `db:"sku_id"`
	Quantity  int64     `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
