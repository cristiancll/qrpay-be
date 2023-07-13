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

type Stock interface {
	Creater[service.Stock, proto.StockCreateRequest, proto.StockCreateResponse]
	Updater[service.Stock, proto.StockUpdateRequest, proto.StockUpdateResponse]
	Deleter[service.Stock, proto.StockDeleteRequest, proto.StockDeleteResponse]
	Lister[service.Stock, proto.StockListRequest, proto.StockListResponse]
	proto.StockServiceServer
}

type stock struct {
	service service.Stock
	proto.UnimplementedStockServiceServer
}

func NewStock(s service.Stock) Stock {
	return &stock{service: s}
}

func (s stock) Create(ctx context.Context, req *proto.StockCreateRequest) (*proto.StockCreateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if req.SkuUUID == "" {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}
	if req.Quantity == 0 {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}

	stock, err := s.service.Create(ctx, req.SkuUUID, req.Quantity)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	res := &proto.StockCreateResponse{
		Stock: &proto.Stock{
			Uuid: stock.UUID,
			Sku: &proto.SKU{
				Uuid:      stock.SKU.UUID,
				Name:      stock.SKU.Name,
				CreatedAt: timestamppb.New(stock.SKU.CreatedAt),
				UpdatedAt: timestamppb.New(stock.SKU.UpdatedAt),
			},
			Quantity:  stock.Quantity,
			CreatedAt: timestamppb.New(stock.CreatedAt),
			UpdatedAt: timestamppb.New(stock.UpdatedAt),
		},
	}
	return res, nil
}

func (s stock) Update(ctx context.Context, req *proto.StockUpdateRequest) (*proto.StockUpdateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if req.Uuid == "" {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}
	if req.Quantity == 0 {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}
	stock, err := s.service.Update(ctx, req.Uuid, req.Quantity)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	res := &proto.StockUpdateResponse{
		Stock: &proto.Stock{
			Uuid: stock.UUID,
			Sku: &proto.SKU{
				Uuid:      stock.SKU.UUID,
				Name:      stock.SKU.Name,
				CreatedAt: timestamppb.New(stock.SKU.CreatedAt),
				UpdatedAt: timestamppb.New(stock.SKU.UpdatedAt),
			},
			Quantity:  stock.Quantity,
			CreatedAt: timestamppb.New(stock.CreatedAt),
			UpdatedAt: timestamppb.New(stock.UpdatedAt),
		},
	}
	return res, nil
}

func (s *stock) Delete(ctx context.Context, req *proto.StockDeleteRequest) (*proto.StockDeleteResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if req.Uuid == "" {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}
	err = s.service.Delete(ctx, req.Uuid)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	return &proto.StockDeleteResponse{}, nil

}

func (s stock) List(ctx context.Context, req *proto.StockListRequest) (*proto.StockListResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	stocks, err := s.service.List(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	res := &proto.StockListResponse{
		Stocks: make([]*proto.Stock, 0),
	}
	for _, stock := range stocks {
		res.Stocks = append(res.Stocks, &proto.Stock{
			Uuid: stock.UUID,
			Sku: &proto.SKU{
				Uuid:      stock.SKU.UUID,
				Name:      stock.SKU.Name,
				CreatedAt: timestamppb.New(stock.SKU.CreatedAt),
				UpdatedAt: timestamppb.New(stock.SKU.UpdatedAt),
			},
			Quantity:  stock.Quantity,
			CreatedAt: timestamppb.New(stock.CreatedAt),
			UpdatedAt: timestamppb.New(stock.UpdatedAt),
		})
	}
	return res, nil
}
