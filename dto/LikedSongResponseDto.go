package dto

type LikedSongResponseDto struct {
	SongID   int    `json:"song_id"`
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Duration string `json:"duration"`
}
