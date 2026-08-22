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
