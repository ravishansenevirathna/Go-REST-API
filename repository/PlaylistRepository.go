package repository

import (
	"gorm.io/gorm"
	"rest-api/models"
)

// PlaylistRepository handles database operations for playlists
type PlaylistRepository struct {
	DB *gorm.DB
}

// NewPlaylistRepository creates a new instance of PlaylistRepository
func NewPlaylistRepository(db *gorm.DB) *PlaylistRepository {
	return &PlaylistRepository{DB: db}
}

// Create saves a new playlist to the database
func (r *PlaylistRepository) Create(playlist *models.Playlist) error {
	if err := r.DB.Create(playlist).Error; err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all playlists from the database
func (r *PlaylistRepository) GetAll() ([]models.Playlist, error) {
	var playlists []models.Playlist
	if err := r.DB.Preload("Songs").Find(&playlists).Error; err != nil { // Preload songs for each playlist
		return nil, err
	}
	return playlists, nil
}

// GetByID retrieves a single playlist by its ID
func (r *PlaylistRepository) GetByID(id int) (*models.Playlist, error) {
	var playlist models.Playlist
	if err := r.DB.Preload("Songs").First(&playlist, id).Error; err != nil {
		return nil, err
	}
	return &playlist, nil
}

// AddSongToPlaylist adds a song to a playlist by updating its Songs association
func (r *PlaylistRepository) AddSongToPlaylist(playlistID int, song models.Song) error {
	var playlist models.Playlist
	if err := r.DB.First(&playlist, playlistID).Error; err != nil {
		return err
	}

	// Append the song to the playlist's Songs slice
	if err := r.DB.Model(&playlist).Association("Songs").Append(&song); err != nil {
		return err
	}

	return nil
}
