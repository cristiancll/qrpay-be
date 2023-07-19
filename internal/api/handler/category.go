package handler

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Category interface {
	Creater[model.Category, proto.CategoryCreateRequest, proto.CategoryCreateResponse]
	Updater[model.Category, proto.CategoryUpdateRequest, proto.CategoryUpdateResponse]
	Deleter[model.Category, proto.CategoryDeleteRequest, proto.CategoryDeleteResponse]
	Lister[model.Category, proto.CategoryListRequest, proto.CategoryListResponse]
	proto.CategoryServiceServer
}

type category struct {
	service service.Category
	proto.UnimplementedCategoryServiceServer
}

func NewCategory(s service.Category) Category {
	return &category{service: s}
}

func (c *category) Create(ctx context.Context, req *proto.CategoryCreateRequest) (*proto.CategoryCreateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}
	if req.Name == "" {
		return nil, errs.New(errors.New(errMsg.NameRequired), errCode.InvalidArgument)
	}
	category, err := c.service.Create(ctx, req.Name)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedCreateCategory, req.Name)
	}
	res := &proto.CategoryCreateResponse{
		Category: &proto.Category{
			Uuid:      category.UUID,
			Name:      category.Name,
			CreatedAt: timestamppb.New(category.CreatedAt),
			UpdatedAt: timestamppb.New(category.UpdatedAt),
		},
	}
	return res, nil
}

func (c *category) Update(ctx context.Context, req *proto.CategoryUpdateRequest) (*proto.CategoryUpdateResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}
	if req.Uuid == "" {
		return nil, errs.New(errors.New(errMsg.UUIDRequired), errCode.InvalidArgument)
	}
	if req.Name == "" {
		return nil, errs.New(errors.New(errMsg.NameRequired), errCode.InvalidArgument)
	}
	category, err := c.service.Update(ctx, req.Uuid, req.Name)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedUpdateCategory, req.Uuid, req.Name)
	}
	res := &proto.CategoryUpdateResponse{
		Category: &proto.Category{
			Uuid:      category.UUID,
			Name:      category.Name,
			CreatedAt: timestamppb.New(category.CreatedAt),
			UpdatedAt: timestamppb.New(category.UpdatedAt),
		},
	}
	return res, nil
}

func (c *category) Delete(ctx context.Context, req *proto.CategoryDeleteRequest) (*proto.CategoryDeleteResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}
	if req.Uuid == "" {
		return nil, errs.New(errors.New(errMsg.UUIDRequired), errCode.InvalidArgument)
	}
	err = c.service.Delete(ctx, req.Uuid)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedDeleteCategory, req.Uuid)
	}
	res := &proto.CategoryDeleteResponse{}
	return res, nil
}

func (c *category) List(ctx context.Context, req *proto.CategoryListRequest) (*proto.CategoryListResponse, error) {
	err := checkAdminAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedAuthCheck)
	}
	categories, err := c.service.List(ctx)
	if err != nil {
		return nil, errs.Wrap(err, errMsg.FailedGetAllCategory)
	}
	res := &proto.CategoryListResponse{}
	for _, category := range categories {
		res.Categories = append(res.Categories, &proto.Category{
			Uuid:      category.UUID,
			Name:      category.Name,
			CreatedAt: timestamppb.New(category.CreatedAt),
			UpdatedAt: timestamppb.New(category.UpdatedAt),
		})
	}
	return res, nil
}
