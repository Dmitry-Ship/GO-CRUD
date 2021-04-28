package main

import (
	"GO-CRUD/books"
	"net/http"

	"github.com/jinzhu/gorm"
)

func HandleRequests(db *gorm.DB) {
	booksStore := books.NewBooksStore(db)
	booksService := books.NewService(booksStore)

	http.HandleFunc("/books", books.GetBooks(booksService))
	http.HandleFunc("/books/create", books.AddBook(booksService))
	http.HandleFunc("/books/delete", books.RemoveBook(booksService))
}
