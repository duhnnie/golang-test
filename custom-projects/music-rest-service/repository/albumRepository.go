package repository

import (
	"errors"
	"music-rest-service/domain"
)

type albumRepository struct{}

func NewAlbumRepository() *albumRepository {
	return &albumRepository{}
}

func (ar *albumRepository) GetAlbumById(id string) (*domain.Album, error) {
	for _, album := range DataAlbum {
		if album.Id == id {
			return album, nil
		}
	}

	return nil, errors.New("album not found")
}

func (ar *albumRepository) GetAlbumsByArtistId(artistId string) ([]*domain.Album, error) {
	albums := []*domain.Album{}

	for _, album := range DataAlbum {
		if album.AlbumArtist.Id == artistId {
			albums = append(albums, album)
		}
	}

	return albums, nil
}
