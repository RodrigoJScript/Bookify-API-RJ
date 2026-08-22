package transport

import (
	"bookify/service"
	"database/sql"
	"encoding/json"
	"net/http"
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
