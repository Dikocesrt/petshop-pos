package entity

import (
	"time"

	"gorm.io/gorm"
)

type Tenant struct {
	ID          string `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Name        string `gorm:"type:varchar(255);not null;column:name;unique"`
	Location    string `gorm:"type:varchar(255);column:location"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (Tenant) TableName() string {
	return "tenants"
}