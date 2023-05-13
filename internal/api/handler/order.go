package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type OrderHandler interface {
	proto.OrderServiceServer
}

type orderHandler struct {
	service service.OrderService
	proto.UnimplementedOrderServiceServer
}

func NewOrderHandler(s service.OrderService) OrderHandler {
	return &orderHandler{service: s}
}
