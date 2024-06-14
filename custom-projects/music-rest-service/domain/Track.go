package domain

type Track struct {
	title  string
	artist *Artist
	length uint16
}

func NewTrack(title string, artist *Artist, length uint16) *Track {
	return &Track{title: title, artist: artist, length: length}
}

func (t *Track) Title() string {
	return t.title
}
