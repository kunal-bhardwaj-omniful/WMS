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
	CreateHub(ctx context.Context, hub domain.Hub) error
	CreateSKU(ctx context.Context, sku domain.SKU) error
	DecreaseInventoryQty(ctx context.Context, skuID, hubID uuid.UUID, availableQty, allocatedQty, damagedQty int) error
}

type service struct {
	repo repo.Repository
}

// NewService is the constructor function to create a new instance of ConcreteService.
func NewService(r repo.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateHub(ctx context.Context, hub domain.Hub) error {
	// Ensure valid hub data before creation
	if hub.Name == "" {
		return fmt.Errorf("hub name cannot be empty")
	}

	// Call repository method to create the hub
	err := s.repo.CreateHub(ctx, hub)
	if err != nil {
		return fmt.Errorf("failed to create hub: %w", err)
	}
	return nil
}

func (s *service) CreateSKU(ctx context.Context, sku domain.SKU) error {
	// Ensure valid SKU data before creation
	if sku.Name == "" {
		return fmt.Errorf("SKU name cannot be empty")
	}

	// Call repository method to create the SKU
	err := s.repo.CreateSKU(ctx, sku)
	if err != nil {
		return fmt.Errorf("failed to create SKU: %w", err)
	}
	return nil
}

func (s *service) FetchHubs(ctx context.Context) ([]domain.Hub, error) {
	hubs, err := s.repo.GetAllHubs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hubs: %w", err)
	}
	return hubs, nil
}

func (s *service) FetchSkus(ctx context.Context) ([]domain.SKU, error) {
	skus, err := s.repo.GetAllSkus(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch SKUs: %w", err)
	}
	return skus, nil
}

func (s *service) FetchHubByID(ctx context.Context, id uuid.UUID) (domain.Hub, error) {
	if id == uuid.Nil {
		return domain.Hub{}, fmt.Errorf("invalid hub ID")
	}
	hub, err := s.repo.GetHubByID(ctx, id)
	if err != nil {
		return domain.Hub{}, fmt.Errorf("failed to fetch hub by ID: %w", err)
	}
	return hub, nil
}

func (s *service) FetchSkuByID(ctx context.Context, id uuid.UUID) (domain.SKU, error) {
	if id == uuid.Nil {
		return domain.SKU{}, fmt.Errorf("invalid SKU ID")
	}
	sku, err := s.repo.GetSkuByID(ctx, id)
	if err != nil {
		return domain.SKU{}, fmt.Errorf("failed to fetch SKU by ID: %w", err)
	}
	return sku, nil
}

func (s *service) DecreaseInventoryQty(ctx context.Context, skuID, hubID uuid.UUID, availableQty, allocatedQty, damagedQty int) error {
	// Validate the quantities before proceeding
	if availableQty < 0 || allocatedQty < 0 || damagedQty < 0 {
		return fmt.Errorf("quantities must be non-negative")
	}

	// Delegate to the repository to decrease inventory quantities
	err := s.repo.DecreaseInventoryQty(ctx, skuID, hubID, availableQty, allocatedQty, damagedQty)
	if err != nil {
		return fmt.Errorf("failed to decrease inventory quantities: %w", err)
	}
	return nil
}
