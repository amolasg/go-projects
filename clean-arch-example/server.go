package main

import (
	"fmt"
	"net/http"

	"github.com/amolasg/go-projects/clean-arch-example/controller"
	router "github.com/amolasg/go-projects/clean-arch-example/http"
	"github.com/amolasg/go-projects/clean-arch-example/repository"
	"github.com/amolasg/go-projects/clean-arch-example/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
	//httpRouter     router.Router             = router.NewChiRouter()
)

func main() {

	httpRouter.GET("/", func(resp http.ResponseWriter, w *http.Request) {
		fmt.Fprintln(resp, " UP and Running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(":8080")

}
