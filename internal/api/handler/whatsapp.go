package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type WhatsAppHandler interface {
	Get(ctx context.Context, req *proto.WhatsAppGetRequest) (*proto.WhatsAppGetResponse, error)
	List(ctx context.Context, req *proto.VoidRequest) (*proto.WhatsAppListResponse, error)
	proto.WhatsAppServiceServer
}

type whatsAppHandler struct {
	service service.WhatsAppService
	proto.UnimplementedWhatsAppServiceServer
}

func NewWhatsAppHandler(s service.WhatsAppService) WhatsAppHandler {
	return &whatsAppHandler{service: s}
}
