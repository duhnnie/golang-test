package api

import (
	"music-rest-service/controller"

	"github.com/gin-gonic/gin"
)

type API struct {
	router *gin.Engine
}

// func New(artistController controller.IArtistController, albumController controller.IAlbumController, trackController controller.ITrackController) *API {
func New(artistController controller.IArtistController, albumController controller.IAlbumController) *API {
	router := gin.Default()

	// artists
	router.GET("/artists", artistController.GetAll)
	router.GET("/artists/:id", artistController.GetById)
	router.GET("/artists/:id/albums", artistController.GetAlbums)
	router.GET("/artists/:id/tracks", artistController.GetTracks)
	router.POST("/artists", artistController.Post)

	// albums
	router.GET("/albums", albumController.GetAll)
	router.POST("/albums", albumController.Create)
	router.GET("/albums/:id", albumController.GetById)
	router.GET("/albums/:id/tracks", albumController.GetTracks)
	//
	//// tracks
	//router.GET("/tracks", trackController.GetAll)
	//router.GET("/tracks/:id", trackController.GetByID)

	return &API{router: router}
}

func (api *API) Run(port string) error {
	return api.router.Run(":" + port)
}
