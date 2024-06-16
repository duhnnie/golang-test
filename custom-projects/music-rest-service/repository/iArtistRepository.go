package repository

import "music-rest-service/domain"

type IArtistRepository interface {
	GetAll() []*domain.Artist
	GetById(id string) (*domain.Artist, error)
	GetByName(name string) (*domain.Artist, error)
	Create(name string) (*domain.Artist, error)
}
