package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/abhinavrobinson/Go-Mux-REST/app/models"
)

// Get all books
func GetBooks(w http.ResponseWriter, r *http.Request, books []models.Book) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
