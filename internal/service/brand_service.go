package service

import (
	"context"

	"petshop-pos/internal/dto"
	"petshop-pos/pkg/exception"
	"petshop-pos/pkg/response"
)

type BrandService interface {
    Create(ctx context.Context, request dto.CreateBrandRequest, tenantName string) (*dto.BrandResponse, *exception.Exception)
    GetByID(ctx context.Context, id string, tenantName string) (*dto.BrandResponse, *exception.Exception)
    GetAll(ctx context.Context, page, limit int, tenantName string) ([]dto.BrandResponse, *response.Metadata, *exception.Exception)
    Update(ctx context.Context, id string, request dto.UpdateBrandRequest, tenantName string) (*dto.BrandResponse, *exception.Exception)
    Delete(ctx context.Context, id string, tenantName string) *exception.Exception
}