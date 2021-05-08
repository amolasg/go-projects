package controller

import (
	"encoding/json"
	"net/http"

	"github.com/amolasg/go-projects/clean-arch-example/entity"
	"github.com/amolasg/go-projects/clean-arch-example/errors"
	"github.com/amolasg/go-projects/clean-arch-example/service"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

func NewPostController(ser service.PostService) PostController {
	postService = ser
	return &controller{}
}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliction/json")

	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the post"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)

}

func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliction/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error getting posts"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := postService.Create(post)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "err savong the post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
