package server

import "github.com/cristiancll/qrpay-be/internal/api/handler"

type handlers struct {
	user      handler.UserHandler
	auth      handler.AuthHandler
	wpp       handler.WhatsAppHandler
	category  handler.CategoryHandler
	item      handler.ItemHandler
	opLog     handler.OperationLogHandler
	order     handler.OrderHandler
	orderItem handler.OrderItemHandler
	retrieval handler.RetrievalHandler
	sku       handler.SKUHandler
	stock     handler.StockHandler
}

func (s *Server) createHandlers() {
	s.handlers.user = handler.NewUserHandler(s.services.user)
	s.handlers.auth = handler.NewAuthHandler(s.services.auth)
	s.handlers.wpp = handler.NewWhatsAppHandler(s.services.wpp)
	s.handlers.category = handler.NewCategoryHandler(s.services.category)
	s.handlers.item = handler.NewItemHandler(s.services.item)
	s.handlers.opLog = handler.NewOperationLogHandler(s.services.opLog)
	s.handlers.order = handler.NewOrderHandler(s.services.order)
	s.handlers.orderItem = handler.NewOrderItemHandler(s.services.orderItem)
	s.handlers.retrieval = handler.NewRetrievalHandler(s.services.retrieval)
	s.handlers.sku = handler.NewSKUHandler(s.services.sku)
	s.handlers.stock = handler.NewStockHandler(s.services.stock)
}
