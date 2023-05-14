package server

import "github.com/cristiancll/qrpay-be/internal/api/handler"

type handlers struct {
	user      handler.User
	auth      handler.Auth
	wpp       handler.WhatsApp
	category  handler.Category
	item      handler.Item
	opLog     handler.OperationLog
	sale      handler.Sale
	saleItem  handler.SaleItem
	retrieval handler.Retrieval
	sku       handler.SKU
	stock     handler.Stock
}

func (s *Server) createHandlers() {
	s.handlers.user = handler.NewUser(s.services.user)
	s.handlers.auth = handler.NewAuth(s.services.auth)
	s.handlers.wpp = handler.NewWhatsApp(s.services.wpp)
	s.handlers.category = handler.NewCategory(s.services.category)
	s.handlers.item = handler.NewItem(s.services.item)
	s.handlers.opLog = handler.NewOperationLog(s.services.opLog)
	s.handlers.sale = handler.NewSale(s.services.sale)
	s.handlers.saleItem = handler.NewSaleItem(s.services.saleItem)
	s.handlers.retrieval = handler.NewRetrieval(s.services.retrieval)
	s.handlers.sku = handler.NewSKU(s.services.sku)
	s.handlers.stock = handler.NewStock(s.services.stock)
}
