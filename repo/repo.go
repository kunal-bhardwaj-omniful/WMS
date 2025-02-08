package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/omniful/go_commons/db/sql/postgres"
	"sync"
	"wms/domain"
)

type Repository interface {
	GetAllHubs(ctx context.Context) ([]domain.Hub, error)
	GetAllSkus(ctx context.Context) ([]domain.SKU, error)
	GetHubByID(ctx context.Context, id uuid.UUID) (domain.Hub, error)
	GetSkuByID(ctx context.Context, id uuid.UUID) (domain.SKU, error)
	CreateHub(ctx context.Context, hub domain.Hub) error
	CreateSKU(ctx context.Context, sku domain.SKU) error
	DecreaseAvailableQty(ctx context.Context, skuID, hubID uuid.UUID, qty int) error
	DecreaseAllocatedQty(ctx context.Context, skuID, hubID uuid.UUID, qty int) error
	DecreaseDamagedQty(ctx context.Context, skuID, hubID uuid.UUID, qty int) error
	DecreaseInventoryQty(ctx context.Context, skuID, hubID uuid.UUID, availableQty, allocatedQty, damagedQty int) error
	GetInventory(ctx context.Context, skuID, hubID uuid.UUID) (domain.Inventory, error)
}

type repository struct {
	db *postgres.DbCluster
}

var repo *repository
var repoOnce sync.Once

func NewRepository(db *postgres.DbCluster) Repository {
	repoOnce.Do(func() {
		// Initialize the Repository with a given DbCluster.
		repo = &repository{
			db: db,
		}
	})
	return repo
}

func (r *repository) CreateHub(ctx context.Context, hub domain.Hub) error {
	// Insert the new hub into the database
	err := r.db.GetMasterDB(ctx).Create(&hub).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CreateSKU(ctx context.Context, sku domain.SKU) error {
	// Insert the new SKU into the database
	err := r.db.GetMasterDB(ctx).Create(&sku).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllHubs(ctx context.Context) ([]domain.Hub, error) {
	var hubs []domain.Hub
	err := r.db.GetMasterDB(ctx).Find(&hubs).Error
	if err != nil {
		return nil, errors.New("failed to fetch hubs")
	}
	return hubs, nil
}

func (r *repository) GetAllSkus(ctx context.Context) ([]domain.SKU, error) {
	var skus []domain.SKU
	err := r.db.GetMasterDB(ctx).Find(&skus).Error
	if err != nil {
		return nil, errors.New("failed to fetch SKUs")
	}
	return skus, nil
}

// GetHubByID fetches a single hub by ID from the database
func (r *repository) GetHubByID(ctx context.Context, id uuid.UUID) (domain.Hub, error) {
	var hub domain.Hub
	err := r.db.GetMasterDB(ctx).Where("id = ?", id).First(&hub).Error
	if err != nil {
		return domain.Hub{}, errors.New("hub not found")
	}
	return hub, nil
}

// GetSkuByID fetches a single SKU by ID from the database
func (r *repository) GetSkuByID(ctx context.Context, id uuid.UUID) (domain.SKU, error) {
	var sku domain.SKU
	err := r.db.GetMasterDB(ctx).Where("id = ?", id).First(&sku).Error
	if err != nil {
		return domain.SKU{}, errors.New("SKU not found")
	}
	return sku, nil
}

func (r *repository) DecreaseAvailableQty(ctx context.Context, skuID, hubID uuid.UUID, qty int) error {
	// Decrease available_qty by the specified quantity
	result := r.db.GetMasterDB(ctx).Exec(`
		UPDATE inventories
		SET available_qty = available_qty - $1, updated_at = CURRENT_TIMESTAMP
		WHERE sku_id = $2 AND hub_id = $3 AND available_qty >= $1
	`, qty, skuID, hubID)

	if result.Error != nil {
		return fmt.Errorf("failed to decrease available quantity: %v", result.Error)
	}

	// If no rows were affected, it means there wasn't enough stock
	if result.RowsAffected == 0 {
		return fmt.Errorf("not enough available quantity")
	}

	return nil
}

func (r *repository) DecreaseAllocatedQty(ctx context.Context, skuID, hubID uuid.UUID, qty int) error {
	// Decrease allocated_qty by the specified quantity
	result := r.db.GetMasterDB(ctx).Exec(`
		UPDATE inventories
		SET allocated_qty = allocated_qty - $1, updated_at = CURRENT_TIMESTAMP
		WHERE sku_id = $2 AND hub_id = $3 AND allocated_qty >= $1
	`, qty, skuID, hubID)

	if result.Error != nil {
		return fmt.Errorf("failed to decrease allocated quantity: %v", result.Error)
	}

	// If no rows were affected, it means there wasn't enough stock allocated
	if result.RowsAffected == 0 {
		return fmt.Errorf("not enough allocated quantity")
	}

	return nil
}
func (r *repository) DecreaseDamagedQty(ctx context.Context, skuID, hubID uuid.UUID, qty int) error {
	// Decrease damaged_qty by the specified quantity
	result := r.db.GetMasterDB(ctx).Exec(`
		UPDATE inventories
		SET damaged_qty = damaged_qty - $1, updated_at = CURRENT_TIMESTAMP
		WHERE sku_id = $2 AND hub_id = $3 AND damaged_qty >= $1
	`, qty, skuID, hubID)

	if result.Error != nil {
		return fmt.Errorf("failed to decrease damaged quantity: %v", result.Error)
	}

	// If no rows were affected, it means there wasn't enough damaged stock
	if result.RowsAffected == 0 {
		return fmt.Errorf("not enough damaged quantity")
	}

	return nil
}

func (r *repository) DecreaseInventoryQty(ctx context.Context, skuID, hubID uuid.UUID, availableQty, allocatedQty, damagedQty int) error {
	// Start a transaction
	tx := r.db.GetMasterDB(ctx).Begin()

	// Ensure the transaction is rolled back if any error occurs
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update available_qty
	if availableQty > 0 {
		if err := r.DecreaseAvailableQty(ctx, skuID, hubID, availableQty); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update allocated_qty
	if allocatedQty > 0 {
		if err := r.DecreaseAllocatedQty(ctx, skuID, hubID, allocatedQty); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Update damaged_qty
	if damagedQty > 0 {
		if err := r.DecreaseDamagedQty(ctx, skuID, hubID, damagedQty); err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction if all operations are successful
	tx.Commit()

	return nil
}

func (r *repository) GetInventory(ctx context.Context, skuID, hubID uuid.UUID) (domain.Inventory, error) {
	var inventory domain.Inventory

	err := r.db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuID, hubID).First(&inventory).Error
	if err != nil {
		return domain.Inventory{}, fmt.Errorf("failed to fetch inventory: %v", err)
	}

	return inventory, nil
}
