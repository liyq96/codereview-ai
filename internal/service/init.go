package service

type Services struct {
	CodeReviewService *CodeReviewService
}

func InitServices() *Services {
	return &Services{
		CodeReviewService: NewCodeReviewService(),
	}
}
