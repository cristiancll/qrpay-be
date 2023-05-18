package service

import (
	"context"
	"github.com/cristiancll/qrpay-be/internal/api/model"
	"github.com/cristiancll/qrpay-be/internal/api/repository"
	"github.com/cristiancll/qrpay-be/internal/wpp"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Sale interface {
	Create(ctx context.Context, userUUID string, sellerUUID string, saleItems map[string]int64) (*model.Sale, error)
}

type sale struct {
	pool         *pgxpool.Pool
	repo         repository.Sale
	skuRepo      repository.SKU
	userRepo     repository.User
	saleItemRepo repository.SaleItem
	wpp          wpp.WhatsAppSystem
}

func NewSale(pool *pgxpool.Pool, wpp wpp.WhatsAppSystem, r repository.Sale, skuRepo repository.SKU, userRepo repository.User, saleItemRepo repository.SaleItem) Sale {
	return &sale{
		pool:         pool,
		repo:         r,
		skuRepo:      skuRepo,
		userRepo:     userRepo,
		saleItemRepo: saleItemRepo,
		wpp:          wpp,
	}
}

func (s sale) Create(ctx context.Context, userUUID string, sellerUUID string, saleUnits map[string]int64) (*model.Sale, error) {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// Validates user
	user, err := s.userRepo.TGetByUUID(ctx, tx, userUUID)
	if err != nil {
		return nil, err
	}

	// Validates seller
	seller, err := s.userRepo.TGetByUUID(ctx, tx, sellerUUID)
	if err != nil {
		return nil, err
	}

	// Gets SKUs UUIDs
	skusIDs := make([]string, len(saleUnits))
	for skuUUID, _ := range saleUnits {
		skusIDs = append(skusIDs, skuUUID)
	}

	// Validates SKUs
	skus, err := s.skuRepo.TGetAllByUUIDs(ctx, tx, skusIDs)
	if err != nil {
		return nil, err
	}

	// Calculates total amount
	var total int64
	for _, sku := range skus {
		total += saleUnits[sku.UUID] * sku.Price
	}

	// Creates sale
	sale := &model.Sale{
		User:   *user,
		Seller: *seller,
		Total:  total,
		Paid:   true,
	}
	err = s.repo.TCreate(ctx, tx, sale)
	if err != nil {
		return nil, err
	}

	// Creates sale items
	saleItems := make([]*model.SaleItem, len(skus))
	for i, sku := range skus {
		saleItems[i] = &model.SaleItem{
			Sale:     *sale,
			SKU:      *sku,
			Quantity: saleUnits[sku.UUID],
		}
		err = s.saleItemRepo.TCreate(ctx, tx, saleItems[i])
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}
	go s.wpp.SendText(user, user.NewSale(sale, saleItems))
	return sale, nil
}
