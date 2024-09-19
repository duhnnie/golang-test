package usecase

import "music-rest-service/domain"

type IAlbumUseCase interface {
	GetAll() []*domain.Album
	GetById(id string) (*domain.Album, error)
	GetTracks(id string) []*domain.Track
	Create(title string, albumArtist string) (*domain.Album, error)
}
