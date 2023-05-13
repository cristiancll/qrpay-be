package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type CategoryHandler interface {
	proto.CategoryServiceServer
}

type categoryHandler struct {
	service service.CategoryService
	proto.UnimplementedCategoryServiceServer
}

func NewCategoryHandler(s service.CategoryService) CategoryHandler {
	return &categoryHandler{service: s}
}
