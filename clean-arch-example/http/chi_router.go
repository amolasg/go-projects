package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	ChiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	ChiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	ChiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi HTTP Server running on 8080")
	http.ListenAndServe(port, ChiDispatcher)

}
