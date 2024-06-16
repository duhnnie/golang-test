package repository

import "music-rest-service/domain"

type ITrackRepository interface {
	GetById(id string) (*domain.Track, error)
	GetByArtistId(artistId string) []*domain.Track
	GetByAlbumId(albumId string) []*domain.Track
}
