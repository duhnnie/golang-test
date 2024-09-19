package repository

import "music-rest-service/domain"

type IAlbumRepository interface {
	GetAll() []*domain.Album
	GetAlbumById(id string) (*domain.Album, error)
	GetAlbumsByArtistId(artistId string) ([]*domain.Album, error)
	Create(title string, albumArtist *domain.Artist) (*domain.Album, error)
	// GetTracks(albumId string) ([]*domain.Track, error)
}
