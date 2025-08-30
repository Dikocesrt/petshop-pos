package repository

import (
	"context"

	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"

	"gorm.io/gorm"
)

type brandRepositoryImpl struct {
    db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
    db.AutoMigrate(&entity.Brand{})
    return &brandRepositoryImpl{db: db}
}

func (r *brandRepositoryImpl) Create(ctx context.Context, brand *entity.Brand) *exception.Exception {
    if err := r.db.WithContext(ctx).Create(brand).Error; err != nil {
        return exception.Internal("failed to create brand", err)
    }
    return nil
}

func (r *brandRepositoryImpl) FindByIDAndTenantID(ctx context.Context, id string, tenantID string) (*entity.Brand, *exception.Exception) {
    var brand entity.Brand
    if err := r.db.WithContext(ctx).Preload("Tenant").Where("id = ? AND tenant_id = ?", id, tenantID).First(&brand).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, exception.NotFound("brand not found")
        }
        return nil, exception.Internal("failed to find brand", err)
    }
    return &brand, nil
}

func (r *brandRepositoryImpl) FindAllByTenantID(ctx context.Context, tenantID string, page, limit int) ([]entity.Brand, int64, *exception.Exception) {
    var brands []entity.Brand
    var total int64

    offset := (page - 1) * limit

    if err := r.db.WithContext(ctx).Model(&entity.Brand{}).Where("tenant_id = ?", tenantID).Count(&total).Error; err != nil {
        return nil, 0, exception.Internal("failed to count brands", err)
    }

    if err := r.db.WithContext(ctx).Preload("Tenant").Where("tenant_id = ?", tenantID).Offset(offset).Limit(limit).Find(&brands).Error; err != nil {
        return nil, 0, exception.Internal("failed to get brands", err)
    }

    return brands, total, nil
}

func (r *brandRepositoryImpl) Update(ctx context.Context, brand *entity.Brand) *exception.Exception {
    if err := r.db.WithContext(ctx).Save(brand).Error; err != nil {
        return exception.Internal("failed to update brand", err)
    }
    return nil
}

func (r *brandRepositoryImpl) Delete(ctx context.Context, id string, tenantID string) *exception.Exception {
    if err := r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&entity.Brand{}).Error; err != nil {
        return exception.Internal("failed to delete brand", err)
    }
    return nil
}

func (r *brandRepositoryImpl) IsBrandExistsByIDAndTenantID(ctx context.Context, id string, tenantID string) (bool, *exception.Exception) {
    var count int64
    if err := r.db.WithContext(ctx).Model(&entity.Brand{}).Where("id = ? AND tenant_id = ?", id, tenantID).Count(&count).Error; err != nil {
        return false, exception.Internal("failed to check if brand exists", err)
    }
    return count > 0, nil
}