package infrastructure

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")

	options := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := gorm.Open(driver, options)
	if err != nil {
		panic("Could not connect to database")
	}
	fmt.Println(fmt.Sprintf("Connected to database %s", dbname))

	// Migrate the schema
	// db.AutoMigrate(&books.Book{})
	// db.AutoMigrate(&proxy.Changelog{})

	return db
}
