package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type StockHandler interface {
	proto.StockServiceServer
}

type stockHandler struct {
	service service.StockService
	proto.UnimplementedStockServiceServer
}

func NewStockHandler(s service.StockService) StockHandler {
	return &stockHandler{service: s}
}
