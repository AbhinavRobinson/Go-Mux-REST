package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	store "github.com/abhinavrobinson/Go-Mux-REST/app/controllers"
	"github.com/abhinavrobinson/Go-Mux-REST/app/models"
	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - !todo - add database
	store.Books = append(store.Books, models.Book{ID: "1", Isbn: "0001", Title: "Default Book A", Author: &models.Author{FirstName: "John", LastName: "Doe"}})
	store.Books = append(store.Books, models.Book{ID: "2", Isbn: "0002", Title: "Default Book B", Author: &models.Author{FirstName: "Kevin", LastName: "Smith"}})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/books", store.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", store.GetBook).Methods("GET")
	r.HandleFunc("/api/books", store.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", store.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", store.DeleteBook).Methods("DELETE")

	// Serve to port
	const port int = 8000
	fmt.Printf("Starting Server on port %d ...", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
