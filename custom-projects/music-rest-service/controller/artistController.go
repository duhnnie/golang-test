package controller

import (
	"github.com/gin-gonic/gin"
	"music-rest-service/api/response"
	"music-rest-service/dto"
	"music-rest-service/usecase"
	"net/http"
)

type ArtistController struct {
	useCase usecase.IArtistUseCase
}

func NewArtistController(useCase usecase.IArtistUseCase) *ArtistController {
	return &ArtistController{useCase: useCase}
}

func (ac *ArtistController) GetAll(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, ac.useCase.GetAll())
}

func (ac *ArtistController) GetById(ctx *gin.Context) {
	artistId := ctx.Param("id")
	artist, err := ac.useCase.GetById(artistId)

	if err != nil {
		errResponse := response.NewErrorResponseFrom(err)
		ctx.IndentedJSON(404, errResponse)
	} else {
		ctx.IndentedJSON(http.StatusOK, artist)
	}
}

func (ac *ArtistController) GetAlbums(ctx *gin.Context) {
	artistId := ctx.Param("id")
	_, err := ac.useCase.GetById(artistId)

	if err != nil {
		errResponse := response.NewErrorResponseFrom(err)
		ctx.IndentedJSON(http.StatusNotFound, errResponse)
		return
	}

	albums, err := ac.useCase.GetAlbums(artistId)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, response.NewErrorResponseFrom(err))
	} else {
		ctx.IndentedJSON(http.StatusOK, albums)
	}
}

func (ac *ArtistController) GetTracks(ctx *gin.Context) {
	artistId := ctx.Param("id")
	_, err := ac.useCase.GetById(artistId)

	if err != nil {
		ctx.IndentedJSON(404, response.NewErrorResponseFrom(err))
	} else {
		ctx.IndentedJSON(http.StatusOK, ac.useCase.GetTracks(artistId))
	}
}

func (ac *ArtistController) Post(ctx *gin.Context) {
	var artistDto dto.NewArtistDTO

	if err := ctx.ShouldBindJSON(&artistDto); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	artist, err := ac.useCase.Create(artistDto.Name)

	if err != nil {
		ctx.Error(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response.NewErrorResponseFrom(err))
		return
	}

	ctx.IndentedJSON(http.StatusOK, artist)
}
