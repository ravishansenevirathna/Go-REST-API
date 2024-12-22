package repository

import (
	"gorm.io/gorm"
	"rest-api/models"
)

type LikedSongRepository struct {
	DB *gorm.DB
}

func NewLikedSongRepository(db *gorm.DB) *LikedSongRepository {
	return &LikedSongRepository{DB: db}
}

func (r *LikedSongRepository) LikeSong(userID int, songID int) error {
	likedSong := models.LikedSong{UserID: userID, SongID: songID}
	return r.DB.Create(&likedSong).Error
}

func (r *LikedSongRepository) GetLikedSongsByUser(userID int) ([]models.Song, error) {
	var songs []models.Song
	err := r.DB.Table("songs").Select("songs.*").
		Joins("INNER JOIN liked_songs ON liked_songs.song_id = songs.id").
		Where("liked_songs.user_id = ?", userID).Find(&songs).Error
	return songs, err
}
