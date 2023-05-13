package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type Stock interface {
	proto.StockServiceServer
}

type stock struct {
	service service.Stock
	proto.UnimplementedStockServiceServer
}

func NewStock(s service.Stock) Stock {
	return &stock{service: s}
}
