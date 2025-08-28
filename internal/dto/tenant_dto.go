package dto

type CreateTenantRequest struct {
    Name     string `json:"name" validate:"required,min=1,max=255"`
    Location string `json:"location" validate:"required,min=1"`
    Phone    string `json:"phone" validate:"required,min=1"`
}

type UpdateTenantRequest struct {
    Name    string `json:"name" validate:"required,min=1,max=255"`
    Location string `json:"location" validate:"required,min=1"`
    Phone   string `json:"phone" validate:"required,min=1"`
}

type TenantResponse struct {
    ID        string `json:"id"`
    Name      string    `json:"name"`
    Location  string    `json:"location"`
    CreatedAt string    `json:"createdAt,omitempty"`
    UpdatedAt string    `json:"updatedAt,omitempty"`
}