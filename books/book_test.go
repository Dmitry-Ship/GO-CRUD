package books

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {

	book := Book{
		Title:  "hello",
		Author: "John Doe",
		ISBN:   "123",
	}

	expectedBook := Book{
		Title:  "hello",
		Author: "John Doe",
		ISBN:   "123",
		ID:     "498081",
	}

	result := CreateBook(book)

	assert.Equal(t, result, expectedBook, "The two books should be the same.")
}

func TestGetBookById(t *testing.T) {
	book := Book{
		Title:  "hello",
		Author: "John Doe",
		ISBN:   "123",
	}

	createdBook := CreateBook(book)

	foundBook, _, _ := GetBookById(createdBook.ID)

	assert.Equal(t, createdBook, foundBook, "The two words should be the same.")
}

func TestGetBookByIdNotFound(t *testing.T) {
	_, _, err := GetBookById("123")

	assert.Equal(t, err, errors.New("book not found"), "Shuuld return error if book not found.")
}

func TestDeleteBook(t *testing.T) {
	book := Book{
		Title:  "hello",
		Author: "John Doe",
		ISBN:   "123",
	}

	createdBook := CreateBook(book)

	deletedBook, err := DeleteBook(createdBook.ID)

	if err != nil {
		t.Fatalf("book not found")
	}

	assert.Equal(t, createdBook, deletedBook, "The two words should be the same.")
}

func TestDeleteBookNotFound(t *testing.T) {
	_, err := DeleteBook("123")

	assert.Equal(t, err, errors.New("Book not found"), "The two words should be the same.")
}
