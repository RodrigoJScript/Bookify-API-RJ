package store

import (
	"bookify/model"
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"
)

var createBooksTable = `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		description TEXT,
		price REAL NOT NULL,
		quantity INTEGER NOT NULL
	);
`

var insertBooks = `
	INSERT INTO books (title, author, description, price, quantity) VALUES ('El Quijote', 'Miguel de Cervantes', 'Novela de caballerías', 19.99, 12), ('Cien años de soledad', 'Gabriel García Márquez', 'Realismo mágico', 24.50, 8), ('Rayuela', 'Julio Cortázar', 'Novela experimental', 18.00, 5);
`

func SeedBooks(db *sql.DB) error {
	hasTable, err := HasTable(db, "books")
	if err != nil {
		return err
	}
	if !hasTable {
		_, err = db.Exec(createBooksTable)
		if err != nil {
			return err
		}
	}

	hasBooks, err := GetBooks(db)
	if err != nil {
		return err
	}
	if len(hasBooks) == 0 {
		_, err = db.Exec(insertBooks)
		return err
	}
	return nil
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

func HasTable(db *sql.DB, tableName string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type = 'table' AND name = ?", tableName).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetBookById(db *sql.DB, bookID int) (model.Book, error) {
	var book model.Book
	err := db.QueryRow("SELECT * FROM books WHERE id = ?", bookID).Scan(&book.ID, &book.Title, &book.Author, &book.Description, &book.Price, &book.Quantity)
	if err != nil {
		return model.Book{}, err
	}
	return book, nil
}

func CreateBook(db *sql.DB, book model.Book) (model.Book, error) {
	result, err := db.Exec("INSERT INTO books (title, author, description, price, quantity) VALUES (?, ?, ?, ?, ?)", book.Title, book.Author, book.Description, book.Price, book.Quantity)
	if err != nil {
		return model.Book{}, err
	}
	bookID, err := result.LastInsertId()
	if err != nil {
		return model.Book{}, err
	}
	book.ID = int(bookID)
	return book, nil
}

func DeleteBookById(db *sql.DB, bookID int) error {
	result, err := db.Exec("DELETE FROM books WHERE id = ?", bookID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}

func UpdateBookById(db *sql.DB, bookID int, book model.Book) (model.Book, error) {
	result, err := db.Exec("UPDATE books SET title = ?, author = ?, description = ?, price = ?, quantity = ? WHERE id = ?", book.Title, book.Author, book.Description, book.Price, book.Quantity, bookID)
	if err != nil {
		return model.Book{}, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.Book{}, err
	}
	if rowsAffected == 0 {
		return model.Book{}, errors.New("book not found")
	}
	return book, nil
}
