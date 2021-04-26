package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/abhinavrobinson/Go-Mux-REST/app/controllers"
	"github.com/abhinavrobinson/Go-Mux-REST/app/models"
	"github.com/gorilla/mux"
)

// Init Book slice
var books []models.Book

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - !todo - add database
	books = append(books, models.Book{ID: "1", Isbn: "0001", Title: "Default Book A", Author: &models.Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, models.Book{ID: "2", Isbn: "0002", Title: "Default Book B", Author: &models.Author{FirstName: "Kevin", LastName: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Serve to port
	const port int = 8000
	fmt.Printf("Starting Server on port %d ...", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}

// ------ DISPATCHERS ------

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	controllers.GetBooks(w, r, books)
}

// Get one books
func getBook(w http.ResponseWriter, r *http.Request) {
	controllers.GetBook(w, r, books)
}

// Create new book
func createBook(w http.ResponseWriter, r *http.Request) {
	controllers.CreateBook(w, r, books)
}

// Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	controllers.UpdateBook(w, r, books)
}

// Delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	controllers.DeleteBook(w, r, books)
}
