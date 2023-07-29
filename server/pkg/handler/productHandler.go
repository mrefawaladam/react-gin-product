package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/pkg/entity"
	"github.com/mrefawaladam/server/pkg/response"
	"github.com/mrefawaladam/server/pkg/usecase"
)

type ProductHandler interface {
	GetProductByID(c *gin.Context)
	CreateProduct(c *gin.Context)
}

type productHandler struct {
	ProductUsecase      usecase.ProductUsecase
	ProductBrandUsecase usecase.ProductBrandUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase, productBrandUsecase usecase.ProductBrandUsecase) ProductHandler {
	return &productHandler{
		ProductUsecase:      productUsecase,
		ProductBrandUsecase: productBrandUsecase,
	}
}

func (h *productHandler) GetProductByID(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.ProductUsecase.GetProductByID(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	productResponse := response.NewProductResponse(product)
	c.JSON(http.StatusOK, productResponse)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var product response.ProductResponse
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	productBrand, err := h.ProductBrandUsecase.GetProductBrandByID(product.ProductBrandID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product brand ID"})
		return
	}

	newProduct := &entity.Product{
		Name:           product.Name,
		ProductBrandID: product.ProductBrandID,
		ProductBrand:   *productBrand,
	}

	if err := h.ProductUsecase.CreateProduct(newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
}
