package handler

import (
	"net/http"
	"petshop-pos/internal/dto"
	"petshop-pos/internal/service"
	"petshop-pos/pkg/response"

	"github.com/gin-gonic/gin"
)

type BrandHandler struct {
	service service.BrandService
}

func NewBrandHandler(service service.BrandService) *BrandHandler {
	return &BrandHandler{
		service: service,
	}
}

func (h *BrandHandler) CreateBrand(c *gin.Context) {
	// bind request
	var request dto.CreateBrandRequest
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

	c.JSON(http.StatusCreated, response.NewBaseSuccessResponse("Brand created successfully", nil))
}

func (h *BrandHandler) GetByID(c *gin.Context) {
	// get query parameter
	id := c.Param("id")

	// call service
	brand, ex := h.service.GetByID(c.Request.Context(), id, c.GetHeader("x-tenant-name"))
	if ex != nil {
		code, resp := response.MapExceptionToHTTP(ex)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Brand retrieved successfully", brand))
}

func (h *BrandHandler) GetAll(c *gin.Context) {
	// get pagination info
	page := c.Query("page")
	limit := c.Query("limit")

	// convert pagination info
	pageInt, limitInt := response.ValidateAndConvertPagination(page, limit)

	// call service
	brands, metadata, ex := h.service.GetAll(c.Request.Context(), pageInt, limitInt, c.GetHeader("x-tenant-name"))
	if ex != nil {
		code, resp := response.MapExceptionToHTTP(ex)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseMetadataSuccessResponse("Brands retrieved successfully", *metadata, brands))
}

func (h *BrandHandler) Update(c *gin.Context) {
	// bind request
	var request dto.UpdateBrandRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Invalid request body"))
		return
	}

	brandID := c.Param("id")

	// call service
	brand, ex := h.service.Update(c.Request.Context(), brandID, request, c.GetHeader("x-tenant-name"))
	if ex != nil {
		code, resp := response.MapExceptionToHTTP(ex)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Brand updated successfully", brand))
}

func (h *BrandHandler) Delete(c *gin.Context) {
	brandID := c.Param("id")

	// call service
	ex := h.service.Delete(c.Request.Context(), brandID, c.GetHeader("x-tenant-name"))
	if ex != nil {
		code, resp := response.MapExceptionToHTTP(ex)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Brand deleted successfully", nil))
}