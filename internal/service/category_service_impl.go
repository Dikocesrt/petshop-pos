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

type CategoryServiceImpl struct {
	repo     repository.CategoryRepository
	validator *xvalidator.Validator
	tenantRepo repository.TenantRepository
}

func NewCategoryService(repo repository.CategoryRepository, validator *xvalidator.Validator, tenantRepo repository.TenantRepository) CategoryService {
	return &CategoryServiceImpl{
		repo:     repo,
		validator: validator,
		tenantRepo: tenantRepo,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request dto.CreateCategoryRequest, tenantName string) (*dto.CategoryResponse, *exception.Exception) {
	// validate request
	if err := s.validator.Struct(request); err != nil {
		return nil, exception.InvalidArgument(err)
	}

	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// to category entity
	category := &entity.Category{
		Name:     request.Name,
		TenantID: tenantID,
	}

	// call repository
	if err := s.repo.Create(ctx, category); err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
		CreatedAt: time.FormatTimeToJakarta(category.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.FormatTimeToJakarta(category.UpdatedAt).Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *CategoryServiceImpl) GetByID(ctx context.Context, id string, tenantName string) (*dto.CategoryResponse, *exception.Exception) {
	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// call repository
	category, err := s.repo.FindByIDAndTenantID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
		Tenant: dto.TenantResponse{
			ID:      category.Tenant.ID,
			Name:    category.Tenant.Name,
			Location: category.Tenant.Location,
		},
		CreatedAt: time.FormatTimeToJakarta(category.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.FormatTimeToJakarta(category.UpdatedAt).Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *CategoryServiceImpl) GetAll(ctx context.Context, page, limit int, tenantName string) ([]dto.CategoryResponse, *response.Metadata, *exception.Exception) {
	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, nil, err
	}

	// call repository
	categories, total, err := s.repo.FindAllByTenantID(ctx, page, limit, tenantID)
	if err != nil {
		return nil, nil, err
	}

	// to category response
	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
			Tenant: dto.TenantResponse{
				ID:      category.Tenant.ID,
				Name:    category.Tenant.Name,
				Location: category.Tenant.Location,
			},
			CreatedAt: time.FormatTimeToJakarta(category.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.FormatTimeToJakarta(category.UpdatedAt).Format("2006-01-02 15:04:05"),
		})
	}

	// metadata
	metadata := &response.Metadata{
		Page:      page,
		Limit:    limit,
		Total:    int(total),
		TotalPages: response.CalculateTotalPages(int(total), limit),
	}

	return categoryResponses, metadata, nil
}

func (s *CategoryServiceImpl) Update(ctx context.Context, id string, request dto.UpdateCategoryRequest, tenantName string) (*dto.CategoryResponse, *exception.Exception) {
	// validate request
	if err := s.validator.Struct(request); err != nil {
		return nil, exception.InvalidArgument(err)
	}

	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// call repository
	category, err := s.repo.FindByIDAndTenantID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	// update category
	category.Name = request.Name

	// call repository
	if err := s.repo.Update(ctx, category); err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
		Tenant: dto.TenantResponse{
			ID:      category.Tenant.ID,
			Name:    category.Tenant.Name,
			Location: category.Tenant.Location,
		},
		CreatedAt: time.FormatTimeToJakarta(category.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.FormatTimeToJakarta(category.UpdatedAt).Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, id string, tenantName string) *exception.Exception {
	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return err
	}

	// call repository
	if err := s.repo.Delete(ctx, id, tenantID); err != nil {
		return err
	}
	return nil
}