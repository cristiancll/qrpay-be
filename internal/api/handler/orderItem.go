package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type OrderItem interface {
	proto.OrderItemServiceServer
}

type orderItem struct {
	service service.OrderItem
	proto.UnimplementedOrderItemServiceServer
}

func NewOrderItem(s service.OrderItem) OrderItem {
	return &orderItem{service: s}
}
