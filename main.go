package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"rest-api/controller"
	"rest-api/models"
	"rest-api/repository"
	"rest-api/service"
)

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/rest_api?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}

	// Auto-migrate the User table
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate User table: %v", err)
	}

	// Initialize components
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Set up Gin routes
	r := gin.Default()
	r.POST("/users", userController.CreateUser)

	// Start the server
	r.Run(":8080")
}
