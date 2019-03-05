package services

import "michiganhackers/class-review-backend/repositories"

type Services struct {
	ReviewService IReviewService
}

func DefaultServices(repos *repositories.Repositories) *Services {
	return &Services{
		ReviewService: DefaultReviewServices(repos),
	}

}
