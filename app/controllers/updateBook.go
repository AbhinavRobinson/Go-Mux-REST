package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/abhinavrobinson/Go-Mux-REST/app/models"
	"github.com/gorilla/mux"
)

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request, books []models.Book) []models.Book {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, book := range books {
		if book.ID == params["id"] {
			// remove old entry
			books = append(books[:index], books[index+1:]...)
			// get new details
			var book models.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = strconv.Itoa(rand.Intn(10000)) // Mock ID - not safe
			// add book
			books = append(books, book)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
	return books
}
