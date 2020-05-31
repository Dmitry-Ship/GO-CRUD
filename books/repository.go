package books

import "github.com/jinzhu/gorm"

type BookStorage struct {
	db *gorm.DB
}

type BookRepository interface {
	GetAllBooks(limit int) []Book
	CreateBook(book Book) Book
	GetBookById(bookId string) Book
	DeleteBook(bookId string) Book
}

func (bs *BookStorage) GetAllBooks(limit int) ([]Book, error) {
	if limit == 0 {
		limit = 50
	}
	books := []Book{}

	err := bs.db.Limit(limit).Find(&books).Error

	return books, err
}

func (bs *BookStorage) CreateBook(book Book) (Book, error) {
	err := bs.db.Create(&book).Error

	return book, err
}

func (bs *BookStorage) GetBookById(bookId string) (Book, error) {
	book := Book{}
	err := bs.db.Where(`id = ?`, bookId).First(&book).Error

	return book, err
}

func (bs *BookStorage) DeleteBook(bookId string) (Book, error) {
	book := Book{}
	err := bs.db.Delete(&book, bookId).Error
	return book, err
}

func NewBooksStore(db *gorm.DB) BookStorage {
	return BookStorage{db}
}