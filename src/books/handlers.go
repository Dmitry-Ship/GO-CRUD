package books

import (
	"GO-CRUD/common"
	"encoding/json"
	"fmt"
	"net/http"
)

type BookByIDRequest struct {
	ID int `json:"id"`
}

func HandleRequests(booksService Service) {
	http.HandleFunc("/api/books", GetBooks(booksService))
	http.HandleFunc("/api/books/create", AddBook(booksService))
	http.HandleFunc("/api/books/delete", RemoveBook(booksService))
}

type BooksResponse struct {
	Data []Book `json:"data"`
}

func GetBooks(bookService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := bookService.GetAllBooks(50)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		resp := BooksResponse{Data: result}

		common.SendJSONresponse(resp, w)
	}
}

func AddBook(bookService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		book := Book{}
		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newBook, _ := bookService.CreateBook(book)
		common.SendJSONresponse(newBook, w)
	}
}

func RemoveBook(bookService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bookById := BookByIDRequest{}
		err := json.NewDecoder(r.Body).Decode(&bookById)

		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		deletedBook, deletionError := bookService.DeleteBook(bookById.ID)

		if deletionError != nil {
			http.Error(w, deletionError.Error(), http.StatusBadRequest)
			return
		}

		common.SendJSONresponse(deletedBook, w)
	}
}
