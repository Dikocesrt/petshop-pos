package handler

import (
	"net/http"
	"petshop-pos/internal/dto"
	"petshop-pos/internal/service"
	"petshop-pos/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) Create(c *gin.Context) {
	// bind request
	var request dto.CreateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Invalid request body"))
		return
	}

	// call service
	if err := h.service.Create(c.Request.Context(), c.GetHeader("x-tenant-name"), request); err != nil {
		code, resp := response.MapExceptionToHTTP(err)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusCreated, response.NewBaseSuccessResponse("Product created successfully", nil))
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	// call service
	product, err := h.service.GetByIDAndTenantID(c.Request.Context(), c.Param("id"), c.GetHeader("x-tenant-name"))
	if err != nil {
		code, resp := response.MapExceptionToHTTP(err)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Product retrieved successfully", product))
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	// query params
	limit, page := c.Query("limit"), c.Query("page")

	// convert string params to int
	pageInt, limitInt := response.ValidateAndConvertPagination(page, limit)

	// call service
	products, metadata, exception := h.service.GetAllByTenantID(c.Request.Context(), c.GetHeader("x-tenant-name"), pageInt, limitInt)
	if exception != nil {
		code, resp := response.MapExceptionToHTTP(exception)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseMetadataSuccessResponse("Products retrieved successfully", *metadata, products))
}

func (h *ProductHandler) Update(c *gin.Context) {
	// bind request
	var request dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBaseErrorResponse("Invalid request body"))
		return
	}

	// call service
	products, err := h.service.Update(c.Request.Context(), c.Param("id"), c.GetHeader("x-tenant-name"), request);
	if err != nil {
		code, resp := response.MapExceptionToHTTP(err)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Product updated successfully", products))
}

func (h *ProductHandler) Delete(c *gin.Context) {
	// call service
	ex := h.service.Delete(c.Request.Context(), c.Param("id"), c.GetHeader("x-tenant-name"))
	if ex != nil {
		code, resp := response.MapExceptionToHTTP(ex)
		c.JSON(code, resp)
		return
	}

	c.JSON(http.StatusOK, response.NewBaseSuccessResponse("Product deleted successfully", nil))
}