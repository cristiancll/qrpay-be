package model

import "time"

type Sale struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	User      User      `db:"user"`
	Seller    User      `db:"seller"`
	Total     int64     `db:"total"`
	Paid      bool      `db:"paid"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
