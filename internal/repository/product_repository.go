package repository

import (
	"context"

	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"
)

type ProductRepository interface {
    Create(ctx context.Context, product *entity.Product) *exception.Exception
    FindByIDAndTenantID(ctx context.Context, id string, tenantID string) (*entity.Product, *exception.Exception)
    FindByTenantID(ctx context.Context, tenantID string, page, limit int) ([]entity.Product, int64, *exception.Exception)
    Update(ctx context.Context, product *entity.Product) *exception.Exception
    Delete(ctx context.Context, id string, tenantID string) *exception.Exception
}