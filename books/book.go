package books

import (
	"math/rand"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
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
