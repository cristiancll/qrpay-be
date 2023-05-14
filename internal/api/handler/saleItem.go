package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type SaleItem interface {
	proto.SaleItemServiceServer
}

type saleItem struct {
	service service.SaleItem
	proto.UnimplementedSaleItemServiceServer
}

func NewSaleItem(s service.SaleItem) SaleItem {
	return &saleItem{service: s}
}
