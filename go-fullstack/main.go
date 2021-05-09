package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

// books empty slice will holding book records
var books []Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {

	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("APP_DB_USERNAME"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_NAME"))

	var err error
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	error := db.Ping()
	if err != nil {
		log.Fatal("error in connection", error)
	}

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
