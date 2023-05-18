package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Sale interface {
	Creater[proto.Sale, proto.SaleCreateRequest, proto.SaleCreateResponse]
	proto.SaleServiceServer
}

type sale struct {
	service service.Sale
	proto.UnimplementedSaleServiceServer
}

func NewSale(s service.Sale) Sale {
	return &sale{service: s}
}

func (s *sale) Create(ctx context.Context, req *proto.SaleCreateRequest) (*proto.SaleCreateResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	sellerUUID, err := getUUIDFromContext(ctx)
	if err != nil {
		return nil, err
	}
	err = checkValidUUID(req.UserUUID)
	if err != nil {
		return nil, err
	}
	saleUnits := req.GetUnits()
	if len(req.GetUnits()) == 0 {
		return nil, err
	}
	saleItemsMap := make(map[string]int64, len(saleUnits))
	for _, unit := range saleUnits {
		err = checkValidUUID(unit.SkuUUID)
		if err != nil {
			return nil, err
		}
		if unit.Quantity == 0 {
			return nil, err
		}
		saleItemsMap[unit.SkuUUID] = unit.Quantity
	}

	sale, err := s.service.Create(ctx, req.UserUUID, sellerUUID, saleItemsMap)
	if err != nil {
		return nil, err
	}

	res := &proto.SaleCreateResponse{
		Sale: &proto.Sale{
			Uuid: sale.UUID,
			User: &proto.User{
				Uuid: sale.User.UUID,
			},
			Seller: &proto.User{
				Uuid: sale.Seller.UUID,
			},
			Total:     sale.Total,
			Paid:      sale.Paid,
			CreatedAt: timestamppb.New(sale.CreatedAt),
			UpdatedAt: timestamppb.New(sale.UpdatedAt),
		},
	}
	return res, nil
}
