package service

import (
	"rest-api/models"
	"rest-api/repository"
)

type RecommendationService struct {
	UserRepo *repository.UserRepository
	SongRepo *repository.SongRepository
}

func NewRecommendationService(userRepo *repository.UserRepository, songRepo *repository.SongRepository) *RecommendationService {
	return &RecommendationService{
		UserRepo: userRepo,
		SongRepo: songRepo,
	}
}

func (s *RecommendationService) RecommendSongs(userID uint, limit int) ([]models.Song, error) {

	return nil, nil
}
