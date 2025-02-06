package domain

import (
	"github.com/google/uuid"
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
