package repository

import (
	"context"

	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"
)

type BrandRepository interface {
    Create(ctx context.Context, brand *entity.Brand) *exception.Exception
    FindByIDAndTenantID(ctx context.Context, id string, tenantID string) (*entity.Brand, *exception.Exception)
    FindAllByTenantID(ctx context.Context, tenantID string, page, limit int) ([]entity.Brand, int64, *exception.Exception)
    Update(ctx context.Context, brand *entity.Brand) *exception.Exception
    Delete(ctx context.Context, id string, tenantID string) *exception.Exception
}