package service

import "github.com/liyq96/codereview-ai/internal/config"

type Services struct {
	CodeReviewService *CodeReviewService
}

func InitServices(cfg *config.SystemConfig) *Services {
	return &Services{
		CodeReviewService: NewCodeReviewService(cfg),
	}
}
