package model

import "time"

type OrderItem struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	OrderID   int64     `db:"order_id"`
	SKUID     int64     `db:"sku_id"`
	Quantity  int64     `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
