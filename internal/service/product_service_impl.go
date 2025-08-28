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

type ProductServiceImpl struct {
	repo     repository.ProductRepository
	validator *xvalidator.Validator
	tenantRepo repository.TenantRepository
}

func NewProductService(repo repository.ProductRepository, validator *xvalidator.Validator, tenantRepo repository.TenantRepository) ProductService {
	return &ProductServiceImpl{
		repo:     repo,
		validator: validator,
		tenantRepo: tenantRepo,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, tenantName string, request dto.CreateProductRequest) (*exception.Exception) {
	// validate request
	if err := s.validator.Struct(request); err != nil {
		return exception.InvalidArgument(err)
	}

	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return err
	}

	// to product entity
	product := &entity.Product{
		Name:     request.Name,
		Stock:   request.Stock,
		Price:    request.Price,
		TenantID: tenantID,
		BrandID: request.BrandID,
		CategoryID: request.CategoryID,
	}

	// call repository
	if err := s.repo.Create(ctx, product); err != nil {
		return nil
	}

	return nil
}

func (s *ProductServiceImpl) GetByIDAndTenantID(ctx context.Context, id string, tenantName string) (*dto.ProductResponse, *exception.Exception) {
	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// call repository
	product, err := s.repo.FindByIDAndTenantID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Stock:   product.Stock,
		Price:    product.Price,
		Tenant:   &dto.TenantResponse{
			ID:      product.Tenant.ID,
			Name:    product.Tenant.Name,
			Location: product.Tenant.Location,
		},
		Brand: &dto.BrandResponseWithoutTenant{
			ID:   product.Brand.ID,
			Name: product.Brand.Name,
		},
		Category: &dto.CategoryResponseWithoutTenant{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
		CreatedAt: time.FormatTimeToJakarta(product.CreatedAt).Format("2006-01-02 15:04:05"),
		UpdatedAt: time.FormatTimeToJakarta(product.UpdatedAt).Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *ProductServiceImpl) GetAllByTenantID(ctx context.Context, tenantName string, page, limit int) ([]dto.ProductResponse, *response.Metadata, *exception.Exception) {
	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, nil, err
	}

	// call repository
	products, total, err := s.repo.FindByTenantID(ctx, tenantID, page, limit)
	if err != nil {
		return nil, nil, err
	}

	// to product response
	var productResponses []dto.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, dto.ProductResponse{
			ID:       product.ID,
			Name:     product.Name,
			Stock:    product.Stock,
			Price:    product.Price,
			Tenant:   &dto.TenantResponse{
				ID:      product.Tenant.ID,
				Name:    product.Tenant.Name,
				Location: product.Tenant.Location,
			},
			Brand: &dto.BrandResponseWithoutTenant{
				ID:   product.Brand.ID,
				Name: product.Brand.Name,
			},
			Category: &dto.CategoryResponseWithoutTenant{
				ID:   product.Category.ID,
				Name: product.Category.Name,
			},
			CreatedAt: time.FormatTimeToJakarta(product.CreatedAt).Format("2006-01-02 15:04:05"),
			UpdatedAt: time.FormatTimeToJakarta(product.UpdatedAt).Format("2006-01-02 15:04:05"),
		})
	}

	// metadata
	metadata := &response.Metadata{
		Page:      page,
		Limit:    limit,
		Total:    int(total),
		TotalPages: response.CalculateTotalPages(int(total), limit),
	}

	return productResponses, metadata, nil
}

func (s *ProductServiceImpl) Update(ctx context.Context, id string, tenantName string, request dto.UpdateProductRequest) (*dto.ProductResponse, *exception.Exception) {
	// validate request
	if err := s.validator.Struct(request); err != nil {
		return nil, exception.InvalidArgument(err)
	}

	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return nil, err
	}

	// check if product exists
	product, err := s.repo.FindByIDAndTenantID(ctx, id, tenantID)
	if err != nil {
		return nil, err
	}

	// update product
	product.Name = request.Name
	product.Stock = request.Stock
	product.Price = request.Price
	product.BrandID = request.BrandID
	product.CategoryID = request.CategoryID

	// call repository
	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}

	return &dto.ProductResponse{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		Stock:    product.Stock,
		Tenant: &dto.TenantResponse{
			ID:      product.Tenant.ID,
			Name:    product.Tenant.Name,
			Location: product.Tenant.Location,
		},
		Brand: &dto.BrandResponseWithoutTenant{
			ID:   product.Brand.ID,
			Name: product.Brand.Name,
		},
		Category: &dto.CategoryResponseWithoutTenant{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		},
	}, nil
}

func (s *ProductServiceImpl) Delete(ctx context.Context, id string, tenantName string) *exception.Exception {
	// find tenant id
	tenantID, err := s.tenantRepo.FindIDByName(ctx, tenantName)
	if err != nil {
		return err
	}

	// check if product exists
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