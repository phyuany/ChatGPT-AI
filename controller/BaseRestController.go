package controller

import "github.com/gin-gonic/gin"

type BaseRestController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Show(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
