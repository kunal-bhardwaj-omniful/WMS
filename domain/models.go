package domain

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Tenant struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null;unique" json:"name"`
	Email     string         `gorm:"type:varchar(100);not null;unique" json:"email"`
	GSTIN     *string        `gorm:"type:varchar(15)" json:"gstin,omitempty"`
	CreatedAt time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete support
}

type Hub struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TenantID  uuid.UUID `gorm:"type:uuid;not null;index" json:"tenant_id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Code      string    `gorm:"type:varchar(20);not null" json:"code"`
	Address   string    `gorm:"type:varchar(255);not null" json:"address"`
	City      *string   `gorm:"type:varchar(100)" json:"city,omitempty"`
	State     *string   `gorm:"type:varchar(100)" json:"state,omitempty"`
	Country   *string   `gorm:"type:varchar(100)" json:"country,omitempty"`
	Pincode   *string   `gorm:"type:varchar(20)" json:"pincode,omitempty"`
	Location  *string   `gorm:"type:varchar(30)" json:"location,omitempty"`
	CreatedAt time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete support

	Tenant Tenant `gorm:"foreignKey:TenantID;constraint:OnDelete:RESTRICT" json:"tenant"` // Relation with Tenant
}

type SKU struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	SellerID    uuid.UUID      `gorm:"type:uuid;not null" json:"seller_id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Code        string         `gorm:"type:varchar(50);not null" json:"code"`
	Description string         `gorm:"type:varchar(500)" json:"description"`
	Category    string         `gorm:"type:varchar(100)" json:"category"`
	Subcategory string         `gorm:"type:varchar(100)" json:"subcategory"`
	Brand       string         `gorm:"type:varchar(100)" json:"brand"`
	Model       string         `gorm:"type:varchar(100)" json:"model"`
	UOM         string         `gorm:"type:varchar(20);not null" json:"uom"` // Unit of Measure
	Weight      float64        `gorm:"type:numeric(10,3)" json:"weight"`
	Dimensions  datatypes.JSON `gorm:"type:jsonb" json:"dimensions"` // JSONB for storing dimensions
	CreatedAt   time.Time      `gorm:"type:timestamptz;default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;default:current_timestamp" json:"updated_at"`

	// Associations
	Seller Seller `gorm:"foreignKey:SellerID;references:ID" json:"seller"`
}
type Seller struct {
	ID            uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	TenantID      uuid.UUID `gorm:"type:uuid;not null" json:"tenant_id"`
	Name          string    `gorm:"size:50;not null" json:"name"`
	Code          string    `gorm:"size:20;not null;unique:tenant_code" json:"code"`
	ContactPerson string    `gorm:"size:100" json:"contact_person"`
	Email         string    `gorm:"size:100" json:"email"`
	Phone         string    `gorm:"size:20" json:"phone"`
	CreatedAt     time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Inventory struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	SkuID         uuid.UUID  `gorm:"type:uuid;not null;index" json:"sku_id"`
	HubID         uuid.UUID  `gorm:"type:uuid;not null;index" json:"hub_id"`
	AvailableQty  int        `gorm:"not null;default:0;check:available_qty >= 0" json:"available_qty"`
	AllocatedQty  int        `gorm:"not null;default:0;check:allocated_qty >= 0" json:"allocated_qty"`
	DamagedQty    int        `gorm:"not null;default:0;check:damaged_qty >= 0" json:"damaged_qty"`
	Zone          string     `gorm:"type:varchar(50)" json:"zone"`
	Rack          string     `gorm:"type:varchar(50)" json:"rack"`
	Bin           string     `gorm:"type:varchar(50)" json:"bin"`
	MinThreshold  int        `gorm:"default:0" json:"min_threshold"`
	MaxThreshold  int        `gorm:"default:0" json:"max_threshold"`
	LastCountedAt *time.Time `gorm:"type:timestamptz" json:"last_counted_at"`
	CreatedAt     time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updated_at"`
}
