package usecase

import "music-rest-service/domain"

type IArtistUseCase interface {
	GetAll() []*domain.Artist
	GetById(id string) (*domain.Artist, error)
	GetAlbums(id string) ([]*domain.Album, error)
	GetTracks(id string) []*domain.Track
	Create(name string) (*domain.Artist, error)
}
