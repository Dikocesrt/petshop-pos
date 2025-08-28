package repository

import (
	"context"
	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"

	"gorm.io/gorm"
)

type TenantRepositoryImpl struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepositoryImpl {
	db.AutoMigrate(&entity.Tenant{})
	return &TenantRepositoryImpl{db: db}
}

func (r *TenantRepositoryImpl) FindIDByName(ctx context.Context, name string) (string, *exception.Exception) {
	var tenant entity.Tenant
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&tenant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil
		}
		return "", exception.Internal("failed to find tenant by name", err)
	}
	return tenant.ID, nil
}
