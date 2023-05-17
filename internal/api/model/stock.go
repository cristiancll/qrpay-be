package model

import "time"

type Stock struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	SKU       SKU       `db:"sku"`
	Quantity  int64     `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
