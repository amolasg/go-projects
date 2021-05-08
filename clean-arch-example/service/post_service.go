package service

import (
	"errors"
	"math/rand"

	"github.com/amolasg/go-projects/clean-arch-example/entity"
	"github.com/amolasg/go-projects/clean-arch-example/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct {
	repo repository.PostRepository
}

// constructor func
func NewPostService(repo repository.PostRepository) PostService {
	return &service{repo: repo}
}

func (s *service) Validate(post *entity.Post) error {

	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}

	return nil
}

func (s *service) Create(post entity.Post) (*entity.Post, error) {
	post.ID = int64(rand.Int())
	entity, err := s.repo.Save(&post)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (s *service) FindAll() ([]entity.Post, error) {
	return s.repo.FindAll()
}
