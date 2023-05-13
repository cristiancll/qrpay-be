package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type OperationLogHandler interface {
	proto.OperationLogServiceServer
}

type operationLogHandler struct {
	service service.OperationLogService
	proto.UnimplementedOperationLogServiceServer
}

func NewOperationLogHandler(s service.OperationLogService) OperationLogHandler {
	return &operationLogHandler{service: s}
}
