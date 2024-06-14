package repository

import (
	"errors"
	"music-rest-service/domain"
)

type ArtistRepository struct{}

func (ar *ArtistRepository) exists(artist domain.Artist) bool {
	for _, savedArtist := range DataArtist {
		if artist.Id() == savedArtist.Id() {
			return true
		}
	}

	return false
}

func (ar *ArtistRepository) GetAll() []*domain.Artist {
	return DataArtist
}

func (ar *ArtistRepository) GetById(id string) (*domain.Artist, error) {
	for _, artist := range DataArtist {
		if id == artist.Id() {
			return artist, nil
		}
	}

	return nil, errors.New("artist not found")
}

func (ar *ArtistRepository) Create(artist *domain.Artist) error {
	foundArtist, _ := ar.GetById(artist.Id())

	if foundArtist != nil {
		return errors.New("artist already exists")
	}

	DataArtist = append(DataArtist, artist)
	return nil
}
