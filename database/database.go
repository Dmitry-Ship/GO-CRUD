package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type dbSettings struct {
	User     string `json:"user"`
	Database string `json:"dbname"`
	Port     int    `json:"port"`
	Host     string `json:"host"`
}

var DbConfig = dbSettings{
	User:     "dmitryshipunov",
	Database: "dmitryshipunov",
	Port:     5432,
	Host:     "localhost",
}

func GetDatabaseConnection() *gorm.DB {
	options := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", DbConfig.Host, DbConfig.Port, DbConfig.User, DbConfig.Database)
	db, err := gorm.Open("postgres", options)
	if err != nil {
		panic("Could not connect to database")
	}
	fmt.Println(fmt.Sprintf("Connected to database %s", DbConfig.Database))

	// Migrate the schema
	// db.AutoMigrate(&books.Book{})
	// db.AutoMigrate(&proxy.Changelog{})

	return db
}
