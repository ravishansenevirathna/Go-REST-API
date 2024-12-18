package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/models"
	"rest-api/service"
	"strconv"
)

type SongController struct {
	service *service.SongService
}

// Constructor
func NewSongController(service *service.SongService) *SongController {
	return &SongController{service: service}
}

// CreateSong - Handles creating a song
func (c *SongController) CreateSong(ctx *gin.Context) {
	var song models.Song
	if err := ctx.ShouldBindJSON(&song); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateSong(&song); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create song"})
		return
	}

	ctx.JSON(http.StatusCreated, song)
}

// GetAllSongs - Fetch all songs
func (c *SongController) GetAllSongs(ctx *gin.Context) {
	songs, err := c.service.GetAllSongs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch songs"})
		return
	}
	ctx.JSON(http.StatusOK, songs)
}

// GetSongByID - Fetch a song by ID
func (c *SongController) GetSongByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	song, err := c.service.GetSongByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}
	ctx.JSON(http.StatusOK, song)
}

// UpdateSong - Update a song
func (c *SongController) UpdateSong(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var song models.Song
	if err := ctx.ShouldBindJSON(&song); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	song.ID = int(uint(id)) // Ensure correct ID is set
	if err := c.service.UpdateSong(&song); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update song"})
		return
	}
	ctx.JSON(http.StatusOK, song)
}

// DeleteSong - Delete a song by ID
func (c *SongController) DeleteSong(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.service.DeleteSong(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete song"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
}
