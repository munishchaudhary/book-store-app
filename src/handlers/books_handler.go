package handlers

import (
    "context"
    "log"
    "time"

    "github.com/go-openapi/runtime/middleware"
    "github.com/munishchaudhary/book-store-app/src/db"
    "github.com/munishchaudhary/book-store-app/src/generated/restapi/operations"
    "github.com/munishchaudhary/book-store-app/src/validation"
)

type BookHandler struct {
	bookDB *db.BookDB
}

var globalBookHandler *BookHandler

func NewBookHandler(bookDB *db.BookDB) *BookHandler {
	return &BookHandler{bookDB: bookDB}
}

func SetGlobalBookHandler(handler *BookHandler) {
	globalBookHandler = handler
}

func CreateBookHandler(params operations.CreateBookParams) middleware.Responder {
    start := time.Now()
    log.Printf("[CreateBook] request received")
	if globalBookHandler == nil {
        log.Printf("[CreateBook] handler not initialized")
		return StatusInternalServerError("Handler not initialized")
	}

	if params.Book == nil {
        log.Printf("[CreateBook] missing book payload")
		return StatusBadRequest("Book data is required")
	}

	book := *params.Book

	// Validate book fields
	if validationErrors := validation.ValidateBook(&book); len(validationErrors) > 0 {
		errorMessage := validation.FormatValidationErrors(validationErrors)
        log.Printf("[CreateBook] validation failed: %s", errorMessage)
		return StatusBadRequest("Validation failed: " + errorMessage)
	}

	ctx := context.Background()

	err := globalBookHandler.bookDB.CreateBook(ctx, book)
	if err != nil {
        log.Printf("[CreateBook] db error: %v", err)
		return StatusInternalServerError("Failed to create book: " + err.Error())
	}

    log.Printf("[CreateBook] success in %s", time.Since(start))
	return StatusCreated("Book created successfully")
}

func GetBookHandler(params operations.GetBookParams) middleware.Responder {
    start := time.Now()
	if globalBookHandler == nil {
        log.Printf("[GetBook] handler not initialized")
		return StatusInternalServerError("Handler not initialized")
	}

	bookID := int(params.ID)
    log.Printf("[GetBook] id=%d", bookID)

	// Check if book exists
	exists, err := globalBookHandler.bookDB.BookExists(bookID)
	if err != nil {
        log.Printf("[GetBook] existence check error: %v", err)
		return StatusInternalServerError("Failed to check book existence: " + err.Error())
	}

	if !exists {
        log.Printf("[GetBook] not found id=%d", bookID)
		return StatusNotFound("Book not found")
	}

	book, err := globalBookHandler.bookDB.GetBookByID(bookID)
	if err != nil {
        log.Printf("[GetBook] db error: %v", err)
		return StatusInternalServerError("Failed to retrieve book: " + err.Error())
	}

    log.Printf("[GetBook] success in %s", time.Since(start))
	return StatusOK(book)
}

func GetBooksHandler(params operations.GetBooksParams) middleware.Responder {
    start := time.Now()
	if globalBookHandler == nil {
        log.Printf("[GetBooks] handler not initialized")
		return StatusInternalServerError("Handler not initialized")
	}

	books, err := globalBookHandler.bookDB.GetAllBooks()
	if err != nil {
        log.Printf("[GetBooks] db error: %v", err)
		return StatusInternalServerError("Failed to retrieve books: " + err.Error())
	}

    log.Printf("[GetBooks] success count=%d in %s", len(books), time.Since(start))
	return StatusOK(books)
}

func UpdateBookHandler(params operations.UpdateBookParams) middleware.Responder {
    start := time.Now()
	if globalBookHandler == nil {
        log.Printf("[UpdateBook] handler not initialized")
		return StatusInternalServerError("Handler not initialized")
	}

	if params.Book == nil {
        log.Printf("[UpdateBook] missing book payload")
		return StatusBadRequest("Book data is required")
	}

	book := *params.Book
	book.ID = params.ID
    log.Printf("[UpdateBook] id=%d", book.ID)

	// Validate book fields including ID
	if validationErrors := validation.ValidateBookForUpdate(&book); len(validationErrors) > 0 {
		errorMessage := validation.FormatValidationErrors(validationErrors)
        log.Printf("[UpdateBook] validation failed: %s", errorMessage)
		return StatusBadRequest("Validation failed: " + errorMessage)
	}

	// Check if book exists
	exists, err := globalBookHandler.bookDB.BookExists(int(book.ID))
	if err != nil {
        log.Printf("[UpdateBook] existence check error: %v", err)
		return StatusInternalServerError("Failed to check book existence: " + err.Error())
	}

	if !exists {
        log.Printf("[UpdateBook] not found id=%d", book.ID)
		return StatusNotFound("Book not found")
	}

	ctx := context.Background()

	updatedBook, err := globalBookHandler.bookDB.UpdateBook(ctx, &book)
	if err != nil {
        log.Printf("[UpdateBook] db error: %v", err)
		return StatusInternalServerError("Failed to update book: " + err.Error())
	}

    log.Printf("[UpdateBook] success in %s", time.Since(start))
	return StatusOK(updatedBook)
}

func DeleteBookHandler(params operations.DeleteBookParams) middleware.Responder {
    start := time.Now()
	if globalBookHandler == nil {
        log.Printf("[DeleteBook] handler not initialized")
		return StatusInternalServerError("Handler not initialized")
	}

	bookID := int(params.ID)
    log.Printf("[DeleteBook] id=%d", bookID)

	// Check if book exists
	exists, err := globalBookHandler.bookDB.BookExists(bookID)
	if err != nil {
        log.Printf("[DeleteBook] existence check error: %v", err)
		return StatusInternalServerError("Failed to check book existence: " + err.Error())
	}

	if !exists {
        log.Printf("[DeleteBook] not found id=%d", bookID)
		return StatusNotFound("Book not found")
	}

	ctx := context.Background()
	err = globalBookHandler.bookDB.DeleteBook(ctx, bookID)
	if err != nil {
        log.Printf("[DeleteBook] db error: %v", err)
		return StatusInternalServerError("Failed to delete book: " + err.Error())
	}

    log.Printf("[DeleteBook] success in %s", time.Since(start))
	return StatusOK("Book deleted successfully")
}
