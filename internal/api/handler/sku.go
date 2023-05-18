package handler

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SKU interface {
	Creater[model.SKU, proto.SKUCreateRequest, proto.SKUCreateResponse]
	Updater[model.SKU, proto.SKUUpdateRequest, proto.SKUUpdateResponse]
	Deleter[model.SKU, proto.SKUDeleteRequest, proto.SKUDeleteResponse]
	Lister[model.SKU, proto.SKUListRequest, proto.SKUListResponse]
	proto.SKUServiceServer
}

type sku struct {
	service service.SKU
	proto.UnimplementedSKUServiceServer
}

func NewSKU(s service.SKU) SKU {
	return &sku{service: s}
}

func (s *sku) Create(ctx context.Context, req *proto.SKUCreateRequest) (*proto.SKUCreateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errors.NAME_REQUIRED)
	}
	if req.ItemUUID == "" {
		return nil, status.Error(codes.InvalidArgument, errors.ITEM_REQUIRED)
	}
	if req.Price == 0 {
		return nil, status.Error(codes.InvalidArgument, errors.PRICE_REQUIRED)
	}

	sku, err := s.service.Create(ctx, req.ItemUUID, req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}

	res := &proto.SKUCreateResponse{
		Sku: &proto.SKU{
			Uuid: sku.UUID,
			Name: sku.Name,
			Item: &proto.Item{
				Uuid: sku.Item.UUID,
				Name: sku.Item.Name,
				Category: &proto.Category{
					Uuid:      sku.Item.Category.UUID,
					Name:      sku.Item.Category.Name,
					CreatedAt: timestamppb.New(sku.Item.Category.CreatedAt),
					UpdatedAt: timestamppb.New(sku.Item.Category.UpdatedAt),
				},
				CreatedAt: timestamppb.New(sku.Item.CreatedAt),
				UpdatedAt: timestamppb.New(sku.Item.UpdatedAt),
			},
			Description: sku.Description,
			Price:       sku.Price,
			CreatedAt:   timestamppb.New(sku.CreatedAt),
			UpdatedAt:   timestamppb.New(sku.UpdatedAt),
		},
	}
	return res, nil
}

func (s *sku) Update(ctx context.Context, req *proto.SKUUpdateRequest) (*proto.SKUUpdateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, errors.UUID_REQUIRED)
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errors.NAME_REQUIRED)
	}
	if req.ItemUUID == "" {
		return nil, status.Error(codes.InvalidArgument, errors.ITEM_REQUIRED)
	}
	if req.Price == 0 {
		return nil, status.Error(codes.InvalidArgument, errors.PRICE_REQUIRED)
	}
	sku, err := s.service.Update(ctx, req.Uuid, req.ItemUUID, req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	res := &proto.SKUUpdateResponse{
		Sku: &proto.SKU{
			Uuid: sku.UUID,
			Name: sku.Name,
			Item: &proto.Item{
				Uuid: sku.Item.UUID,
				Name: sku.Item.Name,
				Category: &proto.Category{
					Uuid:      sku.Item.Category.UUID,
					Name:      sku.Item.Category.Name,
					CreatedAt: timestamppb.New(sku.Item.Category.CreatedAt),
					UpdatedAt: timestamppb.New(sku.Item.Category.UpdatedAt),
				},
				CreatedAt: timestamppb.New(sku.Item.CreatedAt),
				UpdatedAt: timestamppb.New(sku.Item.UpdatedAt),
			},
			Description: sku.Description,
			Price:       sku.Price,
			CreatedAt:   timestamppb.New(sku.CreatedAt),
			UpdatedAt:   timestamppb.New(sku.UpdatedAt),
		},
	}
	return res, nil
}

func (s *sku) Delete(ctx context.Context, req *proto.SKUDeleteRequest) (*proto.SKUDeleteResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, errors.UUID_REQUIRED)
	}
	err = s.service.Delete(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &proto.SKUDeleteResponse{}, nil
}

func (s *sku) List(ctx context.Context, req *proto.SKUListRequest) (*proto.SKUListResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	skus, err := s.service.List(ctx)
	if err != nil {
		return nil, err
	}
	res := &proto.SKUListResponse{
		Skus: make([]*proto.SKU, 0),
	}
	for _, sku := range skus {
		res.Skus = append(res.Skus, &proto.SKU{
			Uuid: sku.UUID,
			Name: sku.Name,
			Item: &proto.Item{
				Uuid: sku.Item.UUID,
				Name: sku.Item.Name,
				Category: &proto.Category{
					Uuid:      sku.Item.Category.UUID,
					Name:      sku.Item.Category.Name,
					CreatedAt: timestamppb.New(sku.Item.Category.CreatedAt),
					UpdatedAt: timestamppb.New(sku.Item.Category.UpdatedAt),
				},
				CreatedAt: timestamppb.New(sku.Item.CreatedAt),
				UpdatedAt: timestamppb.New(sku.Item.UpdatedAt),
			},
			Description: sku.Description,
			Price:       sku.Price,
			CreatedAt:   timestamppb.New(sku.CreatedAt),
			UpdatedAt:   timestamppb.New(sku.UpdatedAt),
		})
	}
	return res, nil
}
