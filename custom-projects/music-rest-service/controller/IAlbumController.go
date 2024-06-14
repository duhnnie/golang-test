package controller

import (
	"github.com/gin-gonic/gin"
)

type IAlbumController interface {
	GetAll(ctx *gin.Context)
	GetById(ctx *gin.Context)
	GetTracks(ctx *gin.Context)
}
