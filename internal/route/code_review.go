package route

import (
	"github.com/gin-gonic/gin"
	"github.com/liyq96/codereview-ai/internal/handler"
)

func RegisterCodeReviewRoutes(r *gin.RouterGroup, codeReviewHandler *handler.CodeReviewHandler) {
	codeReviewGroup := r.Group("/code_review")
	{
		codeReviewGroup.POST("/review_ai", codeReviewHandler.SubmitCodeReview)
	}
}
