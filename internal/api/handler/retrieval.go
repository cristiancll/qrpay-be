package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type RetrievalHandler interface {
	proto.RetrievalServiceServer
}

type retrievalHandler struct {
	service service.RetrievalService
	proto.UnimplementedRetrievalServiceServer
}

func NewRetrievalHandler(s service.RetrievalService) RetrievalHandler {
	return &retrievalHandler{service: s}
}
