package usecase

import "music-rest-service/domain"

type ITrackUseCase interface {
	getAll() []*domain.Track
	getById(id string) (*domain.Track, error)
	create(track *domain.Track) (*domain.Track, error)
}
