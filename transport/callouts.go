package transport

import (
	"bookify/model"
	"bookify/service"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetBooks(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	books, err := service.GetBooks(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBookById(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	bookID := r.PathValue("id")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	book, err := service.GetBookById(db, bookIDInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	book := model.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if book.Title == "" || book.Author == "" || book.Price < 0 || book.Quantity < 0 {
		http.Error(w, "Invalid book data: title, author, price, and quantity are required", http.StatusBadRequest)
		return
	}
	createdBook, err := service.CreateBook(db, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBook)
}
