package helpers

import (
	"encoding/json"
	"net/http"

	model "github.com/abhinavrobinson/Go-Mux-REST/app/models"
	"github.com/gorilla/mux"
)

// Get one books
func GetBook(w http.ResponseWriter, r *http.Request, books []model.Book) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find id
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Book{ID: "-1", Isbn: "", Title: "ERR: Book Not Found", Author: &model.Author{}})
}
