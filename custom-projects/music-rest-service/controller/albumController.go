package controller

import (
	"music-rest-service/api/response"
	"music-rest-service/dto"
	"music-rest-service/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type albumController struct {
	albumUseCase usecase.IAlbumUseCase
}

func NewAlbumController(albumUseCase usecase.IAlbumUseCase) *albumController {
	return &albumController{albumUseCase: albumUseCase}
}

func (ac *albumController) GetAll(ctx *gin.Context) {
	albums := ac.albumUseCase.GetAll()
	ctx.JSON(http.StatusOK, albums)
}

func (ac *albumController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	album, err := ac.albumUseCase.GetById(id)

	if err != nil {
		// TODO: how to know if the error is due a not found album?
		ctx.JSON(http.StatusNotFound, response.NewErrorResponseFrom(err))
	} else {
		ctx.JSON(http.StatusOK, album)
	}
}

func (ac *albumController) GetTracks(ctx *gin.Context) {
	id := ctx.Param("id")
	tracks := ac.albumUseCase.GetTracks(id)
	ctx.JSON(http.StatusOK, tracks)
}

func (ac *albumController) Create(ctx *gin.Context) {
	var albumCreateDTO dto.AlbumCreateDTO

	if err := ctx.ShouldBindJSON(&albumCreateDTO); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponseFrom(err))
		return
	}

	album, err := ac.albumUseCase.Create(albumCreateDTO.Title, albumCreateDTO.AlbumArtist)

	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponseFrom(err))
		return
	}

	ctx.IndentedJSON(http.StatusCreated, album)
}
