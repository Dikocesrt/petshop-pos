package dto

type CreateBrandRequest struct {
    Name string `json:"name" validate:"required,min=1,max=255"`
}

type UpdateBrandRequest struct {
    Name string `json:"name" validate:"required,min=1,max=255"`
}

type BrandResponse struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    CreatedAt string    `json:"createdAt,omitempty"`
    UpdatedAt string    `json:"updatedAt,omitempty"`
    Tenant    TenantResponse `json:"tenant,omitempty"`
}

type BrandResponseWithoutTenant struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    CreatedAt string    `json:"createdAt,omitempty"`
    UpdatedAt string    `json:"updatedAt,omitempty"`
}