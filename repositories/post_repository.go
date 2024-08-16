package repositories

import (
	"github.com/printSANO/go-gin-example/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	FindByID(id uint) (*models.Post, error)
	Create(post *models.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, id).Error
	return &post, err
}

func (r *postRepository) Create(post *models.Post) error {
	return r.db.Create(post).Error
}
