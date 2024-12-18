package dto

type CreateSongRequest struct {
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Album  string `json:"album"`
}
