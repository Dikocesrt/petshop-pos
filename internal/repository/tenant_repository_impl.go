package repository

import (
	"petshop-pos/internal/entity"

	"gorm.io/gorm"
)

type TenantRepositoryImpl struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepositoryImpl {
	db.AutoMigrate(&entity.Tenant{})
	return &TenantRepositoryImpl{db: db}
}
