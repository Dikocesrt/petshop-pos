package dto

type CreateCategoryRequest struct {
    Name string `json:"name" validate:"required,min=1,max=255"`
}

type UpdateCategoryRequest struct {
    Name string `json:"name" validate:"required,min=1,max=255"`
}

type CategoryResponse struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
    Tenant    TenantResponse `json:"tenant,omitempty"`
}

type CategoryResponseWithoutTenant struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    CreatedAt string `json:"createdAt,omitempty"`
    UpdatedAt string `json:"updatedAt,omitempty"`
}