package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type OrderItemHandler interface {
	proto.OrderItemServiceServer
}

type orderItemHandler struct {
	service service.OrderItemService
	proto.UnimplementedOrderItemServiceServer
}

func NewOrderItemHandler(s service.OrderItemService) OrderItemHandler {
	return &orderItemHandler{service: s}
}
