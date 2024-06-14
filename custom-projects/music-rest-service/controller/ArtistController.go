package controller

import (
	"github.com/gin-gonic/gin"
	"music-rest-service/domain"
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
	ctx.IndentedJSON(http.StatusOK, []*domain.Artist{})
}

func (ac *ArtistController) GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, []*domain.Artist{})
}

func (ac *ArtistController) GetTracks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, []*domain.Artist{})
}
