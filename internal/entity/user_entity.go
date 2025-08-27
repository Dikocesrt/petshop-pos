package entity

import (
	userroleconst "petshop-pos/internal/const"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
    ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid();column:id"`
    Name        string         `gorm:"type:varchar(255);not null;column:name"`
    Username    string         `gorm:"type:varchar(255);not null;column:username;unique"`
    Password    string         `gorm:"type:varchar(255);not null;column:password"`
    PhoneNumber string         `gorm:"type:varchar(255);column:phone_number"`
    Role        userroleconst.UserRole `gorm:"type:varchar(255);not null;column:role"`
    TenantID    uuid.UUID      `gorm:"type:uuid;not null;column:tenant_id"`
    CreatedAt   time.Time      `gorm:"autoCreateTime;column:created_at"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime;column:updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (User) TableName() string {
    return "users"
}