package model

import "time"

type Auth struct {
	ID              int64      `db:"id"`
	UserID          int64      `db:"user_id"`
	Password        string     `db:"password"`
	Verified        bool       `db:"verified"`
	Disabled        bool       `db:"disabled"`
	Locked          bool       `db:"locked"`
	ActivationToken *string    `db:"activation_token"`
	ResetToken      *string    `db:"reset_token"`
	LastLogin       *time.Time `db:"last_login"`
	ResetExpiration *time.Time `db:"reset_expiration"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
}
