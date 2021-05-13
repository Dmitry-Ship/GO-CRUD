package books

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepo struct{}

func (mr *MockRepo) GetAllBooks(limit int) ([]Book, error) {
	return []Book{}, nil
}

func (mr *MockRepo) CreateBook(book Book) (Book, error) {
	return book, nil
}

func (mr *MockRepo) GetBookById(bookId int) (Book, error) {
	book := Book{
		ID:       1,
		Title:    "hello",
		Author:   "John Doe",
		ISBN:     "123",
		Language: "English",
	}

	return book, nil
}

func (mr *MockRepo) DeleteBook(bookId int) (Book, error) {
	return Book{}, nil
}

var mockedRepository = &MockRepo{}
var bookService = NewService(mockedRepository)

func TestGetAllBooksService(t *testing.T) {
	result, err := bookService.GetAllBooks(0)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, 0, len(result), "The two words should be the same.")
}

func TestCreateBookService(t *testing.T) {
	book := Book{
		ID:          1,
		Title:       "hello",
		Author:      "John Doe",
		ISBN:        "123",
		Language:    "English",
		PublishedAt: "hello",
	}

	result, err := bookService.CreateBook(book)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, book, result, "The two words should be the same.")
}

func TestGetBookByIdService(t *testing.T) {
	book := Book{
		ID:       1,
		Title:    "hello",
		Author:   "John Doe",
		ISBN:     "123",
		Language: "English",
	}

	result, err := bookService.GetBookById(book.ID)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, book, result, "The two books should be the same.")
}
