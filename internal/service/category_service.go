package service

import (
	"context"

	"petshop-pos/internal/dto"
	"petshop-pos/pkg/exception"
	"petshop-pos/pkg/response"
)

type CategoryService interface {
    Create(ctx context.Context, request dto.CreateCategoryRequest, tenantName string) (*dto.CategoryResponse, *exception.Exception)
    GetByID(ctx context.Context, id string, tenantName string) (*dto.CategoryResponse, *exception.Exception)
    GetAll(ctx context.Context, page, limit int, tenantName string) ([]dto.CategoryResponse, *response.Metadata, *exception.Exception)
    Update(ctx context.Context, id string, request dto.UpdateCategoryRequest, tenantName string) (*dto.CategoryResponse, *exception.Exception)
    Delete(ctx context.Context, id string, tenantName string) *exception.Exception
}