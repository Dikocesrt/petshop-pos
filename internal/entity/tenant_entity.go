package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tenant struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:id"`
	Name      string     `gorm:"type:varchar(255);not null;column:name;unique"`
	Location  string     `gorm:"type:varchar(255);column:location"`
	CreatedAt time.Time  `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (Tenant) TableName() string {
	return "tenants"
}