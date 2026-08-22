package store

import (
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
 mar price REAL NOT NULL,
			quantity INTEGER NOT NULL
		);
		INSERT INTO books (title, author, description, price, quantity) VALUES
			('El Quijote', 'Miguel de Cervantes', 'Novela de caballerías', 19.99, 12),
			('Cien años de soledad', 'Gabriel García Márquez', 'Realismo mágico', 24.50, 8),
			('Rayuela', 'Julio Cortázar', 'Novela experimental', 18.00, 5);
	`)
	return err
}
