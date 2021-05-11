package main

import (
	"GO-CRUD/books"
	"GO-CRUD/infrastructure"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := infrastructure.GetDatabaseConnection()
	defer db.Close()

	booksRepository := books.NewBooksRepository(db)
	booksService := books.NewService(booksRepository)
	books.HandleRequests(booksService)

	port := os.Getenv("PORT")

	fmt.Println("Listening to port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
