package dto

// CreatePlaylistRequest is for creating a new playlist
type CreatePlaylistRequest struct {
	Name string `json:"name" binding:"required"`
}
