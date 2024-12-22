package dto

type LikeSongRequestDto struct {
	UserID int `json:"userID" binding:"required"`
	SongID int `json:"songID" binding:"required"`
}
