package controller

import (
	"github.com/gin-gonic/gin"
)

type IArtistController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetAlbums(ctx *gin.Context)
	GetTracks(ctx *gin.Context)
	Post(ctx *gin.Context)
}
