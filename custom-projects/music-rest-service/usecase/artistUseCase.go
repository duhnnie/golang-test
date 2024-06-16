package usecase

import (
	"music-rest-service/domain"
	"music-rest-service/repository"
)

type artistUseCase struct {
	repository      repository.IArtistRepository
	albumRepository repository.IAlbumRepository
	trackRepository repository.ITrackRepository
}

func NewArtistUseCase(repository repository.IArtistRepository, albumRepository repository.IAlbumRepository, trackRepository repository.ITrackRepository) *artistUseCase {
	return &artistUseCase{repository: repository, albumRepository: albumRepository, trackRepository: trackRepository}
}

func (auc *artistUseCase) GetAll() []*domain.Artist {
	return auc.repository.GetAll()
}

func (auc *artistUseCase) GetById(id string) (*domain.Artist, error) {
	return auc.repository.GetById(id)
}

func (auc *artistUseCase) GetAlbums(id string) ([]*domain.Album, error) {
	albums, err := auc.albumRepository.GetAlbumsByArtistId(id)

	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (auc *artistUseCase) GetTracks(id string) []*domain.Track {
	return auc.trackRepository.GetByArtistId(id)
}

func (auc *artistUseCase) Create(name string) (*domain.Artist, error) {
	return auc.repository.Create(name)
}
