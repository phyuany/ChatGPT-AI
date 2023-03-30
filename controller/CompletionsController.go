package controller

import (
	"aiapp.pro/chat/response"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
	"time"
)

type ICompletionsController interface {
	BaseRestController
}

type CompletionsController struct {
}

func NewCompletionsController() CompletionsController {
	return CompletionsController{}
}

func (com CompletionsController) Show(ginCtx *gin.Context) {
	token := ginCtx.Query("token")
	stringLen := len([]rune(token))

	if stringLen < 2 || stringLen > 100 {
		response.Fail(ginCtx, "error", gin.H{"error_detail": "token length must be greater than 2 and less than 100"})
		return
	}

	// 获取环境变量
	gpt3Key := os.Getenv("GPT3_API_KEY")
	c := openai.NewClient(gpt3Key)
	ctx := context.Background()

	req := openai.CompletionRequest{
		Model:     openai.GPT3TextDavinci003,
		MaxTokens: 512,
		Prompt:    token,
		Stream:    true,
	}
	stream, err := c.CreateCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("CompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	ginCtx.Header("Cache-Control", "no-cache")
	ginCtx.Header("Connection", "keep-alive")
	ginCtx.Header("Access-Control-Allow-Origin", "*")
	ginCtx.Header("Content-Type", "text/event-stream; charset=utf-8")

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			break
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			break
		}

		fmt.Printf("Response: %+v\n", response)
		data := "data: " + response.Choices[0].Text + "\n\n"
		ginCtx.Writer.WriteString(data)

		// 将缓冲区数据刷到客户端
		ginCtx.Writer.Flush()
	}
	ginCtx.Writer.WriteString("event: close\ndata: close\ndata: \n\n")
}

func (com CompletionsController) SSE(ginCtx *gin.Context) {
	ginCtx.Header("Content-Type", "text/event-stream")
	ginCtx.Header("Cache-Control", "no-cache")
	ginCtx.Header("Connection", "keep-alive")

	message := "hello"

	// 设置计数器
	counter := 0
	for {
		// 获取当前时间
		now := time.Now().Format(time.RFC3339)
		// 时间格式化
		now = now[11:19]
		// 设置 SSE 格式消息
		data := "data: " + message + "time: " + now + ", counter:" + fmt.Sprintf("%d", counter) + "\n\n"
		ginCtx.Writer.WriteString(data)

		// 将缓冲区数据刷到客户端
		ginCtx.Writer.Flush()

		// 模拟延迟
		time.Sleep(1 * time.Second)

		// 计数器自增
		counter++

		// 计数器大于 10 时，退出循环
		if counter > 10 {
			break
		}
	}
	// 发送关闭连接的消息
	ginCtx.Writer.WriteString("event: close\ndata: close\ndata: \n\n")
}
