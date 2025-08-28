package repository

import (
	"context"

	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"

	"gorm.io/gorm"
)

type categoryRepositoryImpl struct {
    db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    db.AutoMigrate(&entity.Category{})
    return &categoryRepositoryImpl{db: db}
}

func (r *categoryRepositoryImpl) Create(ctx context.Context, category *entity.Category) *exception.Exception {
    if err := r.db.WithContext(ctx).Create(category).Error; err != nil {
        return exception.Internal("failed to create category", err)
    }
    return nil
}

func (r *categoryRepositoryImpl) FindByIDAndTenantID(ctx context.Context, id string, tenantID string) (*entity.Category, *exception.Exception) {
    var category entity.Category
    if err := r.db.WithContext(ctx).Preload("Tenant").Where("id = ? AND tenant_id = ?", id, tenantID).First(&category).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return nil, exception.NotFound("category not found")
        }
        return nil, exception.Internal("failed to find category", err)
    }
    return &category, nil
}

func (r *categoryRepositoryImpl) FindAllByTenantID(ctx context.Context, page, limit int, tenantID string) ([]entity.Category, int64, *exception.Exception) {
    var categories []entity.Category
    var total int64

    offset := (page - 1) * limit

    if err := r.db.WithContext(ctx).Model(&entity.Category{}).Where("tenant_id = ?", tenantID).Count(&total).Error; err != nil {
        return nil, 0, exception.Internal("failed to count categories", err)
    }

    if err := r.db.WithContext(ctx).Preload("Tenant").Where("tenant_id = ?", tenantID).Offset(offset).Limit(limit).Find(&categories).Error; err != nil {
        return nil, 0, exception.Internal("failed to get categories", err)
    }

    return categories, total, nil
}

func (r *categoryRepositoryImpl) Update(ctx context.Context, category *entity.Category) *exception.Exception {
    if err := r.db.WithContext(ctx).Save(category).Error; err != nil {
        return exception.Internal("failed to update category", err)
    }
    return nil
}

func (r *categoryRepositoryImpl) Delete(ctx context.Context, id string, tenantID string) *exception.Exception {
    if err := r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&entity.Category{}).Error; err != nil {
        return exception.Internal("failed to delete category", err)
    }
    return nil
}