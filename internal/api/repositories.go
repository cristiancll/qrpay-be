package server

import (
	"fmt"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
)

type repositories struct {
	user      repository.UserRepository
	auth      repository.AuthRepository
	wpp       repository.WhatsAppRepository
	category  repository.CategoryRepository
	item      repository.ItemRepository
	opLog     repository.OperationLogRepository
	order     repository.OrderRepository
	orderItem repository.OrderItemRepository
	retrieval repository.RetrievalRepository
	sku       repository.SKURepository
	stock     repository.StockRepository
}

func (s *Server) createRepositories() error {
	s.repos.user = repository.NewUserRepository(s.db)
	if err := s.repos.user.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate user repository: %v", err)
	}

	s.repos.auth = repository.NewAuthRepository(s.db)
	if err := s.repos.auth.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate auth repository: %v", err)
	}

	s.repos.wpp = repository.NewWhatsAppRepository(s.db)
	if err := s.repos.wpp.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate whatsapp repository: %v", err)
	}

	s.repos.category = repository.NewCategoryRepository(s.db)
	if err := s.repos.category.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate category repository: %v", err)
	}

	s.repos.item = repository.NewItemRepository(s.db)
	if err := s.repos.item.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate item repository: %v", err)
	}

	s.repos.opLog = repository.NewOperationLogRepository(s.db)
	if err := s.repos.opLog.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate operation log repository: %v", err)
	}

	s.repos.order = repository.NewOrderRepository(s.db)
	if err := s.repos.order.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate order repository: %v", err)
	}

	s.repos.retrieval = repository.NewRetrievalRepository(s.db)
	if err := s.repos.retrieval.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate retrieval repository: %v", err)
	}

	s.repos.sku = repository.NewSKURepository(s.db)
	if err := s.repos.sku.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate sku repository: %v", err)
	}

	s.repos.stock = repository.NewStockRepository(s.db)
	if err := s.repos.stock.Migrate(s.context); err != nil {
		return fmt.Errorf("unable to migrate stock repository: %v", err)
	}
	return nil
}
