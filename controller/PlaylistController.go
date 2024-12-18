package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/dto"
	"rest-api/models"
	"rest-api/service"
	"strconv"
)

type PlaylistController struct {
	service *service.PlaylistService
}

func NewPlaylistController(service *service.PlaylistService) *PlaylistController {
	return &PlaylistController{service: service}
}

// GetAllPlaylists retrieves all playlists
func (c *PlaylistController) GetAllPlaylists(ctx *gin.Context) {
	playlists, err := c.service.GetAllPlaylists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve playlists"})
		return
	}

	ctx.JSON(http.StatusOK, playlists)
}

// CreatePlaylist creates a new playlist
func (c *PlaylistController) CreatePlaylist(ctx *gin.Context) {
	var request dto.CreatePlaylistRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	playlists, err := c.service.GetAllPlaylists()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve playlists"})
		return
	}

	playlist := models.Playlist{
		ID:    len(playlists) + 1,
		Name:  request.Name,
		Songs: []models.Song{}, // Initialize an empty song list
	}

	c.service.CreatePlaylist(playlist)
	ctx.JSON(http.StatusCreated, playlist)
}

// AddSongToPlaylist adds a song to a playlist
func (c *PlaylistController) AddSongToPlaylist(ctx *gin.Context) {
	playlistID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	var request dto.AddSongRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.service.AddSongToPlaylist(playlistID, request.SongID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Song added successfully"})
}

// GetPlaylist retrieves a single playlist by ID
func (c *PlaylistController) GetPlaylist(ctx *gin.Context) {
	playlistID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid playlist ID"})
		return
	}

	playlist, err := c.service.GetPlaylistByID(playlistID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, playlist)
}
