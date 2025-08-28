package service

import (
	"context"
	"petshop-pos/internal/dto"
	"petshop-pos/internal/entity"
	"petshop-pos/internal/repository"
	"petshop-pos/pkg/exception"
	"petshop-pos/pkg/response"
	"petshop-pos/pkg/time"
	"petshop-pos/pkg/xvalidator"
)

type BrandServiceImpl struct {
	repo repository.BrandRepository
	validator *xvalidator.Validator
	tenantRepo repository.TenantRepository
}

func NewBrandService(repo repository.BrandRepository, validator *xvalidator.Validator, tenantRepo repository.TenantRepository) BrandService {
	return &BrandServiceImpl{
		repo: repo,
		validator: validator,
		tenantRepo: tenantRepo,
	}
}

func (s *BrandServiceImpl) Create(ctx context.Context, request dto.CreateBrandRequest, tenantName string) (*dto.BrandResponse, *exception.Exception) {
	// validate request
	if err := s.validator.Struct(request); err != nil {
		return nil, exception.InvalidArgument(err)
	}

	// find tenant ID
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// to brand entity
	brand := &entity.Brand{
		Name:     request.Name,
		TenantID: tenantID,
	}

	// call repository
	if err := s.repo.Create(ctx, brand); err != nil {
		return nil, err
	}

	return &dto.BrandResponse{
		ID:   brand.ID,
		Name: brand.Name,
		Tenant: dto.TenantResponse{
			ID:      brand.Tenant.ID,
			Name:    brand.Tenant.Name,
			Location: brand.Tenant.Location,
		},
	}, nil
}

func (s *BrandServiceImpl) GetByID(ctx context.Context, id string, tenantName string) (*dto.BrandResponse, *exception.Exception) {
	// find tenant ID
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// call repository
	brand, err := s.repo.FindByIDAndTenantID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	return &dto.BrandResponse{
		ID:   brand.ID,
		Name: brand.Name,
		Tenant: dto.TenantResponse{
			ID:      brand.Tenant.ID,
			Name:    brand.Tenant.Name,
			Location: brand.Tenant.Location,
		},
		CreatedAt: time.FormatTimeToJakarta(brand.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.FormatTimeToJakarta(brand.UpdatedAt).Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *BrandServiceImpl) GetAll(ctx context.Context, page, limit int, tenantName string) ([]dto.BrandResponse, *response.Metadata, *exception.Exception) {
	// find tenant ID
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, nil, err
	}

	// call repository
	brands, total, err := s.repo.FindAllByTenantID(ctx, tenantID, page, limit)
	if err != nil {
		return nil, nil, err
	}

	// to brand response
	var brandResponses []dto.BrandResponse
	for _, brand := range brands {
		brandResponses = append(brandResponses, dto.BrandResponse{
			ID:   brand.ID,
			Name: brand.Name,
			Tenant: dto.TenantResponse{
				ID:   brand.Tenant.ID,
				Name: brand.Tenant.Name,
				Location: brand.Tenant.Location,
			},
			CreatedAt: time.FormatTimeToJakarta(brand.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.FormatTimeToJakarta(brand.UpdatedAt).Format("2006-01-02 15:04:05"),
		})
	}

	// metadata
	metadata := &response.Metadata{
		Page:  page,
		Limit: limit,
		Total: int(total),
		TotalPages: response.CalculateTotalPages(int(total), limit),
	}

	return brandResponses, metadata, nil
}

func (s *BrandServiceImpl) Update(ctx context.Context, id string, request dto.UpdateBrandRequest, tenantName string) (*dto.BrandResponse, *exception.Exception) {
	// validate request
	if err := s.validator.Struct(request); err != nil {
		return nil, exception.InvalidArgument(err)
	}

	// find tenant ID
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// check if brand exists
	brand, err := s.repo.FindByIDAndTenantID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	// update brand
	brand.Name = request.Name

	// call repository
	if err := s.repo.Update(ctx, brand); err != nil {
		return nil, err
	}

	return &dto.BrandResponse{
		ID:   brand.ID,
		Name: brand.Name,
		Tenant: dto.TenantResponse{
			ID:      brand.Tenant.ID,
			Name:    brand.Tenant.Name,
			Location: brand.Tenant.Location,
		},
		CreatedAt: time.FormatTimeToJakarta(brand.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.FormatTimeToJakarta(brand.UpdatedAt).Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *BrandServiceImpl) Delete(ctx context.Context, id string, tenantName string) *exception.Exception {
	// find tenant ID
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return err
	}

	// check if brand exists
	_, err = s.repo.FindByIDAndTenantID(ctx, id, tenantID)
	if err != nil {
		return err
	}

	// call repository
	if err := s.repo.Delete(ctx, id, tenantID); err != nil {
		return err
	}
	return nil
}