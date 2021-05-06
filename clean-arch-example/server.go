package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func(resp http.ResponseWriter, w *http.Request) {
		fmt.Fprintln(resp, "Up...")
	})

	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("server listining on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
