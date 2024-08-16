package services

import (
	"github.com/printSANO/go-gin-example/models"
	"github.com/printSANO/go-gin-example/repositories"
)

type PostService interface {
	GetPostByID(id uint) (*models.Post, error)
	CreatePost(post *models.Post) error
}

type postService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	return s.repo.FindByID(id)
}

func (s *postService) CreatePost(post *models.Post) error {
	return s.repo.Create(post)
}
