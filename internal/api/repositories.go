package server

import (
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
)

type repositories struct {
	user      repository.User
	auth      repository.Auth
	category  repository.Category
	item      repository.Item
	opLog     repository.OperationLog
	sale      repository.Sale
	saleItem  repository.SaleItem
	retrieval repository.Retrieval
	sku       repository.SKU
	stock     repository.Stock
}

func (s *Server) createRepositories() error {
	s.repos.user = repository.NewUser(s.db)
	if err := s.repos.user.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate user repository: %v", err)
	}

	s.repos.auth = repository.NewAuth(s.db)
	if err := s.repos.auth.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate auth repository: %v", err)
	}

	s.repos.category = repository.NewCategory(s.db)
	if err := s.repos.category.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate category repository: %v", err)
	}

	s.repos.item = repository.NewItem(s.db)
	if err := s.repos.item.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate item repository: %v", err)
	}

	s.repos.opLog = repository.NewOperationLog(s.db)
	if err := s.repos.opLog.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate operation log repository: %v", err)
	}

	s.repos.sale = repository.NewSale(s.db)
	if err := s.repos.sale.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate sale repository: %v", err)
	}

	s.repos.sku = repository.NewSKU(s.db)
	if err := s.repos.sku.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate sku repository: %v", err)
	}

	s.repos.saleItem = repository.NewSaleItem(s.db)
	if err := s.repos.saleItem.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate sale item repository: %v", err)
	}

	s.repos.retrieval = repository.NewRetrieval(s.db)
	if err := s.repos.retrieval.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate retrieval repository: %v", err)
	}

	s.repos.stock = repository.NewStock(s.db)
	if err := s.repos.stock.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate stock repository: %v", err)
	}
	return nil
}
