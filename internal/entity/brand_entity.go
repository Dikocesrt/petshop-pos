package entity

import (
	"time"

	"gorm.io/gorm"
)

type Brand struct {
    ID        string      `gorm:"type:char(36);primaryKey;default:(UUID())"`
    Name      string         `gorm:"type:varchar(255);not null"`
    TenantID   string      `gorm:"type:char(36);not null"`
    CreatedAt time.Time      `gorm:"autoCreateTime"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime"`
    DeletedAt gorm.DeletedAt `gorm:"index"`

    // Relations
    Products []Product `gorm:"foreignKey:BrandID"`
    Tenant   Tenant   `gorm:"foreignKey:TenantID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}

func (b *Brand) TableName() string {
    return "brands"
}