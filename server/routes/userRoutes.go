package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/pkg/handler"
)

func SetupUserRoutes(r *gin.Engine, userHandler handler.UserHandler) {
	r.GET("/users/:id", userHandler.GetUserByID)
	r.POST("/users", userHandler.CreateUser)
}
