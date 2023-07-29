package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/pkg/handler"
)

func SetupProductBrandRoutes(r *gin.Engine, productBrandHandler handler.ProductBrandHandler) {
	r.GET("/product-brands/:id", productBrandHandler.GetProductBrandByID)
	r.POST("/product-brands", productBrandHandler.CreateProductBrand)
}
