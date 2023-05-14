package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type Sale interface {
	proto.SaleServiceServer
}

type sale struct {
	service service.Sale
	proto.UnimplementedSaleServiceServer
}

func NewSale(s service.Sale) Sale {
	return &sale{service: s}
}
