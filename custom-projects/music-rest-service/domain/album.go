package domain

import "github.com/google/uuid"

type Album struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	AlbumArtist *Artist  `json:"artist"`
	Tracks      []*Track `json:"tracks"`
}

func NewAlbum(title string, albumArtist *Artist, tracks []*Track) *Album {
	return &Album{uuid.NewString(), title, albumArtist, tracks}
}

func (a *Album) AddTrack(track *Track) {
	a.Tracks = append(a.Tracks, track)
}

func (a *Album) Track(index int) Track {
	return *a.Tracks[index]
}
