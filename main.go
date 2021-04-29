package main

import (
	"GO-CRUD/database"
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

	db := database.GetDatabaseConnection()
	defer db.Close()

	HandleRequests(db)

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Println("Listening to: http://" + host + ":" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
