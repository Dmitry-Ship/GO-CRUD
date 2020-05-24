package books

import (
	"errors"
	"math/rand"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type BookByIDRequest struct {
	ID string `json:"id"`
}

type Author struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var Books []Book

func CreateBook(book Book) Book {
	book.ID = strconv.Itoa(rand.Intn(1000000))

	Books = append(Books, book)

	return book
}

func GetBookById(bookId string) (Book, int, error) {
	for i, book := range Books {
		if book.ID == bookId {
			return book, i, nil
		}
	}
	return Book{}, 0, errors.New("book not found")
}

func DeleteBook(bookId string) (Book, error) {
	book, i, err := GetBookById(bookId)

	if err != nil {
		return Book{}, errors.New("Book not found")
	}

	Books = append(Books[:i], Books[i+1:]...)
	return book, nil
}
