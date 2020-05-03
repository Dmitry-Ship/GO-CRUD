package books

import (
	"GO-CRUD/common"
	"encoding/json"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	common.SendJSONresponse(Books, w)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	book := Book{}
	err = json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBook := CreateBook(book)
	common.SendJSONresponse(newBook, w)
}
