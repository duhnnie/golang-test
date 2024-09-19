package dto

type AlbumCreateDTO struct {
	Title string `json:"title" binding:"required"`
	AlbumArtist string `json:"albumArtist" binding:"required"`
}
