package model

import (
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/common"
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

Para confirmar o seu cadastro, responda *CONFIRMA* para este número.
`
	return fmt.Sprintf(format, u.Name)
}

func (u *User) ConfirmMessage() string {
	format := `Olá *%s*, o seu cadastro foi confirmado com sucesso!

Nosso sistema é super fácil e simples de usar. Basta seguir os passos abaixo:

1. Apresente o seu QR Code ao caixa do evento.

2. O caixa irá escanear o seu QR Code e montar o seu pedido.

3. Confirme o valor e faça o pagamento.

4. Após o pagamento, o caixa irá confirmar e finalizar o seu pedido.

5. Pronto! Agora é só se dirigir ao balcão de retirada e apresentar o seu QR Code para pegar o seu pedido.

Se você escolheu itens avulsos, fique à vontade para retirá-los a qualquer momento.
    
No caso de um combo, lembre-se de que será necessário retirá-lo de uma vez, sem fracionamentos.

E lembre-se: seu QR Code é recarregável! Basta apresentá-lo ao caixa novamente em futuras visitas!
`
	return fmt.Sprintf(format, u.Name)
}

func (u *User) BalanceMessage() string {
	format := `O seu saldo de pedidos: `
	return format
}

func (u *User) NewSale(sale *Sale, items []*SaleItem) string {
	format := `Recebemos o seu pedido no valor de *%s*.

Agora é só se dirigir ao balcão de retirada e apresentar o seu QR Code para pegar o seu pedido!`

	return fmt.Sprintf(format, common.FormatCentsToBRL(sale.Total))
}
