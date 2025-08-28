package dto

type CreateProductRequest struct {
    Name       string     `json:"name" validate:"required,min=1,max=255"`
    Stock      int        `json:"stock" validate:"required,min=0"`
    Price      int        `json:"price" validate:"required,min=0"`
    BrandID    *string     `json:"brandID"`
    CategoryID *string     `json:"categoryID"`
}

type UpdateProductRequest struct {
    Name       string     `json:"name" validate:"required,min=1,max=255"`
    Stock      int        `json:"stock" validate:"required,min=0"`
    Price      int        `json:"price" validate:"required,min=0"`
    BrandID    *string   `json:"brandID"`
    CategoryID *string   `json:"categoryID"`
}

type ProductResponse struct {
    ID         string        `json:"id"`
    Name       string        `json:"name"`
    Stock      int           `json:"stock"`
    Price      int           `json:"price"`
    CreatedAt  string          `json:"created_at,omitempty"`
    UpdatedAt  string          `json:"updated_at,omitempty"`
    Brand      *BrandResponseWithoutTenant  `json:"brand,omitempty"`
    Category   *CategoryResponseWithoutTenant `json:"category,omitempty"`
    Tenant     *TenantResponse   `json:"tenant,omitempty"`
}