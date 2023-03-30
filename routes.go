package main

import (
	"aiapp.pro/chat/controller"
	"aiapp.pro/chat/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.GET("", controller.Index)
	// 预留一个需鉴权的接口
	r.GET("/auth/info", middleware.AuthMiddleware(), controller.Info)

	completionsRoutes := r.Group("/completions")
	completionsController := controller.NewCompletionsController()
	completionsRoutes.GET("", completionsController.Show)
	completionsRoutes.GET("/sse", completionsController.SSE)

	return r
}
