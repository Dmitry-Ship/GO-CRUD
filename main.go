package main

import (
	"GO-CRUD/database"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db := database.GetDatabaseConnection()
	defer db.Close()

	HandleRequests(db)

	fmt.Println("Listening to: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
