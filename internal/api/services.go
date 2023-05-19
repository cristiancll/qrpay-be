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
	sale      service.Sale
	retrieval service.Retrieval
	sku       service.SKU
	stock     service.Stock
}

func (s *Server) createServices(wppSystem wpp.WhatsAppSystem) {
	s.services.user = service.NewUser(s.db, wppSystem, s.repos.user, s.repos.auth, s.repos.opLog)
	s.services.auth = service.NewAuth(s.db, s.repos.auth, s.repos.user, s.repos.opLog)
	s.services.wpp = service.NewWhatsApp(s.db, wppSystem, s.repos.wpp, s.repos.opLog)
	s.services.category = service.NewCategory(s.db, s.repos.category, s.repos.opLog)
	s.services.item = service.NewItem(s.db, s.repos.item, s.repos.category, s.repos.opLog)
	s.services.sale = service.NewSale(s.db, wppSystem, s.repos.sale, s.repos.sku, s.repos.user, s.repos.saleItem, s.repos.stock, s.repos.opLog)
	s.services.retrieval = service.NewRetrieval(s.db, wppSystem, s.repos.retrieval, s.repos.user, s.repos.saleItem, s.repos.opLog)
	s.services.sku = service.NewSKU(s.db, s.repos.sku, s.repos.item, s.repos.opLog)
	s.services.stock = service.NewStock(s.db, s.repos.stock, s.repos.sku, s.repos.opLog)
}
