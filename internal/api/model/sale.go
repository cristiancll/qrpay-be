package model

import "time"

type Sale struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	UserId    int64     `db:"user_id"`
	SellerId  int64     `db:"seller_id"`
	Amount    int64     `db:"amount"`
	Paid      bool      `db:"paid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
