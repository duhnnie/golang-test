package controller

import "github.com/gin-gonic/gin"

type ITrackController interface {
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}
