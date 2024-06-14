package main

import (
	"fmt"
	"music-rest-service/api"
	"music-rest-service/controller"
	"music-rest-service/domain"
	"music-rest-service/repository"
	"music-rest-service/usecase"
)

func main() {
	artist := domain.NewArtist("asdfasdf", "Foo Fighters")
	track1 := domain.NewTrack("Doll", artist, 345)
	track2 := domain.NewTrack("Monkey Wrench", artist, 345)
	track3 := domain.NewTrack("Hey, Jonny Park!", artist, 345)
	track4 := domain.NewTrack("My Poor Brain", artist, 345)
	track5 := domain.NewTrack("Wind Up", artist, 345)
	track6 := domain.NewTrack("Up In Arms", artist, 345)
	track7 := domain.NewTrack("My Hero", artist, 345)
	track8 := domain.NewTrack("See You", artist, 345)
	track9 := domain.NewTrack("Enough Space", artist, 345)
	track10 := domain.NewTrack("February Stars", artist, 345)
	track11 := domain.NewTrack("Everlong", artist, 345)
	//track12 := domain.NewTrack("Walking After You", artist, 345)
	//track13 := domain.NewTrack("New Way Home", artist, 345)

	album := domain.NewAlbum("The Colour and the Shape", artist, []*domain.Track{track1, track2, track3})
	album.AddTrack(track4)
	album.AddTrack(track5)
	album.AddTrack(track6)
	album.AddTrack(track7)
	album.AddTrack(track8)
	album.AddTrack(track9)
	album.AddTrack(track10)
	album.AddTrack(track11)

	//fmt.Printf("%v - \"%v\"", artist.Name(), track.Title())
	fmt.Println(album)

	secondTrack := album.Track(1)
	fmt.Println(secondTrack)

	secondTrack = album.Track(1)
	fmt.Println(secondTrack)

	artistRepository := repository.ArtistRepository{}
	artistUseCase := usecase.NewArtistUseCase(&artistRepository)
	artistController := controller.NewArtistController(artistUseCase)
	apiServer := api.New(artistController)
	apiServer.Run("8080")
}
