package service

import (
	"context"

	"petshop-pos/internal/dto"
	"petshop-pos/pkg/exception"
	"petshop-pos/pkg/response"

	"github.com/google/uuid"
)

type TenantService interface {
    Create(ctx context.Context, request dto.CreateTenantRequest) (*dto.TenantResponse, *exception.Exception)
    GetByID(ctx context.Context, id uuid.UUID) (*dto.TenantResponse, *exception.Exception)
    GetAll(ctx context.Context, page, limit int) ([]dto.TenantResponse, *response.Metadata, *exception.Exception)
    Update(ctx context.Context, id uuid.UUID, request dto.UpdateTenantRequest) (*dto.TenantResponse, *exception.Exception)
    Delete(ctx context.Context, id uuid.UUID) *exception.Exception
}