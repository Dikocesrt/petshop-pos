package service

import (
	"context"

	"petshop-pos/internal/dto"
	"petshop-pos/pkg/exception"
	"petshop-pos/pkg/response"
)

type ProductService interface {
    Create(ctx context.Context, tenantName string, request dto.CreateProductRequest) (*exception.Exception)
	GetByIDAndTenantID(ctx context.Context, id string, tenantName string) (*dto.ProductResponse, *exception.Exception)
	GetAllByTenantID(ctx context.Context, tenantName string, page, limit int) ([]dto.ProductResponse, *response.Metadata, *exception.Exception)
	Update(ctx context.Context, id string, tenantName string, request dto.UpdateProductRequest) (*dto.ProductResponse, *exception.Exception)
	Delete(ctx context.Context, id string, tenantName string) *exception.Exception
}