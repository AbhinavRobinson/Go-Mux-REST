package controllers

import (
	"net/http"

	helper "github.com/abhinavrobinson/Go-Mux-REST/app/controllers/helpers"
	model "github.com/abhinavrobinson/Go-Mux-REST/app/models"
)

// Init Book slice
var Books []model.Book

// ------ DISPATCHERS ------

// Get All Books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	helper.GetBooks(w, r, Books)
}

// Get one books
func GetBook(w http.ResponseWriter, r *http.Request) {
	helper.GetBook(w, r, Books)
}

// Create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	Books = helper.CreateBook(w, r, Books)
}

// Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	Books = helper.UpdateBook(w, r, Books)
}

// Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	Books = helper.DeleteBook(w, r, Books)
}
