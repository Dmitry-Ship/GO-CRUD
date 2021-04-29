package main

import (
	"GO-CRUD/books"
	"net/http"

	"github.com/jinzhu/gorm"
)

func HandleRequests(db *gorm.DB) {
	booksStore := books.NewBooksStore(db)
	booksService := books.NewService(booksStore)

	http.HandleFunc("/api/books", books.GetBooks(booksService))
	http.HandleFunc("/api/books/create", books.AddBook(booksService))
	http.HandleFunc("/api/books/delete", books.RemoveBook(booksService))
}
