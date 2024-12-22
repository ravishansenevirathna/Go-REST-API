package models

type LikedSong struct {
	ID     int `json:"id" gorm:"primaryKey"`
	UserID int `json:"user_id"`
	SongID int `json:"song_id"`
}
