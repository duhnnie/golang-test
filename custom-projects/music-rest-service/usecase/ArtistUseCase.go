package usecase

import (
	"music-rest-service/domain"
	"music-rest-service/repository"
)

type ArtistUseCase struct {
	repository repository.IArtistRepository
}

func NewArtistUseCase(repository repository.IArtistRepository) *ArtistUseCase {
	return &ArtistUseCase{repository: repository}
}

func (auc *ArtistUseCase) GetAll() []*domain.Artist {
	return auc.repository.GetAll()
}

func (auc *ArtistUseCase) GetById(id string) (*domain.Artist, error) {
	return auc.repository.GetById(id)
}

func (auc *ArtistUseCase) Create(artist *domain.Artist) (*domain.Artist, error) {
	err := auc.repository.Create(artist)

	if err != nil {
		return &(domain.Artist{}), err
	}

	return artist, nil
}
