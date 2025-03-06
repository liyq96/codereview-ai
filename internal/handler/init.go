package handler

import "github.com/liyq96/codereview-ai/internal/service"

type Handlers struct {
	CodeReviewHandler *CodeReviewHandler
}

func InitHandlers(services *service.Services) *Handlers {
	return &Handlers{
		CodeReviewHandler: NewCodeReviewHandler(services.CodeReviewService),
	}
}
