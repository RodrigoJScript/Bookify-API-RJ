package service

import (
	"bookify/model"
	"bookify/store"
	"database/sql"
)

func GetBooks(db *sql.DB) ([]model.Book, error) {
	books, err := store.GetBooks(db)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func SeedBooks(db *sql.DB) error {
	return store.SeedBooks(db)
}

func GetBookById(db *sql.DB, bookID int) (model.Book, error) {
	return store.GetBookById(db, bookID)
}

func CreateBook(db *sql.DB, book model.Book) (model.Book, error) {
	return store.CreateBook(db, book)
}

func DeleteBookById(db *sql.DB, bookID int) error {
	return store.DeleteBookById(db, bookID)
}

func UpdateBookById(db *sql.DB, bookID int, book model.Book) (model.Book, error) {
	return store.UpdateBookById(db, bookID, book)
}
