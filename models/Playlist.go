package models

// Playlist represents a music playlist
type Playlist struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Songs []Song `gorm:"many2many:playlist_songs;"`
}
