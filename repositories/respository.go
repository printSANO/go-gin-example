package repositories

import "gorm.io/gorm"

type Repository struct {
	UserRepository UserRepository
	PostRepository PostRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		PostRepository: NewPostRepository(db),
	}
}
