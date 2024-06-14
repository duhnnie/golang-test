package usecase

import "music-rest-service/domain"

type IArtistUseCase interface {
	GetAll() []*domain.Artist
	GetById(id string) (*domain.Artist, error)
	Create(artist *domain.Artist) (*domain.Artist, error)
}
