package repository

import (
	"context"

	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"

	"gorm.io/gorm"
)

type productRepositoryImpl struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&entity.Product{})
    return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) Create(ctx context.Context, product *entity.Product) *exception.Exception {
	if err := r.db.WithContext(ctx).Create(product).Error; err != nil {
		return exception.Internal("failed to create product", err)
	}
	return nil
}

func (r *productRepositoryImpl) FindByIDAndTenantID(ctx context.Context, id string, tenantID string) (*entity.Product, *exception.Exception) {
	var product entity.Product
	if err := r.db.WithContext(ctx).Preload("Brand").Preload("Category").Preload("Tenant").Where("id = ? AND tenant_id = ?", id, tenantID).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NotFound("product not found")
		}
		return nil, exception.Internal("failed to find product", err)
	}
	return &product, nil
}

func (r *productRepositoryImpl) FindByTenantID(ctx context.Context, tenantID string, page, limit int) ([]entity.Product, int64, *exception.Exception) {
	var products []entity.Product
	var total int64

	if err := r.db.WithContext(ctx).Model(&entity.Product{}).Where("tenant_id = ?", tenantID).Count(&total).Error; err != nil {
		return nil, 0, exception.Internal("failed to count products", err)
	}

	if err := r.db.WithContext(ctx).Preload("Brand").Preload("Category").Preload("Tenant").Where("tenant_id = ?", tenantID).Offset((page - 1) * limit).Limit(limit).Find(&products).Error; err != nil {
		return nil, 0, exception.Internal("failed to find products", err)
	}

	return products, total, nil
}

func (r *productRepositoryImpl) Update(ctx context.Context, product *entity.Product) *exception.Exception {
    if err := r.db.WithContext(ctx).Save(product).Error; err != nil {
        return exception.Internal("failed to update product", err)
    }
    return nil
}

func (r *productRepositoryImpl) Delete(ctx context.Context, id string, tenantID string) *exception.Exception {
    if err := r.db.WithContext(ctx).Where("id = ? AND tenant_id = ?", id, tenantID).Delete(&entity.Product{}).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            return exception.NotFound("product not found")
        }
        return exception.Internal("failed to delete product", err)
    }
    return nil
}