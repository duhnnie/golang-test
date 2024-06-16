package domain

import "github.com/google/uuid"

type Track struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist *Artist `json:"artist"`
	Length uint16  `json:"length"`
}

func NewTrack(title string, artist *Artist, length uint16) *Track {
	return &Track{Id: uuid.NewString(), Title: title, Artist: artist, Length: length}
}
