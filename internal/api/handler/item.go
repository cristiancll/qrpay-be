package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type Item interface {
	proto.ItemServiceServer
}

type item struct {
	service service.Item
	proto.UnimplementedItemServiceServer
}

func NewItem(s service.Item) Item {
	return &item{service: s}
}
