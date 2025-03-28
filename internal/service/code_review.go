package service

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/liyq96/codereview-ai/internal/config"
	openai "github.com/sashabaranov/go-openai"
)

type CodeReviewService struct {
	cfg *config.SystemConfig
}

func NewCodeReviewService(cfg *config.SystemConfig) *CodeReviewService {
	return &CodeReviewService{cfg: cfg}
}

func (s *CodeReviewService) SubmitCodeReview(code string) {
	config := openai.DefaultConfig(s.cfg.Ai.ApiKey)
	config.BaseURL = s.cfg.Ai.BaseUrl
	client := openai.NewClientWithConfig(config)

	// 解析提示词文档
	prompt, err := ReadMarkdownFile(s.cfg.Ai.PromptFilePath)
	if err != nil {
		fmt.Printf("ReadMarkdownFile error: %v\n", err)
		prompt = "请帮我审查一下这段代码，依照markdown的格式回复建议，按照下面的分类分别给出改进意见及原因，\n\n # 必需要改进内容 ## 原因 \n\n 建议优化的内容 ## 原因 \n\n 规范建议"
	}
	fmt.Println("----- standard request -----")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: s.cfg.Ai.ModelId,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: code,
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

func ReadMarkdownFile(filePath string) (string, error) {
	// 打开指定路径的 Markdown 文件
	file, err := os.Open(filePath)
	if err != nil {
		// 如果打开文件时出现错误，返回错误信息
		return "", err
	}
	// 确保在函数返回时关闭文件，避免资源泄漏
	defer file.Close()

	// 创建一个字符串构建器，用于高效地拼接字符串
	var contentBuilder string
	// 创建一个扫描器，用于逐行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// 将每一行内容添加到字符串构建器中，并添加换行符
		contentBuilder += scanner.Text() + "\n"
	}

	// 检查扫描过程中是否出现错误
	if err := scanner.Err(); err != nil {
		// 如果出现错误，返回错误信息
		return "", err
	}

	// 返回读取到的文件内容字符串
	return contentBuilder, nil
}
