package service

import (
	"rest-api/models"
	"rest-api/repository"
)

type LikedSongService struct {
	repo *repository.LikedSongRepository
}

func NewLikedSongService(repo *repository.LikedSongRepository) *LikedSongService {
	return &LikedSongService{repo: repo}
}

func (s *LikedSongService) LikeSong(userID int, songID int) error {
	return s.repo.LikeSong(userID, songID)
}

func (s *LikedSongService) GetLikedSongsByUser(userID int) ([]models.Song, error) {
	return s.repo.GetLikedSongsByUser(userID)
}
