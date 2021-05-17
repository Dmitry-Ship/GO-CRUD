package main

import (
	"GO-CRUD/books"
	"GO-CRUD/infrastructure"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := infrastructure.GetDatabaseConnection()
	// defer db.Close()

	booksRepository := books.NewBooksRepository(db)
	booksService := books.NewService(booksRepository)
	books.HandleRequests(booksService)

	port := os.Getenv("PORT")

	fmt.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
