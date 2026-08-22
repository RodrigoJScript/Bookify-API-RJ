package main

import (
	"bookify/transport"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	db, err := sql.Open("sqlite", "store/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		transport.GetBooks(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
