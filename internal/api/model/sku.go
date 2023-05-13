package model

import "time"

type SKU struct {
	ID          int64     `db:"id"`
	UUID        string    `db:"uuid"`
	ItemId      int64     `db:"item_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Price       int64     `db:"price"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
