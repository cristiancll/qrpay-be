package model

import "time"

type Auth struct {
	ID         int64      `db:"id"`
	UserID     int64      `db:"user_id"`
	Password   string     `db:"password"`
	Verified   bool       `db:"verified"`
	Disabled   bool       `db:"disabled"`
	ResetToken *string    `db:"reset_token"`
	LastLogin  *time.Time `db:"last_login"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
}
