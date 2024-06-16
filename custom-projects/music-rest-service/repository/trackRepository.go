package repository

import (
	"errors"
	"music-rest-service/domain"
)

type trackRepository struct{}

func NewTrackRepository() *trackRepository {
	return &trackRepository{}
}

func (tr *trackRepository) GetById(id string) (*domain.Track, error) {
	for _, track := range DataTrack {
		if track.Id == id {
			return track, nil
		}
	}

	return nil, errors.New("Track not found.")
}

func (tr *trackRepository) GetByArtistId(artistId string) []*domain.Track {
	var tracks []*domain.Track

	for _, track := range DataTrack {
		if track.Artist.Id == artistId {
			tracks = append(tracks, track)
		}
	}

	return tracks
}

func (tr *trackRepository) GetByAlbumId(albumId string) []*domain.Track {
	for _, album := range DataAlbum {
		if album.Id == albumId {
			return album.Tracks
		}
	}

	return []*domain.Track{}
}
