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

	// User Routes
	router.POST("/users", userController.CreateUser)
	router.GET("/getAllUsers", userController.GetAllUsers)

	// Song Routes
	router.POST("/saveSong", songController.CreateSong)
	router.GET("/getAllSongs", songController.GetAllSongs)
	router.GET("/getSongByID/:id", songController.GetSongByID)
	router.PUT("/updateSong/:id", songController.UpdateSong)
	router.DELETE("/deleteSong/:id", songController.DeleteSong)

	// Playlist Routes
	router.GET("/playlists", playlistController.GetAllPlaylists)
	router.POST("/playlists", playlistController.CreatePlaylist)
	router.POST("/playlists/:id/songs", playlistController.AddSongToPlaylist)
	router.GET("/playlists/:id", playlistController.GetPlaylist)

	// Liked Songs Routes
	router.POST("/likedSongs/like", likedSongController.LikeSong)
	router.GET("/likedSongs/:userID", likedSongController.GetLikedSongs)

	// Start the server
	router.Run(":8080")
}
