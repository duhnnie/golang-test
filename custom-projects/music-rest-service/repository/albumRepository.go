package repository

import (
	"errors"
	"music-rest-service/domain"
)

type albumRepository struct{}

func NewAlbumRepository() *albumRepository {
	return &albumRepository{}
}

func (ar *albumRepository) GetAll() ([]*domain.Album) {
	return DataAlbum
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

func (ar *albumRepository) Create(title string, albumArtist *domain.Artist) (*domain.Album, error) {
	album := domain.NewAlbum(title, albumArtist, []*domain.Track{})
	DataAlbum = append(DataAlbum, album)
	return album, nil
}

// func (ar *albumRepository) GetTracks(albumId string) ([]*domain.Track, error) {
// 	album, err := ar.GetAlbumById(albumId)

// 	if err != nil {
// 			return nil, errors.New("album not found")
// 	}

// 	return album.Tracks, nil
// }
