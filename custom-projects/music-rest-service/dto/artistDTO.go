package dto

type NewArtistDTO struct {
	Name string `json:"name" binding:"required"`
}
