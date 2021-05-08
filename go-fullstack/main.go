package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

// books empty slice will holding book records
var books []Book

func main() {

	router := mux.NewRouter()
	// api endpoints
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	// server listning on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

// implementing handler functions

func getBooks(w http.ResponseWriter, r *http.Request) {

}
func getBook(w http.ResponseWriter, r *http.Request) {

}
func addBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}
