package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"wms/domain"
	"wms/repo"
)

type Service interface {
	FetchHubs(ctx context.Context) ([]domain.Hub, error)
	FetchSkus(ctx context.Context) ([]domain.SKU, error)
	FetchHubByID(ctx context.Context, id uuid.UUID) (domain.Hub, error)
	FetchSkuByID(ctx context.Context, id uuid.UUID) (domain.SKU, error)
	FetchInventory(ctx context.Context, skuID, hubID uuid.UUID) (domain.Inventory, error)
	CreateHub(ctx context.Context, hub domain.Hub) error
	CreateSKU(ctx context.Context, sku domain.SKU) error
	DecreaseInventoryQty(ctx context.Context, skuID, hubID uuid.UUID, Qty int) error
}

type service struct {
	repo repo.Repository
}

// NewService creates a new instance of the service.
func NewService(r repo.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateHub(ctx context.Context, hub domain.Hub) error {
	if hub.Name == "" {
		return fmt.Errorf("hub name cannot be empty")
	}
	return s.repo.CreateHub(ctx, hub)
}

func (s *service) CreateSKU(ctx context.Context, sku domain.SKU) error {
	if sku.Name == "" {
		return fmt.Errorf("SKU name cannot be empty")
	}
	return s.repo.CreateSKU(ctx, sku)
}

func (s *service) FetchHubs(ctx context.Context) ([]domain.Hub, error) {
	return s.repo.GetAllHubs(ctx)
}

func (s *service) FetchSkus(ctx context.Context) ([]domain.SKU, error) {
	return s.repo.GetAllSkus(ctx)
}

func (s *service) FetchHubByID(ctx context.Context, id uuid.UUID) (domain.Hub, error) {
	if id == uuid.Nil {
		return domain.Hub{}, fmt.Errorf("invalid hub ID")
	}
	return s.repo.GetHubByID(ctx, id)
}

func (s *service) FetchSkuByID(ctx context.Context, id uuid.UUID) (domain.SKU, error) {
	if id == uuid.Nil {
		return domain.SKU{}, fmt.Errorf("invalid SKU ID")
	}
	return s.repo.GetSkuByID(ctx, id)
}

// FetchInventory retrieves inventory details based on SKU ID and Hub ID
func (s *service) FetchInventory(ctx context.Context, skuID, hubID uuid.UUID) (domain.Inventory, error) {
	if skuID == uuid.Nil || hubID == uuid.Nil {
		return domain.Inventory{}, fmt.Errorf("invalid SKU ID or Hub ID")
	}
	return s.repo.GetInventory(ctx, skuID, hubID)
}

func (s *service) DecreaseInventoryQty(ctx context.Context, skuID, hubID uuid.UUID, Qty int) error {
	if Qty < 0 {
		return fmt.Errorf("quantities must be non-negative")
	}
	return s.repo.DecreaseAvailableQty(ctx, skuID, hubID, Qty)
}
