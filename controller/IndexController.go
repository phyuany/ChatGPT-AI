package controller

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
	"jkapp.net/ai/response"
)

func Info(ctx *gin.Context) {
	token, _ := ctx.Get("token")
	response.Success(ctx, gin.H{"token": token}, "success")
}

func Index(ctx *gin.Context) {

	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatalln(err)
	}

	// 渲染到浏览器
	tmpl.Execute(ctx.Writer, nil)
}
