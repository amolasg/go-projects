package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
}

var (
	MuxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	MuxDispatcher.HandleFunc(uri, f).Methods("GET")
}
func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	MuxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Println("Mux HTTP Server running on 8080")

	http.ListenAndServe(port, MuxDispatcher)

}
