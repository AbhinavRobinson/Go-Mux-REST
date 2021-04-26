package helpers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	model "github.com/abhinavrobinson/Go-Mux-REST/app/models"
)

// Create new book
func CreateBook(w http.ResponseWriter, r *http.Request, books []model.Book) []model.Book {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	return books
}
