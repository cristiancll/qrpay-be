package model

import "time"

type Item struct {
	ID         int64     `db:"id"`
	UUID       string    `db:"uuid"`
	CategoryID int64     `db:"category_id"`
	Name       string    `db:"name"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
