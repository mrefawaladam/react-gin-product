package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrefawaladam/server/pkg/entity"
	"github.com/mrefawaladam/server/pkg/response"
	"github.com/mrefawaladam/server/pkg/usecase"
)

type UserHandler interface {
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{UserUsecase: userUsecase}
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.UserUsecase.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userResponse := response.NewUserResponse(user)
	c.JSON(http.StatusOK, userResponse)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var user response.UserResponse
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newUser := &entity.User{
		Name:  user.Name,
		Email: user.Email,
	}

	if err := h.UserUsecase.CreateUser(newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
