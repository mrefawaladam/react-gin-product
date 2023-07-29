package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/pkg/handler"
)

func SetupProductRoutes(r *gin.Engine, productHandler handler.ProductHandler) {
	r.GET("/products/:id", productHandler.GetProductByID)
	r.POST("/products", productHandler.CreateProduct)
}
