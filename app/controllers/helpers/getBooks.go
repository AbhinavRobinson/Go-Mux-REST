package helpers

import (
	"encoding/json"
	"net/http"

	model "github.com/abhinavrobinson/Go-Mux-REST/app/models"
)

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request, books []model.Book) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
