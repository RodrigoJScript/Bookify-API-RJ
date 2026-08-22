package main

import (
	"bookify/store"
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := store.SeedBooks(db); err != nil {
		log.Fatal(err)
	}
}
