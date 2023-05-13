package server

import (
	"github.com/cristiancll/qrpay-be/internal/api/service"
	"github.com/cristiancll/qrpay-be/internal/wpp"
)

type services struct {
	user      service.User
	auth      service.Auth
	wpp       service.WhatsApp
	category  service.Category
	item      service.Item
	opLog     service.OperationLog
	order     service.Order
	orderItem service.OrderItem
	retrieval service.Retrieval
	sku       service.SKU
	stock     service.Stock
}

func (s *Server) createServices(wppSystem wpp.WhatsAppSystem) {
	s.services.user = service.NewUser(s.db, wppSystem, s.repos.user, s.repos.auth)
	s.services.auth = service.NewAuth(s.db, s.repos.auth, s.repos.user)
	s.services.wpp = service.NewWhatsApp(s.db, wppSystem, s.repos.wpp)
	s.services.category = service.NewCategory(s.db, s.repos.category)
	s.services.item = service.NewItem(s.db, s.repos.item)
	s.services.opLog = service.NewOperationLog(s.db, s.repos.opLog)
	s.services.order = service.NewOrder(s.db, s.repos.order)
	s.services.orderItem = service.NewOrderItem(s.db, s.repos.orderItem)
	s.services.retrieval = service.NewRetrieval(s.db, s.repos.retrieval)
	s.services.sku = service.NewSKU(s.db, s.repos.sku)
	s.services.stock = service.NewStock(s.db, s.repos.stock)
}
