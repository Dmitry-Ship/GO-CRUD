package main

import (
	"GO-CRUD/books"
	"GO-CRUD/proxy"
	"fmt"
	"log"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/proxy", proxy.Proxy)

	http.HandleFunc("/books", books.GetBooks)
	http.HandleFunc("/books/create", books.AddBook)
	http.HandleFunc("/books/delete", books.RemoveBook)

	fmt.Println("Listening to: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
