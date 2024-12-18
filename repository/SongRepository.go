package repository

import (
	"gorm.io/gorm"
	"rest-api/models"
)

type SongRepository struct {
	DB *gorm.DB
}

// Constructor
func NewSongRepository(db *gorm.DB) *SongRepository {
	return &SongRepository{DB: db}
}

// CreateSong saves a new song to the database
func (r *SongRepository) CreateSong(song *models.Song) error {
	if err := r.DB.Create(song).Error; err != nil {
		return err
	}
	return nil
}

// GetAllSongs fetches all songs from the database
func (r *SongRepository) GetAllSongs() ([]models.Song, error) {
	var songs []models.Song
	if err := r.DB.Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

// GetSongByID fetches a single song by its ID
func (r *SongRepository) GetSongByID(id uint) (*models.Song, error) {
	var song models.Song
	if err := r.DB.First(&song, id).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

// UpdateSong updates an existing song
func (r *SongRepository) UpdateSong(song *models.Song) error {
	if err := r.DB.Save(song).Error; err != nil {
		return err
	}
	return nil
}

// DeleteSong deletes a song by its ID
func (r *SongRepository) DeleteSong(id uint) error {
	if err := r.DB.Delete(&models.Song{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *SongRepository) GetByID(id int) (*models.Song, error) {
	var song models.Song
	if err := r.DB.First(&song, id).Error; err != nil {
		return nil, err
	}
	return &song, nil
}
