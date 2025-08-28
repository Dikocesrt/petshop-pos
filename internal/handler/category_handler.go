package handler

import (
	"net/http"
	"petshop-pos/internal/dto"
	"petshop-pos/internal/service"
	"petshop-pos/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
    service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
    return &CategoryHandler{
        service: service,
    }
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
    // bind request
    var request dto.CreateCategoryRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Invalid request body"))
        return
    }

    // call service
    _, err := h.service.Create(c.Request.Context(), request, c.GetHeader("x-tenant-name"))
    if err != nil {
        code, resp := response.MapExceptionToHTTP(err)
        c.JSON(code, resp)
        return
    }

    c.JSON(http.StatusCreated, response.NewBaseSuccessResponse("Category created successfully", nil))
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
    // get query parameter
    id := c.Param("id")

    // call service
    category, ex := h.service.GetByID(c.Request.Context(), id, c.GetHeader("x-tenant-name"))
    if ex != nil {
        code, resp := response.MapExceptionToHTTP(ex)
        c.JSON(code, resp)
        return
    }

    c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Category retrieved successfully", category))
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
    // get pagination info
    page := c.Query("page")
    limit := c.Query("limit")

    // convert pagination info
    pageInt, limitInt := response.ValidateAndConvertPagination(page, limit)

    // call service
    categories, metadata, ex := h.service.GetAll(c.Request.Context(), pageInt, limitInt, c.GetHeader("x-tenant-name"))
    if ex != nil {
        code, resp := response.MapExceptionToHTTP(ex)
        c.JSON(code, resp)
        return
    }

    c.JSON(http.StatusOK, response.NewBaseMetadataSuccessResponse("Categories retrieved successfully", *metadata, categories))
}

func (h *CategoryHandler) Update(c *gin.Context) {
    // bind request
    var request dto.UpdateCategoryRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Invalid request body"))
        return
    }

    categoryID := c.Param("id")

    // call service
    category, ex := h.service.Update(c.Request.Context(), categoryID, request, c.GetHeader("x-tenant-name"))
    if ex != nil {
        code, resp := response.MapExceptionToHTTP(ex)
        c.JSON(code, resp)
        return
    }

    c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Category updated successfully", category))
}

func (h *CategoryHandler) Delete(c *gin.Context) {
    categoryID := c.Param("id")

    // call service
    ex := h.service.Delete(c.Request.Context(), categoryID, c.GetHeader("x-tenant-name"))
    if ex != nil {
        code, resp := response.MapExceptionToHTTP(ex)
        c.JSON(code, resp)
        return
    }

    c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Category deleted successfully", nil))
}