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
	proto.RegisterSaleServiceServer(grpcServer, s.handlers.sale)
	proto.RegisterRetrievalServiceServer(grpcServer, s.handlers.retrieval)
	proto.RegisterSKUServiceServer(grpcServer, s.handlers.sku)
	proto.RegisterStockServiceServer(grpcServer, s.handlers.stock)
}
