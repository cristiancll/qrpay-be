package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type Category interface {
	proto.CategoryServiceServer
}

type category struct {
	service service.Category
	proto.UnimplementedCategoryServiceServer
}

func NewCategory(s service.Category) Category {
	return &category{service: s}
}
