package repository

import "music-rest-service/domain"

type IArtistRepository interface {
	GetAll() []*domain.Artist
	GetById(id string) (*domain.Artist, error)
	Create(artist *domain.Artist) error
}
