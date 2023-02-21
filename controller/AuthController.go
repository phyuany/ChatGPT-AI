package controller

import (
	"github.com/gin-gonic/gin"
	"jkapp.net/ai/response"
)

func Info(ctx *gin.Context) {
	token, _ := ctx.Get("token")
	response.Success(ctx, gin.H{"token": token}, "success")
}

func Index(ctx *gin.Context) {
	domain := ctx.Request.Host

	response.Success(ctx,
		gin.H{
			"completions": domain + "/completions",
		},
		"success")
}
