package store

import (
	"bookify/model"
	"database/sql"

	_ "modernc.org/sqlite"
)

func SeedBooks(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			description TEXT,
			price REAL NOT NULL,
			quantity INTEGER NOT NULL
		);
		INSERT INTO books (title, author, description, price, quantity) VALUES
			('El Quijote', 'Miguel de Cervantes', 'Novela de caballerías', 19.99, 12),
			('Cien años de soledad', 'Gabriel García Márquez', 'Realismo mágico', 24.50, 8),
			('Rayuela', 'Julio Cortázar', 'Novela experimental', 18.00, 5);
	`)
	return err
}

func GetBooks(db *sql.DB) ([]model.Book, error) {
	var books []model.Book
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price, &book.Quantity)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
