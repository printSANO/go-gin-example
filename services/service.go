package services

import "github.com/printSANO/go-gin-example/repositories"

// Service groups all individual services.
type Service struct {
	UserService UserService
	PostService PostService
}

// NewService creates a new instance of Service with all required services.
func NewService(repo *repositories.Repository) *Service {
	return &Service{
		UserService: NewUserService(repo.UserRepository),
		PostService: NewPostService(repo.PostRepository),
	}
}
