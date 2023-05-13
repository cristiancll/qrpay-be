package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type Order interface {
	proto.OrderServiceServer
}

type order struct {
	service service.Order
	proto.UnimplementedOrderServiceServer
}

func NewOrder(s service.Order) Order {
	return &order{service: s}
}
