package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type SKU interface {
	proto.SKUServiceServer
}

type sku struct {
	service service.SKU
	proto.UnimplementedSKUServiceServer
}

func NewSKU(s service.SKU) SKU {
	return &sku{service: s}
}
