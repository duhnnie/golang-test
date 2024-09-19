package usecase

import (
	"errors"
	"music-rest-service/domain"
	"music-rest-service/repository"
)

type albumUseCase struct {
	albumRepository  repository.IAlbumRepository
	artistRepository repository.IArtistRepository
	trackRepository  repository.ITrackRepository
}

func NewAlbumUseCase(
	albumRepository repository.IAlbumRepository,
	artistRepository repository.IArtistRepository,
	trackRepository repository.ITrackRepository,
) *albumUseCase {
	return &albumUseCase{
		albumRepository:  albumRepository,
		artistRepository: artistRepository,
		trackRepository:  trackRepository,
	}
}

func (auc *albumUseCase) GetAll() []*domain.Album {
	return auc.albumRepository.GetAll()
}

func (auc *albumUseCase) GetById(id string) (*domain.Album, error) {
	return auc.albumRepository.GetAlbumById(id)
}

func (auc *albumUseCase) GetTracks(id string) []*domain.Track {
	return auc.trackRepository.GetByAlbumId(id)
}

func (auc *albumUseCase) Create(title string, artistId string) (*domain.Album, error) {
	artist, err := auc.artistRepository.GetById(artistId)

	if err != nil {
		return nil, errors.New("artist not found")
	}

	return auc.albumRepository.Create(title, artist)
}
