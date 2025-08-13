package validation

import (
	"strings"
	"testing"
	"time"

	"github.com/munishchaudhary/book-store-app/src/generated/models"
)

func TestValidateBook(t *testing.T) {
	currentYear := int64(time.Now().Year())

	tests := []struct {
		name    string
		book    *models.Book
		wantErr bool
	}{
		{
			name: "valid book",
			book: &models.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2020,
			},
			wantErr: false,
		},
		{
			name: "empty title",
			book: &models.Book{
				Title:  "",
				Author: "Test Author",
				Year:   2020,
			},
			wantErr: true,
		},
		{
			name: "whitespace title",
			book: &models.Book{
				Title:  "   ",
				Author: "Test Author",
				Year:   2020,
			},
			wantErr: true,
		},
		{
			name: "title too long",
			book: &models.Book{
				Title:  strings.Repeat("a", MaxTitleLength+1),
				Author: "Test Author",
				Year:   2020,
			},
			wantErr: true,
		},
		{
			name: "empty author",
			book: &models.Book{
				Title:  "Test Book",
				Author: "",
				Year:   2020,
			},
			wantErr: true,
		},
		{
			name: "author too long",
			book: &models.Book{
				Title:  "Test Book",
				Author: strings.Repeat("a", MaxAuthorLength+1),
				Year:   2020,
			},
			wantErr: true,
		},
		{
			name: "year too early",
			book: &models.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   MinYear - 1,
			},
			wantErr: true,
		},
		{
			name: "year too late",
			book: &models.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   currentYear + 1,
			},
			wantErr: true,
		},
		{
			name: "future year",
			book: &models.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   currentYear + 1,
			},
			wantErr: true,
		},
		{
			name: "current year",
			book: &models.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   currentYear,
			},
			wantErr: false,
		},
		{
			name: "past year",
			book: &models.Book{
				Title:  "Test Book",
				Author: "Test Author",
				Year:   currentYear - 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := ValidateBook(tt.book)
			if tt.wantErr && len(errors) == 0 {
				t.Errorf("ValidateBook() expected errors but got none")
			}
			if !tt.wantErr && len(errors) > 0 {
				t.Errorf("ValidateBook() unexpected errors: %v", errors)
			}
		})
	}
}

func TestValidateBookForUpdate(t *testing.T) {
	tests := []struct {
		name    string
		book    *models.Book
		wantErr bool
	}{
		{
			name: "valid book for update",
			book: &models.Book{
				ID:     1,
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2020,
			},
			wantErr: false,
		},
		{
			name: "invalid ID",
			book: &models.Book{
				ID:     0,
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2020,
			},
			wantErr: true,
		},
		{
			name: "negative ID",
			book: &models.Book{
				ID:     -1,
				Title:  "Test Book",
				Author: "Test Author",
				Year:   2020,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errors := ValidateBookForUpdate(tt.book)
			if tt.wantErr && len(errors) == 0 {
				t.Errorf("ValidateBookForUpdate() expected errors but got none")
			}
			if !tt.wantErr && len(errors) > 0 {
				t.Errorf("ValidateBookForUpdate() unexpected errors: %v", errors)
			}
		})
	}
}

func TestFormatValidationErrors(t *testing.T) {
	errors := []error{
		BookValidationError{Field: "title", Message: "title is required"},
		BookValidationError{Field: "author", Message: "author is required"},
	}

	formatted := FormatValidationErrors(errors)
	expected := "title: title is required; author: author is required"

	if formatted != expected {
		t.Errorf("FormatValidationErrors() = %v, want %v", formatted, expected)
	}
}

func TestBookValidationError_Error(t *testing.T) {
	err := BookValidationError{
		Field:   "title",
		Message: "title is required",
	}

	expected := "title: title is required"
	if err.Error() != expected {
		t.Errorf("BookValidationError.Error() = %v, want %v", err.Error(), expected)
	}
}
