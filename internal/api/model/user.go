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
	format := `Olá *%s*, Agora você faz parte do sistema *QRPay*!

O QR Code acima é o seu código de acesso ao sistema. Guarde-o com cuidado, pois ele é único e intransferível.

Apresente o seu QR Code para o operador do caixa para carregar o seu saldo.

Apresente o seu QR Code para o balconista para retirar os seus pedidos.

Para confirmar o seu cadastro, responda *CONFIRMA* para este número.
`
	return fmt.Sprintf(format, u.Name)
}

func (u *User) ConfirmMessage() string {
	format := `Olá *%s*, o seu cadastro foi confirmado com sucesso!`
	return fmt.Sprintf(format, u.Name)
}

func (u *User) SaleMessage() string {
	format := `Voce fez um pedido`
	return format
}

func (u *User) BalanceMessage() string {
	format := `O seu saldo de pedidos: `
	return format
}
