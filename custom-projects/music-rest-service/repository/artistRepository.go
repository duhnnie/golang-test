package repository

import (
	"errors"
	"music-rest-service/domain"
)

type ArtistRepository struct{}

func (ar *ArtistRepository) exists(artist domain.Artist) bool {
	for _, savedArtist := range DataArtist {
		if artist.Id == savedArtist.Id {
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
		if id == artist.Id {
			return artist, nil
		}
	}

	return nil, errors.New("artist not found")
}

func (ar *ArtistRepository) GetByName(name string) (*domain.Artist, error) {
	for _, artist := range DataArtist {
		if name == artist.Name {
			return artist, nil
		}
	}

	return nil, errors.New("artist not found")
}

func (ar *ArtistRepository) Create(name string) (*domain.Artist, error) {
	newArtist := domain.NewArtist(name)
	foundArtist, _ := ar.GetByName(newArtist.Name)

	if foundArtist != nil {
		return nil, errors.New("artist already exists")
	}

	DataArtist = append(DataArtist, newArtist)
	return newArtist, nil
}
