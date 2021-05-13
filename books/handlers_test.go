package books

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockService struct{}

func (mr *MockService) GetAllBooks(limit int) ([]Book, error) {
	return []Book{}, nil
}

func (mr *MockService) CreateBook(book Book) (Book, error) {
	return book, nil
}

func (mr *MockService) GetBookById(bookId int) (Book, error) {
	book := Book{
		ID:       1,
		Title:    "hello",
		Author:   "John Doe",
		ISBN:     "123",
		Language: "English",
	}

	return book, nil
}

func (mr *MockService) DeleteBook(bookId int) (Book, error) {
	return Book{}, nil
}

var mockService = &MockService{}

func TestGetBooksHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBooks(mockService))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "handler returned wrong status code")
	assert.Equal(t, "[]", strings.TrimSpace(rr.Body.String()), "handler returned unexpected body")
}

func TestAddBookHandler(t *testing.T) {
	book := Book{
		ID:          1,
		Title:       "hello",
		Author:      "John Doe",
		ISBN:        "123",
		Language:    "English",
		PublishedAt: "hello",
	}

	b, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := strings.NewReader(string(b))

	req, err := http.NewRequest("POST", "/books/create", reader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AddBook(mockService))

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "handler returned wrong status code")
	assert.Equal(t, string(b), strings.TrimSpace(rr.Body.String()), "handler returned unexpected body")
}
