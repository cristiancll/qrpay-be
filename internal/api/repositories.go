package server

import (
	"errors"
	errs "github.com/cristiancll/go-errors"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/errCode"
	"github.com/cristiancll/qrpay-be/internal/errMsg"
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
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.auth = repository.NewAuth(s.db)
	if err := s.repos.auth.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.category = repository.NewCategory(s.db)
	if err := s.repos.category.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.item = repository.NewItem(s.db)
	if err := s.repos.item.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.opLog = repository.NewOperationLog(s.db)
	if err := s.repos.opLog.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.sale = repository.NewSale(s.db)
	if err := s.repos.sale.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.sku = repository.NewSKU(s.db)
	if err := s.repos.sku.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.saleItem = repository.NewSaleItem(s.db)
	if err := s.repos.saleItem.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.retrieval = repository.NewRetrieval(s.db)
	if err := s.repos.retrieval.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}

	s.repos.stock = repository.NewStock(s.db)
	if err := s.repos.stock.Migrate(s.context); err != nil {
		return errs.New(errors.New(errMsg.FailedToMigrateRepository), errCode.Internal)
	}
	return nil
}
