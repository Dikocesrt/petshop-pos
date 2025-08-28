package repository

import (
	"context"

	"petshop-pos/internal/entity"
	"petshop-pos/pkg/exception"
)

type CategoryRepository interface {
    Create(ctx context.Context, category *entity.Category) *exception.Exception
    FindByIDAndTenantID(ctx context.Context, id string, tenantID string) (*entity.Category, *exception.Exception)
    FindAllByTenantID(ctx context.Context, page, limit int, tenantID string) ([]entity.Category, int64, *exception.Exception)
    Update(ctx context.Context, category *entity.Category) *exception.Exception
    Delete(ctx context.Context, id string, tenantID string) *exception.Exception
}