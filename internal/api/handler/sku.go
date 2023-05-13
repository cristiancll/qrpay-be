package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type SKUHandler interface {
	proto.SKUServiceServer
}

type skuHandler struct {
	service service.SKUService
	proto.UnimplementedSKUServiceServer
}

func NewSKUHandler(s service.SKUService) SKUHandler {
	return &skuHandler{service: s}
}
