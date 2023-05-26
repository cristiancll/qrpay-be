package model

import "time"

type Category struct {
	ID        int64     `db:"id"`
	UUID      string    `db:"uuid"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
