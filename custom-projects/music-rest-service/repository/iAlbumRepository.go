package repository

import "music-rest-service/domain"

type IAlbumRepository interface {
	GetAlbumById(id string) (*domain.Album, error)
	GetAlbumsByArtistId(artistId string) ([]*domain.Album, error)
}
