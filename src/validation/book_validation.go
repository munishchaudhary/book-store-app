package validation

import (
	"fmt"
	"strings"
	"time"

	"github.com/munishchaudhary/book-store-app/src/generated/models"
)

const (
	MaxTitleLength  = 255
	MaxAuthorLength = 255
	MinYear         = 2000
)

// BookValidationError represents a validation error for book fields
type BookValidationError struct {
	Field   string
	Message string
}

func (e BookValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateBook validates a book model and returns validation errors
func ValidateBook(book *models.Book) []error {
	var errors []error

	// Validate title
	if err := validateTitle(book.Title); err != nil {
		errors = append(errors, err)
	}

	// Validate author
	if err := validateAuthor(book.Author); err != nil {
		errors = append(errors, err)
	}

	// Validate year
	if err := validateYear(book.Year); err != nil {
		errors = append(errors, err)
	}

	return errors
}

// ValidateBookForUpdate validates a book model for update operations
func ValidateBookForUpdate(book *models.Book) []error {
	var errors []error

	// Validate ID
	if err := validateID(book.ID); err != nil {
		errors = append(errors, err)
	}

	// Validate other fields
	errors = append(errors, ValidateBook(book)...)

	return errors
}

// validateTitle validates the book title
func validateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return BookValidationError{
			Field:   "title",
			Message: "title is required",
		}
	}

	if len(title) > MaxTitleLength {
		return BookValidationError{
			Field:   "title",
			Message: fmt.Sprintf("title cannot exceed %d characters", MaxTitleLength),
		}
	}

	return nil
}

// validateAuthor validates the book author
func validateAuthor(author string) error {
	if strings.TrimSpace(author) == "" {
		return BookValidationError{
			Field:   "author",
			Message: "author is required",
		}
	}

	if len(author) > MaxAuthorLength {
		return BookValidationError{
			Field:   "author",
			Message: fmt.Sprintf("author cannot exceed %d characters", MaxAuthorLength),
		}
	}

	return nil
}

// validateYear validates the book publication year
func validateYear(year int64) error {
	currentYear := int64(time.Now().Year())

	if year < MinYear {
		return BookValidationError{
			Field:   "year",
			Message: fmt.Sprintf("year cannot be earlier than %d", MinYear),
		}
	}

	if year > currentYear {
		return BookValidationError{
			Field:   "year",
			Message: fmt.Sprintf("year cannot be later than %d", currentYear),
		}
	}

	return nil
}

// validateID validates the book ID
func validateID(id int64) error {
	if id <= 0 {
		return BookValidationError{
			Field:   "id",
			Message: "id must be a positive integer",
		}
	}

	return nil
}

// FormatValidationErrors formats multiple validation errors into a single string
func FormatValidationErrors(errors []error) string {
	if len(errors) == 0 {
		return ""
	}

	var messages []string
	for _, err := range errors {
		if validationErr, ok := err.(BookValidationError); ok {
			messages = append(messages, validationErr.Error())
		} else {
			messages = append(messages, err.Error())
		}
	}

	return strings.Join(messages, "; ")
}
