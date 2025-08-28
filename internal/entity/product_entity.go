package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
    ID         string      `gorm:"type:char(36);primaryKey;default:(UUID())"`
    Name       string         `gorm:"type:varchar(255);not null"`
    Stock      int            `gorm:"not null"`
    Price      int            `gorm:"not null"`
    TenantID   string      `gorm:"type:char(36);not null"`
    BrandID    *string     `gorm:"type:char(36)"`
    CategoryID *string      `gorm:"type:char(36)"`
    CreatedAt  time.Time      `gorm:"autoCreateTime"`
    UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
    DeletedAt  gorm.DeletedAt `gorm:"index"`

    // Relations
    Tenant   Tenant   `gorm:"foreignKey:TenantID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
    Brand    *Brand   `gorm:"foreignKey:BrandID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
    Category *Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (p *Product) TableName() string {
    return "products"
}
