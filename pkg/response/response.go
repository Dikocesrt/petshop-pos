package response

import (
	"net/http"
	"petshop-pos/pkg/exception"
	"strconv"
)

// BaseSuccessResponse is the standard response for successful requests.
type BaseSuccessResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
    Data    any    `json:"data,omitempty"`
}

// BaseMetadataSuccessResponse is the standard response for successful requests with pagination metadata.
type BaseMetadataSuccessResponse struct {
    Status   string   `json:"status"`
    Message  string   `json:"message"`
    Metadata Metadata `json:"metadata"`
    Data     any      `json:"data,omitempty"`
}

// BaseErrorResponse is the standard response for failed requests.
type BaseErrorResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

// Metadata holds pagination information.
type Metadata struct {
    Page  int `json:"page"`
    Limit int `json:"limit"`
    Total int `json:"total"`
    TotalPages int `json:"totalPages"`
}

// GetOffset returns the offset for pagination.
func (m Metadata) GetOffset() int {
    return (m.Page - 1) * m.Limit
}

// GetMetadata parses page and limit from string to Metadata struct.
func GetMetadata(page, limit string) Metadata {
    pageInt, err := strconv.Atoi(page)
    if err != nil || pageInt <= 0 {
        pageInt = 1
    }
    limitInt, err := strconv.Atoi(limit)
    if err != nil || limitInt <= 0 {
        limitInt = 10
    }
    return Metadata{Page: pageInt, Limit: limitInt}
}

// NewBaseSuccessResponse creates a new BaseSuccessResponse.
func NewBaseSuccessResponse(message string, data any) BaseSuccessResponse {
    return BaseSuccessResponse{
        Status:  "success",
        Message: message,
        Data:    data,
    }
}

// NewBaseMetadataSuccessResponse creates a new BaseMetadataSuccessResponse.
func NewBaseMetadataSuccessResponse(message string, metadata Metadata, data any) BaseMetadataSuccessResponse {
    return BaseMetadataSuccessResponse{
        Status:   "success",
        Message:  message,
        Metadata: metadata,
        Data:     data,
    }
}

// NewBaseErrorResponse creates a new BaseErrorResponse.
func NewBaseErrorResponse(message string) BaseErrorResponse {
    return BaseErrorResponse{
        Status:  "fail",
        Message: message,
    }
}

// NewErrorResponseFromException creates a BaseErrorResponse from an Exception.
func NewErrorResponseFromException(e *exception.Exception) BaseErrorResponse {
    return BaseErrorResponse{
        Status:  "fail",
        Message: e.Message.(string),
    }
}

func CalculateTotalPages(total, limit int) int {
    if limit <= 0 {
        return 0
    }
    
    if total == 0 {
        return 0
    }
    
    totalPages := total / limit
    if total%limit > 0 {
        totalPages++
    }
    
    return totalPages
}

func ValidateAndConvertPagination(page, limit string) (int, int) {
    pageInt, err := strconv.Atoi(page)
    if err != nil || pageInt <= 0 {
        pageInt = 1
    }

    limitInt, err := strconv.Atoi(limit)
    if err != nil || limitInt <= 0 {
        limitInt = 10
    }

    // Optional: Set maximum limit to prevent abuse
    if limitInt > 100 {
        limitInt = 100
    }

    return pageInt, limitInt
}

func MapExceptionToHTTP(e *exception.Exception) (int, BaseErrorResponse) {
    switch e.Code {
    case exception.InvalidArgumentCode:
        return http.StatusBadRequest, NewBaseErrorResponse(e.Message.(string))
    case exception.NotFoundCode:
        return http.StatusNotFound, NewBaseErrorResponse(e.Message.(string))
    case exception.AlreadyExistsCode:
        return http.StatusConflict, NewBaseErrorResponse(e.Message.(string))
    case exception.PermissionDeniedCode, exception.UnauthenticatedCode:
        return http.StatusUnauthorized, NewBaseErrorResponse(e.Message.(string))
    case exception.UnsupportedMediaTypeCode:
        return http.StatusUnsupportedMediaType, NewBaseErrorResponse(e.Message.(string))
    default:
        return http.StatusInternalServerError, NewBaseErrorResponse(e.Message.(string))
    }
}