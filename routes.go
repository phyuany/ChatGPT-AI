package main

import (
	"github.com/gin-gonic/gin"
	"jkapp.net/ai/controller"
	"jkapp.net/ai/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.GET("", controller.Index)
	// 预留一个需鉴权的接口
	r.GET("/auth/info", middleware.AuthMiddleware(), controller.Info)

	completionsRoutes := r.Group("/completions")
	completionsController := controller.NewCompletionsController()
	completionsRoutes.GET("", completionsController.Show)

	return r
}
