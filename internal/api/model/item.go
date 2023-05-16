package model

import "time"

type Item struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	Category  Category  `db:"category"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
