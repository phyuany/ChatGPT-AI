package controller

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"jkapp.net/ai/response"
)

type ICompletionsController interface {
	BaseRestController
}

type CompletionsController struct {
}

func NewCompletionsController() CompletionsController {
	return CompletionsController{}
}

func (c CompletionsController) Show(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	stringLen := len([]rune(token))

	if stringLen < 2 || stringLen > 100 {
		response.Fail(ginCtx, "error", gin.H{"error_detail": "token length must be greater than 2 and less than 100"})
		return
	}

	// 获取环境变量
	gpt3Key := os.Getenv("GPT3_API_KEY")

	client := gogpt.NewClient(gpt3Key)
	backCtx := context.Background()

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   1024,
		Prompt:      token,
		Temperature: 1,
		TopP:        1,
	}
	resp, err := client.CreateCompletion(backCtx, req)
	if err != nil {
		response.Fail(ginCtx, "error", gin.H{"error_detail": err.Error()})
		return
	}

	response.Success(ginCtx, gin.H{"result": resp.Choices[0].Text}, "success")
}
