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
	sale      service.Sale
	saleItem  service.SaleItem
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
	s.services.sale = service.NewSale(s.db, s.repos.sale)
	s.services.saleItem = service.NewSaleItem(s.db, s.repos.saleItem)
	s.services.retrieval = service.NewRetrieval(s.db, s.repos.retrieval)
	s.services.sku = service.NewSKU(s.db, s.repos.sku)
	s.services.stock = service.NewStock(s.db, s.repos.stock)
}
