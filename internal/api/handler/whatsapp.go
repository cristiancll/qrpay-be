package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WhatsApp interface {
	Get(ctx context.Context, req *proto.WhatsAppGetRequest) (*proto.WhatsAppGetResponse, error)
	List(ctx context.Context, req *proto.VoidRequest) (*proto.WhatsAppListResponse, error)
	proto.WhatsAppServiceServer
}

type whatsApp struct {
	service service.WhatsApp
	proto.UnimplementedWhatsAppServiceServer
}

func NewWhatsApp(s service.WhatsApp) WhatsApp {
	return &whatsApp{service: s}
}

func (h *whatsApp) Get(ctx context.Context, req *proto.WhatsAppGetRequest) (*proto.WhatsAppGetResponse, error) {
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, errors.UUID_REQUIRED)
	}
	whats, err := h.service.Get(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	res := &proto.WhatsAppGetResponse{
		WhatsApp: &proto.WhatsApp{
			Uuid:      whats.UUID,
			Qr:        whats.QR,
			Phone:     whats.Phone,
			Active:    whats.Active,
			Banned:    whats.Banned,
			CreatedAt: timestamppb.New(whats.CreatedAt),
			UpdatedAt: timestamppb.New(whats.UpdatedAt),
		},
	}
	return res, nil
}

func (h *whatsApp) List(ctx context.Context, _ *proto.VoidRequest) (*proto.WhatsAppListResponse, error) {
	whatsList, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	pWhatsList := make([]*proto.WhatsApp, 0)
	for _, whats := range whatsList {
		pWhatsList = append(pWhatsList, &proto.WhatsApp{
			Uuid:      whats.UUID,
			Qr:        whats.QR,
			Phone:     whats.Phone,
			Active:    whats.Active,
			Banned:    whats.Banned,
			CreatedAt: timestamppb.New(whats.CreatedAt),
			UpdatedAt: timestamppb.New(whats.UpdatedAt),
		})
	}
	res := &proto.WhatsAppListResponse{
		WhatsAppList: pWhatsList,
	}
	return res, nil
}
