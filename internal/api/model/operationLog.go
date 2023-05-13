package model

import "time"

type OperationLog struct {
	ID          int64     `db:"id"`
	UUID        string    `db:"uuid"`
	UserId      int64     `db:"user_id"`
	SellerId    int64     `db:"seller_id"`
	Operation   string    `db:"operation"`
	OperationId int64     `db:"operation_id"`
	Metadata    string    `db:"metadata"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
