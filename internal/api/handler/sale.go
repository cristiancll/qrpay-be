package handler

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Sale interface {
	Creater[proto.Sale, proto.SaleCreateRequest, proto.SaleCreateResponse]
	ListSaleItemsByUser(context.Context, *proto.ListSaleItemsByUserRequest) (*proto.ListSaleItemsByUserResponse, error)
	ListAvailableSaleItemsByUser(context.Context, *proto.ListAvailableSaleItemsByUserRequest) (*proto.ListAvailableSaleItemsByUserResponse, error)
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
		return nil, errs.Wrap(err, "")
	}
	sellerUUID, err := getUUIDFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	err = checkValidUUID(req.UserUUID)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	saleUnits := req.GetUnits()
	if len(req.GetUnits()) == 0 {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}
	saleItemsMap := make(map[string]int64, len(saleUnits))
	for _, unit := range saleUnits {
		err = checkValidUUID(unit.SkuUUID)
		if err != nil {
			return nil, errs.New(errors.New(""), errCode.InvalidArgument)
		}
		if unit.Quantity == 0 {
			return nil, errs.New(errors.New(""), errCode.InvalidArgument)
		}
		saleItemsMap[unit.SkuUUID] = unit.Quantity
	}

	sale, err := s.service.Create(ctx, req.UserUUID, sellerUUID, saleItemsMap)
	if err != nil {
		return nil, errs.Wrap(err, "")
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

func (s *sale) ListSaleItemsByUser(ctx context.Context, req *proto.ListSaleItemsByUserRequest) (*proto.ListSaleItemsByUserResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	err = checkValidUUID(req.UserUUID)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	saleItems, err := s.service.ListSaleItemsByUser(ctx, req.UserUUID)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	res := &proto.ListSaleItemsByUserResponse{
		SaleItems: make([]*proto.SaleItem, len(saleItems)),
	}
	for i, saleItem := range saleItems {
		res.SaleItems[i] = &proto.SaleItem{
			Uuid: saleItem.UUID,
			Sku: &proto.SKU{
				Uuid: saleItem.SKU.UUID,
				Item: &proto.Item{
					Uuid: saleItem.SKU.Item.UUID,
				},
			},
			CreatedAt: timestamppb.New(saleItem.CreatedAt),
			UpdatedAt: timestamppb.New(saleItem.UpdatedAt),
		}
	}

	return res, nil
}

func (s *sale) ListAvailableSaleItemsByUser(ctx context.Context, req *proto.ListAvailableSaleItemsByUserRequest) (*proto.ListAvailableSaleItemsByUserResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	err = checkValidUUID(req.UserUUID)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	saleItems, err := s.service.ListAvailableSaleItemsByUser(ctx, req.UserUUID)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	res := &proto.ListAvailableSaleItemsByUserResponse{
		SaleItems: make([]*proto.SaleItem, len(saleItems)),
	}
	for i, saleItem := range saleItems {
		res.SaleItems[i] = &proto.SaleItem{
			Uuid: saleItem.UUID,
			Sku: &proto.SKU{
				Uuid: saleItem.SKU.UUID,
				Item: &proto.Item{
					Uuid: saleItem.SKU.Item.UUID,
				},
			},
			CreatedAt: timestamppb.New(saleItem.CreatedAt),
			UpdatedAt: timestamppb.New(saleItem.UpdatedAt),
		}
	}

	return res, nil
}
