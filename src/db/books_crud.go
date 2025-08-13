package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/munishchaudhary/book-store-app/src/generated/models"
)

type BookDB struct {
	db *sql.DB
}

// CreateBook inserts a new book into the database
func (b *BookDB) CreateBook(ctx context.Context, book models.Book) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to create db transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", book.Title, book.Author, book.Year)
	if err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit db transaction: %w", err)
	}

	return nil
}

// GetBookByID retrieves a single book from the database by its ID
func (b *BookDB) GetBookByID(id int) (*models.Book, error) {
	var book models.Book
	row := b.db.QueryRow("SELECT id, title, author, year FROM books WHERE id = ?", id)

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("book with ID %d not found", id)
		}
		return nil, fmt.Errorf("could not get book with ID %d: %w", id, err)
	}
	return &book, nil
}

// UpdateBook updates an existing book in the database.
func (b *BookDB) UpdateBook(ctx context.Context, book *models.Book) (*models.Book, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?", book.Title, book.Author, book.Year, book.ID)
	if err != nil {
		return nil, fmt.Errorf("could not update book with ID %d: %w", book.ID, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("could not commit transaction: %w", err)
	}

	return book, nil
}

// DeleteBook removes a book from the database by its ID using a transaction.
func (b *BookDB) DeleteBook(ctx context.Context, id int) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback()

	result, err := tx.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("could not delete book with ID %d: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not check rows affected for deletion: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("book with ID %d not found for deletion", id)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}

// GetAllBooks retrieves all books from the database
func (b *BookDB) GetAllBooks() ([]*models.Book, error) {
	rows, err := b.db.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		return nil, fmt.Errorf("could not get books: %w", err)
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, fmt.Errorf("could not scan book: %w", err)
		}
		books = append(books, &book)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over books: %w", err)
	}

	return books, nil
}

// BookExists checks if a book with the given ID exists in the database
func (b *BookDB) BookExists(id int) (bool, error) {
	var exists bool
	err := b.db.QueryRow("SELECT EXISTS(SELECT 1 FROM books WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("could not check if book exists: %w", err)
	}
	return exists, nil
}
