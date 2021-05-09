package main

import (
	"database/sql"
	"encoding/json"
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

func init() {
	gotenv.Load()
}

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("APP_DB_USERNAME"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

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
	db := OpenConnection()

	var book Book

	books = []Book{}

	rows, err := db.Query("select * from books")
	if err != nil {
		log.Fatal("error in getting records", err)
	}

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			log.Fatal("error in scan", err)
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
	defer rows.Close()
	defer db.Close()
}

func getBook(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var book Book

	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		log.Fatal("error in scan", err)
	}
	json.NewEncoder(w).Encode(book)
	defer db.Close()
}
func addBook(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var book Book

	var bookID int

	json.NewDecoder(r.Body).Decode(&book)

	err := db.QueryRow("insert into books(title,author,year)values($1,$2,$3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&bookID)

	if err != nil {
		log.Fatal("error in insert record", err)
	}
	json.NewEncoder(w).Encode(bookID)
	defer db.Close()
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	var book Book

	json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec("update books set title=$1,author=$2,year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)

	if err != nil {
		log.Fatal("error in updating", err)
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		log.Fatal("error in rows upadted", err)
	}
	json.NewEncoder(w).Encode(rowsUpdated)
	defer db.Close()
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	db := OpenConnection()

	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id=$1", params["id"])
	if err != nil {
		log.Fatal("error in delete record")
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		log.Fatal("error in RowsAffected")
	}
	json.NewEncoder(w).Encode(rowsDeleted)
	defer db.Close()
}
