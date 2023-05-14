package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type Retrieval interface {
	proto.RetrievalServiceServer
}

type retrieval struct {
	service service.Retrieval
	proto.UnimplementedRetrievalServiceServer
}

func NewRetrieval(s service.Retrieval) Retrieval {
	return &retrieval{service: s}
}
