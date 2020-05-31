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

	http.HandleFunc("/books", books.GetBooks(booksStore))
	http.HandleFunc("/books/create", books.AddBook(booksStore))
	http.HandleFunc("/books/delete", books.RemoveBook(booksStore))
}
