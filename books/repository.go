package books

import "gorm.io/gorm"

type BookStorage struct {
	db *gorm.DB
}

type BookRepository interface {
	GetAllBooks(limit int) ([]Book, error)
	CreateBook(book Book) (Book, error)
	GetBookById(bookId int) (Book, error)
	DeleteBook(bookId int) (Book, error)
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

func (bs *BookStorage) GetBookById(bookId int) (Book, error) {
	book := Book{}
	err := bs.db.Where(`id = ?`, bookId).First(&book).Error

	return book, err
}

func (bs *BookStorage) DeleteBook(bookId int) (Book, error) {
	book := Book{}
	err := bs.db.Delete(&book, bookId).Error
	return book, err
}

func NewBooksRepository(db *gorm.DB) BookRepository {
	return &BookStorage{db}
}
