package domain

type Album struct {
	title       string
	albumArtist *Artist
	tracks      []*Track
}

func NewAlbum(title string, albumArtist *Artist, tracks []*Track) *Album {
	return &Album{title, albumArtist, tracks}
}

func (a *Album) Title() string {
	return a.title
}

func (a *Album) AddTrack(track *Track) {
	a.tracks = append(a.tracks, track)
}

func (a *Album) Track(index int) Track {
	return *a.tracks[index]
}
