package repository

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/amolasg/go-projects/clean-arch-example/entity"
)

// need this struct to implement interface
type repo struct{}

// constructor to initilize new instance
// NewRepository

func NewFirestoreRepository() PostRepository {
	return &repo{} // refernce to repo struct
}

const (
	projectId      string = "rest-example-95b90"
	collectionName string = "posts"
)

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed adding a new Post: %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post

	iterator := client.Collection(collectionName).Documents(ctx)

	doc, err := iterator.GetAll()
	if err != nil {
		log.Fatalf("Failed to iterate a list of posts %v", err)
		return nil, err
	}

	for _, v := range doc {
		m := v.Data()
		posts = append(posts, entity.Post{
			ID:    m["ID"].(int64),
			Title: m["Title"].(string),
			Text:  m["Text"].(string),
		})
	}

	return posts, nil
}
