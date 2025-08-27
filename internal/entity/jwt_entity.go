package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshToken struct {
    ID          uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID      uuid.UUID      `gorm:"type:uuid;not null;column:user_id"`
    Token       string         `gorm:"type:varchar(255);not null;unique;column:token"`
    ExpiresAt   time.Time      `gorm:"not null;column:expires_at"`
	CreatedAt   time.Time      `gorm:"autoCreateTime;column:created_at"`
    UpdatedAt   time.Time      `gorm:"autoUpdateTime;column:updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (RefreshToken) TableName() string {
    return "refresh_tokens"
}