package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/liyq96/codereview-ai/internal/service"
	response "github.com/liyq96/codereview-ai/pkg/utils/resposne"
)

type CodeReviewHandler struct {
	codeReviewService *service.CodeReviewService
}

func NewCodeReviewHandler(codeReviewService *service.CodeReviewService) *CodeReviewHandler {
	return &CodeReviewHandler{codeReviewService: codeReviewService}
}

func (h *CodeReviewHandler) SubmitCodeReview(c *gin.Context) {
	// TODO code review logic
	resp := "代码审查完成"
	response.Success(resp)
}
