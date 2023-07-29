package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/pkg/handler"
	"github.com/mrefawaladam/server/pkg/repository"
	"github.com/mrefawaladam/server/pkg/usecase"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	productBrandRepo := repository.NewProductBrandRepository(db)
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)

	productBrandUsecase := usecase.NewProductBrandUsecase(*productBrandRepo)
	userUsecase := usecase.NewUserUsecase(*userRepo)
	productrUsecase := usecase.NewProductUsecase(*productRepo, *productBrandRepo)

	productBrandHandler := handler.NewProductBrandHandler(productBrandUsecase)
	userHandler := handler.NewUserHandler(userUsecase)
	productHandler := handler.NewProductHandler(productrUsecase, productBrandUsecase)

	SetupProductBrandRoutes(r, productBrandHandler)
	SetupUserRoutes(r, userHandler)
	SetupProductRoutes(r, productHandler)

}
