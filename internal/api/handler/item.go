package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type ItemHandler interface {
	proto.ItemServiceServer
}

type itemHandler struct {
	service service.ItemService
	proto.UnimplementedItemServiceServer
}

func NewItemHandler(s service.ItemService) ItemHandler {
	return &itemHandler{service: s}
}
