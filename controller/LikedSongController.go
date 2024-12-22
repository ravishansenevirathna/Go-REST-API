package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/dto"
	"rest-api/service"
	"strconv"
)

type LikedSongController struct {
	service *service.LikedSongService
}

func NewLikedSongController(service *service.LikedSongService) *LikedSongController {
	return &LikedSongController{service: service}
}

// Endpoint to like a song
func (c *LikedSongController) LikeSong(ctx *gin.Context) {
	var likeSongDto dto.LikeSongRequestDto
	if err := ctx.ShouldBindJSON(&likeSongDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.LikeSong(likeSongDto.UserID, likeSongDto.SongID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Song liked successfully"})
}

// Endpoint to get liked songs
func (c *LikedSongController) GetLikedSongs(ctx *gin.Context) {
	// Get userID as a string from the URL parameter
	userIDStr := ctx.Param("userID")

	// Convert userID to an integer
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		// Handle the case where the conversion fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the service method with the converted userID
	likedSongs, err := c.service.GetLikedSongsByUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the liked songs as a response
	ctx.JSON(http.StatusOK, likedSongs)
}
