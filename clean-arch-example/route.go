package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "text"}}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliction/json")

	result, err := json.Marshal(posts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error" : "Error marshalling the posts array"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "appliction/json")
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error" : "Error unmarshalling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	w.WriteHeader(http.StatusOK)
	res, err := json.Marshal(posts)
	w.Write(res)

}
