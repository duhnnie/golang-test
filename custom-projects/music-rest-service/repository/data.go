package repository

import "music-rest-service/domain"

var DataArtist = []*domain.Artist{
	domain.NewArtist("123", "Foo Fighters"),
	domain.NewArtist("234", "Radiohead"),
	domain.NewArtist("345", "NOFX"),
	domain.NewArtist("456", "Llegas"),
	domain.NewArtist("567", "Mammut"),
}
