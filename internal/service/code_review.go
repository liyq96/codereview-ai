package service

import (
	"context"
	"fmt"

	"github.com/liyq96/codereview-ai/internal/config"
	openai "github.com/sashabaranov/go-openai"
)

type CodeReviewService struct{}

func NewCodeReviewService() *CodeReviewService {
	return &CodeReviewService{}
}

func (s *CodeReviewService) SubmitCodeReview(code string, cfg *config.SystemConfig) {
	config := openai.DefaultConfig("")
	config.BaseURL = "https://ark.cn-beijing.volces.com/api/v3"
	fmt.Println("config.BaseURL", config.BaseURL)
	client := openai.NewClientWithConfig(config)

	fmt.Println("----- standard request -----")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "deepseek-r1-distill-qwen-32b-250120",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "你是人工智能助手",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "常见的十字花科植物有哪些？",
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Message.Content)

}
