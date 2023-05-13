package handler

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto"
	"github.com/cristiancll/qrpay-be/internal/api/service"
)

type OperationLog interface {
	proto.OperationLogServiceServer
}

type operationLog struct {
	service service.OperationLog
	proto.UnimplementedOperationLogServiceServer
}

func NewOperationLog(s service.OperationLog) OperationLog {
	return &operationLog{service: s}
}
