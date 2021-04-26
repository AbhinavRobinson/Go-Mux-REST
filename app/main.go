package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	controller "github.com/abhinavrobinson/Go-Mux-REST/app/controllers"
	model "github.com/abhinavrobinson/Go-Mux-REST/app/models"
	"github.com/gorilla/mux"
)

func main() {
	// Init Router
	router := mux.NewRouter()

	// !TODO - add database
	controller.Books = append(controller.Books, model.Book{ID: "1", Isbn: "0001", Title: "Default Book A", Author: &model.Author{FirstName: "John", LastName: "Doe"}})
	controller.Books = append(controller.Books, model.Book{ID: "2", Isbn: "0002", Title: "Default Book B", Author: &model.Author{FirstName: "Kevin", LastName: "Smith"}})

	// Route Handlers / Endpoints
	router.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")
	router.HandleFunc("/api/books", controller.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controller.DeleteBook).Methods("DELETE")

	// Serve to port
	const port int = 8000
	fmt.Printf("Starting Server on port %d ...", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}
