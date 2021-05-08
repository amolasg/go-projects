package repository

import "github.com/amolasg/go-projects/clean-arch-example/entity"

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
