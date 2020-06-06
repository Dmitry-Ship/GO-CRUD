package books

import "github.com/jinzhu/gorm"

type bookStorage struct {
	db *gorm.DB
}

type BookRepository interface {
	GetAllBooks(limit int) ([]Book, error)
	CreateBook(book Book) (Book, error)
	GetBookById(bookId string) (Book, error)
	DeleteBook(bookId string) (Book, error)
}

func (bs *bookStorage) GetAllBooks(limit int) ([]Book, error) {
	if limit == 0 {
		limit = 50
	}
	books := []Book{}

	err := bs.db.Limit(limit).Find(&books).Error

	return books, err
}

func (bs *bookStorage) CreateBook(book Book) (Book, error) {
	err := bs.db.Create(&book).Error

	return book, err
}

func (bs *bookStorage) GetBookById(bookId string) (Book, error) {
	book := Book{}
	err := bs.db.Where(`id = ?`, bookId).First(&book).Error

	return book, err
}

func (bs *bookStorage) DeleteBook(bookId string) (Book, error) {
	book := Book{}
	err := bs.db.Delete(&book, bookId).Error
	return book, err
}

func NewBooksStore(db *gorm.DB) BookRepository {
	return &bookStorage{db}
}
