package books

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Repository *BookStorage
var Mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	Mock = mock

	gdb, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	Repository = &BookStorage{db: gdb}
	os.Exit(m.Run())
}

func TestGetBookById(t *testing.T) {
	book := Book{
		ID:       1,
		Title:    "hello",
		Author:   "John Doe",
		ISBN:     "123",
		Language: "English",
	}

	rows := sqlmock.
		NewRows([]string{"id", "title", "author", "isbn", "language"}).
		AddRow(book.ID, book.Title, book.Author, book.ISBN, book.Language)

	const sqlSelectOne = `SELECT * FROM "books" WHERE "books"."id" = $1`
	Mock.ExpectQuery(sqlSelectOne).WithArgs(book.ID).WillReturnRows(rows)

	result, err := Repository.GetBookById(book.ID)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, book, result, "The two books should be the same.")
}

func TestGetAllBooks(t *testing.T) {
	const sqlSelectAll = `SELECT * FROM "books" LIMIT 50`

	Mock.ExpectQuery(sqlSelectAll).
		WillReturnRows(sqlmock.NewRows(nil))

	result, err := Repository.GetAllBooks(0)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, 0, len(result), "The two words should be the same.")
}

func TestCreateBook(t *testing.T) {

	book := Book{
		ID:          1,
		Title:       "hello",
		Author:      "John Doe",
		ISBN:        "123",
		Language:    "English",
		PublishedAt: "hello",
	}

	const sqlInsert = `INSERT INTO "books" ("title","author","isbn","published_at","language","id")
                                        VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`
	const newId = 1
	Mock.ExpectBegin() // begin transaction
	Mock.ExpectQuery(sqlInsert).
		WithArgs(book.Title, book.Author, book.ISBN, book.PublishedAt, book.Language, book.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	Mock.ExpectCommit() // commit transaction

	result, err := Repository.CreateBook(book)

	assert.Equal(t, nil, err, "Error occurred")
	assert.Equal(t, book, result, "The two words should be the same.")
}

// func TestDeleteBookNotFound(t *testing.T) {

// 	book := Book{
// 		ID:          1,
// 		Title:       "hello",
// 		Author:      "John Doe",
// 		ISBN:        "123",
// 		Language:    "English",
// 		PublishedAt: "hello",
// 	}

// 	Repository.CreateBook(book)
// 	const sqlInsert = `DELETE FROM "books"  WHERE ("books"."id" = $1)`
// 	const newId = 1
// 	Mock.ExpectBegin() // begin transaction
// 	Mock.ExpectQuery(sqlInsert).
// 		WithArgs(book.ID).
// 		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
// 	Mock.ExpectCommit() // commit transaction

// 	result, err := Repository.DeleteBook(book.ID)

// 	assert.Equal(t, nil, err, "Error occurred")
// 	assert.Equal(t, book, result, "The two words should be the same.")
// }
