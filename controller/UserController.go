package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/dto"
	"rest-api/service"
)

// UserController handles HTTP requests for users.
type UserController struct {
	Service *service.UserService
}

// NewUserController creates a new UserController instance.
func NewUserController(service *service.UserService) *UserController {
	return &UserController{Service: service}
}

// CreateUser handles the POST /users route.
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the CreateUser method of UserService to create the user
	createdUser, err := c.Service.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the created user DTO as JSON response
	ctx.JSON(http.StatusCreated, createdUser)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {

	createdUser, err := c.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return the created user DTO as JSON response
	ctx.JSON(http.StatusCreated, createdUser)
}
