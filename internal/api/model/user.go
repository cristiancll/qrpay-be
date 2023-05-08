package model

import (
	"github.com/cristiancll/qrpay-be/internal/roles"
	"time"
)

type User struct {
	ID        int64      `db:"id"`
	UUID      string     `db:"uuid"`
	Name      string     `db:"name"`
	Role      roles.Role `db:"role"`
	Email     string     `db:"email"`
	Phone     string     `db:"phone"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
