package domain

import (
	"github.com/google/uuid"
	"strings"
)

type Artist struct {
	Id   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

func NewArtistWithId(id string, name string) *Artist {
	return &Artist{Id: strings.Trim(id, " "), Name: strings.Trim(name, " ")}
}

func NewArtist(name string) *Artist {
	return NewArtistWithId(uuid.NewString(), name)
}
