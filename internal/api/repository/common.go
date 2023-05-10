package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type TCreater[E any] interface {
	TCreate(context.Context, pgx.Tx, *E) error
}
type TUpdater[E any] interface {
	TUpdate(context.Context, pgx.Tx, *E) error
}
type TDeleter[E any] interface {
	TDelete(context.Context, pgx.Tx, *E) error
}
type TGetterById[E any] interface {
	TGetById(context.Context, pgx.Tx, int64) (*E, error)
}
type TGetterByUUID[E any] interface {
	TGetByUUID(context.Context, pgx.Tx, string) (*E, error)
}
type TGetterAll[E any] interface {
	TGetAll(context.Context, pgx.Tx) ([]*E, error)
}

type Creater[E any] interface {
	Create(context.Context, *E) error
}
type Updater[E any] interface {
	Update(context.Context, *E) error
}
type Deleter[E any] interface {
	Delete(context.Context, *E) error
}
type GetterById[E any] interface {
	GetById(context.Context, int64) (*E, error)
}
type GetterByUUID[E any] interface {
	GetByUUID(context.Context, string) (*E, error)
}
type GetterAll[E any] interface {
	GetAll(context.Context) ([]*E, error)
}

type Migrater interface {
	Migrate(context.Context) error
}

type TCRUDer[E any] interface {
	TCreater[E]
	TUpdater[E]
	TDeleter[E]
	TGetterById[E]
	TGetterByUUID[E]
	TGetterAll[E]
}

type CRUDer[E any] interface {
	Creater[E]
	Updater[E]
	Deleter[E]
	GetterById[E]
	GetterByUUID[E]
	GetterAll[E]
}
