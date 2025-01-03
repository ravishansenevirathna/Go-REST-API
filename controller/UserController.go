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

// SignUp handles the POST /signup route.
//func (c *UserController) SignUp(ctx *gin.Context) {
//	var user dto.User
//	if err := ctx.ShouldBindJSON(&user); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Hash the user's password
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
//		return
//	}
//	user.Password = string(hashedPassword)
//
//	// Save the user
//	_, err = c.Service.CreateUser(&user)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
//		return
//	}
//
//	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
//}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := c.Service.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (c *UserController) LogIn(ctx *gin.Context) {
	var user dto.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authenticatedUser, token, err := c.Service.LoginUser(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  authenticatedUser,
		"token": token,
	})
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
