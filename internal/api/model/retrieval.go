package model

import "time"

type Retrieval struct {
	ID         int64     `db:"id"`
	UUID       string    `db:"uuid"`
	UserId     int64     `db:"user_id"`
	SellerId   int64     `db:"seller_id"`
	SaleItemID int64     `db:"sale_item_id"`
	Quantity   int64     `db:"quantity"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
