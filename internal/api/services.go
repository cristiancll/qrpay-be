package server

import (
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/wpp"
)

type services struct {
	user      service.UserService
	auth      service.AuthService
	wpp       service.WhatsAppService
	category  service.CategoryService
	item      service.ItemService
	opLog     service.OperationLogService
	order     service.OrderService
	orderItem service.OrderItemService
	retrieval service.RetrievalService
	sku       service.SKUService
	stock     service.StockService
}

func (s *Server) createServices(wppSystem wpp.WhatsAppSystem) {
	s.services.user = service.NewUserService(s.db, wppSystem, s.repos.user, s.repos.auth)
	s.services.auth = service.NewAuthService(s.db, s.repos.auth, s.repos.user)
	s.services.wpp = service.NewWhatsAppService(s.db, wppSystem, s.repos.wpp)
	s.services.category = service.NewCategoryService(s.db, s.repos.category)
	s.services.item = service.NewItemService(s.db, s.repos.item)
	s.services.opLog = service.NewOperationLogService(s.db, s.repos.opLog)
	s.services.order = service.NewOrderService(s.db, s.repos.order)
	s.services.orderItem = service.NewOrderItemService(s.db, s.repos.orderItem)
	s.services.retrieval = service.NewRetrievalService(s.db, s.repos.retrieval)
	s.services.sku = service.NewSKUService(s.db, s.repos.sku)
	s.services.stock = service.NewStockService(s.db, s.repos.stock)
}
