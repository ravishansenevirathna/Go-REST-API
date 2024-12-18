package service

import (
	"errors"
	"rest-api/models"
	"rest-api/repository"
)

type PlaylistService struct {
	repo     *repository.PlaylistRepository
	songRepo *repository.SongRepository
}

func NewPlaylistService(playlistRepo *repository.PlaylistRepository, songRepo *repository.SongRepository) *PlaylistService {
	return &PlaylistService{
		repo:     playlistRepo,
		songRepo: songRepo,
	}
}

func (s *PlaylistService) GetAllPlaylists() ([]models.Playlist, error) {
	return s.repo.GetAll()
}

func (s *PlaylistService) CreatePlaylist(playlist models.Playlist) (*models.Playlist, error) {
	if err := s.repo.Create(&playlist); err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (s *PlaylistService) AddSongToPlaylist(playlistID int, songID int) error {
	// Fetch the song by ID
	song, err := s.songRepo.GetByID(songID)
	if err != nil {
		return errors.New("song not found")
	}

	// Add the song to the playlist
	if err := s.repo.AddSongToPlaylist(playlistID, *song); err != nil {
		return err
	}
	return nil
}

func (s *PlaylistService) GetPlaylistByID(playlistID int) (*models.Playlist, error) {
	return s.repo.GetByID(playlistID)
}
