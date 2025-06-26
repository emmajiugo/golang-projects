package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/emmajiugo/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name       string `json:"name"`
	Author    string `json:"author"`
	Publication string `json:"publication"`
}

// Initialize the database connection
func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook adds a new book to the database
func CreateBook(book *Book) (*Book, error) {
	// Ensure ID is 0 for new records
    book.ID = 0
    book.CreatedAt = time.Time{} // Let GORM set timestamps
    book.UpdatedAt = time.Time{}

	if err := db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// GetBooks retrieves all books from the database
func GetBooks() ([]Book, error) {
	var books []Book
	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookById retrieves a book by its ID
func GetBookById(id uint) (*Book, error) {
	var book Book
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// UpdateBook updates an existing book in the database
func UpdateBook(book *Book) (*Book, error) {
	if err := db.Save(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// DeleteBook removes a book from the database
func DeleteBook(id uint) error {
	if err := db.Delete(&Book{}, id).Error; err != nil {
		return err
	}
	return nil
}