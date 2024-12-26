package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"rest-api/Auth"
	"rest-api/controller"
	"rest-api/models"
	"rest-api/repository"
	"rest-api/service"
)

func main() {
	// Database DSN (Data Source Name)
	dsn := "root:1234@tcp(127.0.0.1:3306)/rest_api?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}

	// Auto-migrate the models
	err = db.AutoMigrate(&models.User{}, &models.Song{}, &models.Playlist{}, &models.LikedSong{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize User components
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Initialize Song components
	songRepo := repository.NewSongRepository(db)
	songService := service.NewSongService(songRepo)
	songController := controller.NewSongController(songService)

	// Initialize Playlist components
	playlistRepo := repository.NewPlaylistRepository(db)
	playlistService := service.NewPlaylistService(playlistRepo, songRepo)
	playlistController := controller.NewPlaylistController(playlistService)

	// Initialize Liked Songs components
	likedSongRepo := repository.NewLikedSongRepository(db)
	likedSongService := service.NewLikedSongService(likedSongRepo)
	likedSongController := controller.NewLikedSongController(likedSongService)

	// Set up Gin routes
	router := gin.Default()

	// Secure User Routes
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/getAllUsers", Auth.JWTAuthMiddleware(), userController.GetAllUsers)
		userRoutes.POST("/logIn", userController.LogIn)       // No auth required for login
		userRoutes.POST("/signIn", userController.CreateUser) // No auth required for sign-up
	}

	// Secure Song Routes
	songRoutes := router.Group("/songs")
	{
		songRoutes.POST("/saveSong", Auth.JWTAuthMiddleware(), songController.CreateSong)
		songRoutes.GET("/getAllSongs", Auth.JWTAuthMiddleware(), songController.GetAllSongs)
		songRoutes.GET("/getSongByID/:id", Auth.JWTAuthMiddleware(), songController.GetSongByID)
		songRoutes.PUT("/updateSong/:id", Auth.JWTAuthMiddleware(), songController.UpdateSong)
		songRoutes.DELETE("/deleteSong/:id", Auth.JWTAuthMiddleware(), songController.DeleteSong)
	}

	playlistRoutes := router.Group("/playlists")
	{
		playlistRoutes.GET("", Auth.JWTAuthMiddleware(), playlistController.GetAllPlaylists)
		playlistRoutes.POST("", Auth.JWTAuthMiddleware(), playlistController.CreatePlaylist)
		playlistRoutes.POST("/:id/songs", Auth.JWTAuthMiddleware(), playlistController.AddSongToPlaylist)
		playlistRoutes.GET("/:id", Auth.JWTAuthMiddleware(), playlistController.GetPlaylist)
	}

	// Secure Liked Songs Routes
	likedSongsRoutes := router.Group("/likedSongs")
	{
		likedSongsRoutes.POST("/like", Auth.JWTAuthMiddleware(), likedSongController.LikeSong)
		likedSongsRoutes.GET("/:userID", Auth.JWTAuthMiddleware(), likedSongController.GetLikedSongs)
	}

	// Start the server
	router.Run(":8080")
}
