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

type Item interface {
	Creater[model.Item, proto.ItemCreateRequest, proto.ItemCreateResponse]
	Updater[model.Item, proto.ItemUpdateRequest, proto.ItemUpdateResponse]
	Deleter[model.Item, proto.ItemDeleteRequest, proto.ItemDeleteResponse]
	Lister[model.Item, proto.ItemListRequest, proto.ItemListResponse]
	proto.ItemServiceServer
}

type item struct {
	service service.Item
	proto.UnimplementedItemServiceServer
}

func NewItem(s service.Item) Item {
	return &item{service: s}
}

func (i item) Create(ctx context.Context, req *proto.ItemCreateRequest) (*proto.ItemCreateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errors.NAME_REQUIRED)
	}
	if req.CategoryUUID == "" {
		return nil, status.Error(codes.InvalidArgument, errors.CATEGORY_REQUIRED)
	}
	item, err := i.service.Create(ctx, req.Name, req.CategoryUUID)
	if err != nil {
		return nil, err
	}
	res := &proto.ItemCreateResponse{
		Item: &proto.Item{
			Uuid: item.UUID,
			Name: item.Name,
			Category: &proto.Category{
				Uuid: item.Category.UUID,
				Name: item.Category.Name,
			},
			CreatedAt: timestamppb.New(item.CreatedAt),
			UpdatedAt: timestamppb.New(item.UpdatedAt),
		},
	}
	return res, nil
}

func (i item) Update(ctx context.Context, req *proto.ItemUpdateRequest) (*proto.ItemUpdateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, errors.NAME_REQUIRED)
	}
	if req.CategoryUUID == "" {
		return nil, status.Error(codes.InvalidArgument, errors.CATEGORY_REQUIRED)
	}
	item, err := i.service.Update(ctx, req.Uuid, req.Name, req.CategoryUUID)
	if err != nil {
		return nil, err
	}
	res := &proto.ItemUpdateResponse{
		Item: &proto.Item{
			Uuid: item.UUID,
			Name: item.Name,
			Category: &proto.Category{
				Uuid: item.Category.UUID,
				Name: item.Category.Name,
			},
			CreatedAt: timestamppb.New(item.CreatedAt),
			UpdatedAt: timestamppb.New(item.UpdatedAt),
		},
	}
	return res, nil
}

func (i item) Delete(ctx context.Context, req *proto.ItemDeleteRequest) (*proto.ItemDeleteResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	if req.Uuid == "" {
		return nil, status.Error(codes.InvalidArgument, errors.UUID_REQUIRED)
	}
	err = i.service.Delete(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	res := &proto.ItemDeleteResponse{}
	return res, nil

}

func (i item) List(ctx context.Context, req *proto.ItemListRequest) (*proto.ItemListResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, err
	}
	items, err := i.service.List(ctx)
	if err != nil {
		return nil, err
	}
	res := &proto.ItemListResponse{
		Items: make([]*proto.Item, 0),
	}
	for _, item := range items {
		res.Items = append(res.Items, &proto.Item{
			Uuid: item.UUID,
			Name: item.Name,
			Category: &proto.Category{
				Uuid: item.Category.UUID,
				Name: item.Category.Name,
			},
			CreatedAt: timestamppb.New(item.CreatedAt),
			UpdatedAt: timestamppb.New(item.UpdatedAt),
		})
	}
	return res, nil

}
