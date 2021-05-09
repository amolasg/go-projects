package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/amolasg/go-projects/go-fullstack/models"
	bookRepository "github.com/amolasg/go-projects/go-fullstack/repository/book"
	"github.com/amolasg/go-projects/go-fullstack/utils"
	"github.com/gorilla/mux"
)

type Controller struct{}

var books []models.Book

// implementing handler functions

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var book models.Book
		var error models.Error
		books = []models.Book{}
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		params := mux.Vars(r)

		books = []models.Book{}

		bookRepo := bookRepository.BookRepository{}

		id, _ := strconv.Atoi(params["id"])

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			if err == sql.ErrNoRows {
				error.Message = "Not Found"
				utils.SendError(w, http.StatusNotFound, error)
				return
			} else {
				error.Message = "Server error"
				utils.SendError(w, http.StatusInternalServerError, error)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}
		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var book models.Book
		var error models.Error

		json.NewDecoder(r.Body).Decode(&book)

		if book.ID == 0 || book.Author == "" || book.Title == "" || book.Year == "" {
			error.Message = "All feilds required."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		bookRepo := bookRepository.BookRepository{}

		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)

	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)
		bookRepo := bookRepository.BookRepository{}
		id, _ := strconv.Atoi(params["id"])

		rowsDeleted, err := bookRepo.DeleteBook(db, id)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		if rowsDeleted == 0 {
			error.Message = "Not Found"
			utils.SendError(w, http.StatusNotFound, error)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsDeleted)
	}

}
