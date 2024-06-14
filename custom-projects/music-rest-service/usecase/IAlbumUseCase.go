package usecase

import "music-rest-service/domain"

type IAlbumUseCase interface {
	getAll() []*domain.Album
	getById(id string) (*domain.Album, error)
	create(album *domain.Album) error
}
