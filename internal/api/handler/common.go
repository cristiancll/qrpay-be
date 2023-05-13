package handler

import "context"

type Creater[E any, REQ any, RES any] interface {
	Create(context.Context, *REQ) (*RES, error)
}
type Getter[E any, REQ any, RES any] interface {
	Get(context.Context, *REQ) (*RES, error)
}
type Lister[E any, REQ any, RES any] interface {
	List(context.Context, *REQ) (*RES, error)
}
type Updater[E any, REQ any, RES any] interface {
	Update(context.Context, *REQ) (*RES, error)
}
type Deleter[E any, REQ any, RES any] interface {
	Delete(context.Context, *REQ) (*RES, error)
}
