package handler

import (
	"fmt"

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
	var data map[string]any
	err := c.ShouldBindJSON(&data)
	if err != nil {
		response.Error(400, "参数错误")
		return
	}
	code, ok := data["code_content"].(string)
	if !ok {
		response.Error(400, "参数错误")
	}
	fmt.Println("code_content:", code)
	go h.codeReviewService.SubmitCodeReview(code)
	response.Success("")
}
