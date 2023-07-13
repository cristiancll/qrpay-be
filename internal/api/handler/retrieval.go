package handler

import (
	"context"
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	proto "github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Retrieval interface {
	Creater[model.Retrieval, proto.RetrievalCreateRequest, proto.RetrievalCreateResponse]
	Updater[model.Retrieval, proto.RetrievalUpdateRequest, proto.RetrievalUpdateResponse]
	Deleter[model.Retrieval, proto.RetrievalDeleteRequest, proto.RetrievalDeleteResponse]
	Getter[model.Retrieval, proto.RetrievalGetRequest, proto.RetrievalGetResponse]
	Lister[model.Retrieval, proto.RetrievalListRequest, proto.RetrievalListResponse]

	proto.RetrievalServiceServer
}

type retrieval struct {
	service service.Retrieval
	proto.UnimplementedRetrievalServiceServer
}

func NewRetrieval(s service.Retrieval) Retrieval {
	return &retrieval{service: s}
}

func (r *retrieval) Create(ctx context.Context, req *proto.RetrievalCreateRequest) (*proto.RetrievalCreateResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	sellerUUID, err := getUUIDFromContext(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	if req.UserUUID == "" {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}

	err = r.service.Create(ctx, req.UserUUID, sellerUUID, req.SaleItemUUIDs)
	if err != nil {
		return nil, err
	}

	return &proto.RetrievalCreateResponse{}, nil
}

func (r *retrieval) Update(ctx context.Context, req *proto.RetrievalUpdateRequest) (*proto.RetrievalUpdateResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if req.Uuid == "" {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}

	retrieval, err := r.service.Update(ctx, req.Uuid, req.Delivered)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	return &proto.RetrievalUpdateResponse{
		Retrieval: &proto.Retrieval{
			Uuid: retrieval.UUID,
			User: &proto.User{
				Uuid: retrieval.User.UUID,
			},
			Seller: &proto.User{
				Uuid: retrieval.User.UUID,
			},
			SaleItem: &proto.SaleItem{
				Uuid: retrieval.SaleItem.UUID,
			},
			Delivered: retrieval.Delivered,
			CreatedAt: timestamppb.New(retrieval.CreatedAt),
			UpdatedAt: timestamppb.New(retrieval.UpdatedAt),
		},
	}, nil
}

func (r *retrieval) Delete(ctx context.Context, req *proto.RetrievalDeleteRequest) (*proto.RetrievalDeleteResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if req.Uuid == "" {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}

	err = r.service.Delete(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}
	return &proto.RetrievalDeleteResponse{}, nil
}

func (r *retrieval) Get(ctx context.Context, req *proto.RetrievalGetRequest) (*proto.RetrievalGetResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	if req.Uuid == "" {
		return nil, errs.New(errors.New(""), errCode.InvalidArgument)
	}

	retrieval, err := r.service.Get(ctx, req.Uuid)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	return &proto.RetrievalGetResponse{
		Retrieval: &proto.Retrieval{
			Uuid: retrieval.UUID,
			User: &proto.User{
				Uuid: retrieval.User.UUID,
			},
			Seller: &proto.User{
				Uuid: retrieval.User.UUID,
			},
			SaleItem: &proto.SaleItem{
				Uuid: retrieval.SaleItem.UUID,
			},
			Delivered: retrieval.Delivered,
			CreatedAt: timestamppb.New(retrieval.CreatedAt),
			UpdatedAt: timestamppb.New(retrieval.UpdatedAt),
		},
	}, nil
}

func (r *retrieval) List(ctx context.Context, req *proto.RetrievalListRequest) (*proto.RetrievalListResponse, error) {
	err := checkStaffAuthorization(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	retrievals, err := r.service.List(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}

	res := &proto.RetrievalListResponse{
		Retrievals: make([]*proto.Retrieval, 0, len(retrievals)),
	}
	for _, retrieval := range retrievals {
		res.Retrievals = append(res.Retrievals, &proto.Retrieval{
			Uuid: retrieval.UUID,
			User: &proto.User{
				Uuid: retrieval.User.UUID,
			},
			Seller: &proto.User{
				Uuid: retrieval.User.UUID,
			},
			SaleItem: &proto.SaleItem{
				Uuid: retrieval.SaleItem.UUID,
			},
			Delivered: retrieval.Delivered,
			CreatedAt: timestamppb.New(retrieval.CreatedAt),
			UpdatedAt: timestamppb.New(retrieval.UpdatedAt),
		})
	}
	return res, nil
}
