package books

type Service interface {
	GetAllBooks(limit int) ([]Book, error)
	CreateBook(book Book) (Book, error)
	GetBookById(bookId int) (Book, error)
	DeleteBook(bookId int) (Book, error)
}

type service struct {
	repo BookRepository
}

func (s *service) GetAllBooks(limit int) ([]Book, error) {
	books, err := s.repo.GetAllBooks(limit)

	return books, err
}

func (s *service) CreateBook(book Book) (Book, error) {
	book, err := s.repo.CreateBook(book)

	return book, err
}

func (s *service) GetBookById(bookId int) (Book, error) {
	book, err := s.repo.GetBookById(bookId)

	return book, err
}

func (s *service) DeleteBook(bookId int) (Book, error) {
	book, err := s.repo.DeleteBook(bookId)

	return book, err
}

func NewService(repo BookRepository) Service {
	return &service{repo}
}
