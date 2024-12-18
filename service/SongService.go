package service

import (
	"rest-api/models"
	"rest-api/repository"
)

type SongService struct {
	repo *repository.SongRepository
}

// Constructor
func NewSongService(repo *repository.SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) CreateSong(song *models.Song) error {
	return s.repo.CreateSong(song)
}

func (s *SongService) GetAllSongs() ([]models.Song, error) {
	return s.repo.GetAllSongs()
}

func (s *SongService) GetSongByID(id uint) (*models.Song, error) {
	return s.repo.GetSongByID(id)
}

func (s *SongService) UpdateSong(song *models.Song) error {
	return s.repo.UpdateSong(song)
}

func (s *SongService) DeleteSong(id uint) error {
	return s.repo.DeleteSong(id)
}
