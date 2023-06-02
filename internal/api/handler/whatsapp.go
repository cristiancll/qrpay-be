package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type WhatsApp interface {
	proto.WhatsAppServiceServer
}

type whatsApp struct {
	service service.WhatsApp
	proto.UnimplementedWhatsAppServiceServer
}

func NewWhatsApp(s service.WhatsApp) WhatsApp {
	return &whatsApp{service: s}
}
