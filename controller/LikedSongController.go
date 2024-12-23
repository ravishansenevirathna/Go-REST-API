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

func (c *LikedSongController) GetLikedSongs(ctx *gin.Context) {
	userIDStr := ctx.Param("userID")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	likedSongs, err := c.service.GetLikedSongsByUser(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, likedSongs)
}
