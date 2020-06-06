package main

import (
	"GO-CRUD/books"
	"GO-CRUD/proxy"
	"net/http"

	"github.com/jinzhu/gorm"
)

func HandleRequests(db *gorm.DB) {
	http.HandleFunc("/proxy", proxy.Proxy)

	booksStore := books.NewBooksStore(db)
	booksService := books.NewService(booksStore)

	http.HandleFunc("/books", books.GetBooks(booksService))
	http.HandleFunc("/books/create", books.AddBook(booksService))
	http.HandleFunc("/books/delete", books.RemoveBook(booksService))
}
