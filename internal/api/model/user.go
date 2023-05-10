package model

import (
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/roles"
	"time"
)

type User struct {
	ID        int64      `db:"id"`
	UUID      string     `db:"uuid"`
	Name      string     `db:"name"`
	Role      roles.Role `db:"role"`
	Phone     string     `db:"phone"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}

func (u *User) WelcomeMessage() string {
	return fmt.Sprintf("Seja bem vindo ao QRPay, %s!", u.Name)
}

func (u *User) AccessMessage() string {
	return fmt.Sprintf("Seu código de acesso é: %s", u.UUID)
}
