package helpers

import (
	"encoding/json"
	"net/http"

	model "github.com/abhinavrobinson/Go-Mux-REST/app/models"
	"github.com/gorilla/mux"
)

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request, books []model.Book) []model.Book {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
	return books
}
