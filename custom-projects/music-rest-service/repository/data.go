package repository

import "music-rest-service/domain"

var artistFooFighters = domain.NewArtistWithId("123", "Foo Fighters")
var artistRadiohead = domain.NewArtistWithId("234", "Radiohead")
var artistNofx = domain.NewArtistWithId("345", "NOFX")
var artistLlegas = domain.NewArtistWithId("456", "Llegas")
var artistMammut = domain.NewArtistWithId("567", "Mammut")

var trackStacked = domain.NewTrack("Stacked Actors", artistFooFighters, 123)
var trackBreakout = domain.NewTrack("Breakout", artistFooFighters, 345)
var trackGenerator = domain.NewTrack("Generator", artistFooFighters, 45)
var trackAirbag = domain.NewTrack("Airbag", artistRadiohead, 34)
var trackParanoid = domain.NewTrack("Paranoid Android", artistRadiohead, 34)
var trackSubterranean = domain.NewTrack("Subterranean Homesick Alien", artistRadiohead, 34)
var trackLinoleum = domain.NewTrack("Linoleum", artistNofx, 34)
var trackLeave = domain.NewTrack("Leave it Alone", artistNofx, 34)
var trackDig = domain.NewTrack("Dig", artistNofx, 34)
var trackAntifaz = domain.NewTrack("Antifaz", artistLlegas, 34)
var trackRespiros = domain.NewTrack("Respiros", artistLlegas, 34)
var trackHuesos = domain.NewTrack("Huesos", artistLlegas, 34)
var trackLa = domain.NewTrack("La misma sed", artistMammut, 34)
var trackParamos = domain.NewTrack("Páramos", artistMammut, 34)
var trackLento = domain.NewTrack("Lento", artistMammut, 34)

var DataArtist = []*domain.Artist{
	artistFooFighters,
	artistRadiohead,
	artistNofx,
	artistLlegas,
	artistMammut,
}

var DataTrack = []*domain.Track{
	trackStacked,
	trackBreakout,
	trackGenerator,
	trackAirbag,
	trackParanoid,
	trackSubterranean,
	trackLinoleum,
	trackLeave,
	trackDig,
	trackAntifaz,
	trackRespiros,
	trackHuesos,
	trackLa,
	trackParamos,
	trackLento,
}

var albumNothingToLose = domain.NewAlbum(
	"There is Nothing Left To Lose",
	artistFooFighters,
	[]*domain.Track{trackStacked, trackBreakout, trackGenerator},
)

var albumOk = domain.NewAlbum(
	"OK Computer",
	artistRadiohead,
	[]*domain.Track{trackAirbag, trackParanoid, trackSubterranean},
)

var albumPunk = domain.NewAlbum(
	"Punk in Drublic",
	artistNofx,
	[]*domain.Track{trackLinoleum, trackLeave, trackDig},
)

var albumPesanervios = domain.NewAlbum(
	"El Pesanervios",
	artistLlegas,
	[]*domain.Track{trackAntifaz, trackHuesos, trackRespiros},
)

var albumLodo = domain.NewAlbum(
	"Lodo",
	artistMammut,
	[]*domain.Track{trackLa, trackParamos, trackLento})

var DataAlbum = []*domain.Album{
	albumNothingToLose,
	albumLodo,
	albumOk,
	albumPesanervios,
	albumPunk,
}
