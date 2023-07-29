package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/pkg/entity"
	"github.com/mrefawaladam/server/pkg/response"
	"github.com/mrefawaladam/server/pkg/usecase"
)

type ProductBrandHandler interface {
	GetProductBrandByID(c *gin.Context)
	CreateProductBrand(c *gin.Context)
}

type productBrandHandler struct {
	ProductBrandUsecase usecase.ProductBrandUsecase
}

func NewProductBrandHandler(productBrandUsecase usecase.ProductBrandUsecase) ProductBrandHandler {
	return &productBrandHandler{ProductBrandUsecase: productBrandUsecase}
}

func (h *productBrandHandler) GetProductBrandByID(c *gin.Context) {
	productBrandID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product brand ID"})
		return
	}

	productBrand, err := h.ProductBrandUsecase.GetProductBrandByID(uint(productBrandID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product brand not found"})
		return
	}

	productBrandResponse := response.NewProductBrandResponse(productBrand)
	c.JSON(http.StatusOK, productBrandResponse)
}

func (h *productBrandHandler) CreateProductBrand(c *gin.Context) {
	var productBrand response.ProductBrandResponse
	if err := c.ShouldBindJSON(&productBrand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newProductBrand := &entity.ProductBrand{
		Name: productBrand.Name,
	}

	if err := h.ProductBrandUsecase.CreateProductBrand(newProductBrand); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product brand"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product brand created successfully"})
}
