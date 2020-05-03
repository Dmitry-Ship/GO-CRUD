package books

import (
	"GO-CRUD/common"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	common.SendJSONresponse(Books, w)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	book := Book{}
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBook := CreateBook(book)
	common.SendJSONresponse(newBook, w)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	bookById := BookByIDRequest{}
	err := json.NewDecoder(r.Body).Decode(&bookById)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deletedBook, err := DeleteBook(bookById.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	common.SendJSONresponse(deletedBook, w)
}
