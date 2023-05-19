package model

import "time"

type Retrieval struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	User      User      `db:"user"`
	Seller    User      `db:"seller"`
	SaleItem  SaleItem  `db:"sale_item"`
	Delivered bool      `db:"delivered"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
