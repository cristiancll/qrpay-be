package server

import (
	"github.com/cristiancll/qrpay-be/internal/api/proto/generated"
	"google.golang.org/grpc"
)

func (s *Server) registerServices(grpcServer *grpc.Server) {
	proto.RegisterUserServiceServer(grpcServer, s.handlers.user)
	proto.RegisterAuthServiceServer(grpcServer, s.handlers.auth)
	proto.RegisterWhatsAppServiceServer(grpcServer, s.handlers.wpp)
	proto.RegisterCategoryServiceServer(grpcServer, s.handlers.category)
	proto.RegisterItemServiceServer(grpcServer, s.handlers.item)
	proto.RegisterOperationLogServiceServer(grpcServer, s.handlers.opLog)
	proto.RegisterOrderServiceServer(grpcServer, s.handlers.order)
	proto.RegisterOrderItemServiceServer(grpcServer, s.handlers.orderItem)
	proto.RegisterRetrievalServiceServer(grpcServer, s.handlers.retrieval)
	proto.RegisterSKUServiceServer(grpcServer, s.handlers.sku)
	proto.RegisterStockServiceServer(grpcServer, s.handlers.stock)
}
