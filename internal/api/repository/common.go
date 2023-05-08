package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Creater[E any] interface {
	Create(context.Context, pgx.Tx, *E) error
}
type Updater[E any] interface {
	Update(context.Context, pgx.Tx, *E) error
}
type Deleter[E any] interface {
	Delete(context.Context, pgx.Tx, *E) error
}
type GetterById[E any] interface {
	GetById(context.Context, pgx.Tx, int64) (*E, error)
}
type GetterByUUID[E any] interface {
	GetByUUID(context.Context, pgx.Tx, string) (*E, error)
}
type GetterAll[E any] interface {
	GetAll(context.Context, pgx.Tx) ([]E, error)
}

type Migrater interface {
	Migrate(context.Context) error
}

type PublicRepository[E any] interface {
	Migrater
	Creater[E]
	Updater[E]
	Deleter[E]
	GetterById[E]
	GetterByUUID[E]
	GetterAll[E]
}

type PrivateRepository[E any] interface {
	Migrater
	Creater[E]
	Updater[E]
	Deleter[E]
	GetterById[E]
}
