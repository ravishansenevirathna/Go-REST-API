package dto

// AddSongRequest is for adding a song to a playlist
type AddSongRequest struct {
	SongID int `json:"song_id" binding:"required"`
}
