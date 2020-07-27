package books

import (
	"GO-CRUD/common"
	"encoding/json"
	"fmt"
	"net/http"
)

type BookByIDRequest struct {
	ID string `json:"id"`
}

func GetBooks(bookService Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := bookService.GetAllBooks(5)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		common.SendJSONresponse(result, w)
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
