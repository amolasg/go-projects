package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amolasg/go-projects/go-fullstack/controllers"
	"github.com/amolasg/go-projects/go-fullstack/driver"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {

	db := driver.OpenConnection()

	controller := controllers.Controller{}

	router := mux.NewRouter()
	// api endpoints
	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET", "OPTIONS")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.DeleteBook(db)).Methods("DELETE")

	fmt.Println("Server listning on port 8080")
	// handle cors
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
