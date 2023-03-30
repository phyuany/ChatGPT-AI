package middleware

import (
	"aiapp.pro/chat/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(ctx, fmt.Sprint(err), nil)
			}
		}()

		ctx.Next()
	}
}
