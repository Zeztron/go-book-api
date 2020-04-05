package main

/*
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
*/

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Get All Books
func getBooks(w http.ResponseWriter, req *http.Request) {

}

// Get Book
func getBook(w http.ResponseWriter, req *http.Request) {

}

// Create Book
func createBook(w http.ResponseWriter, req *http.Request) {

}

// Update Book
func updateBook(w http.ResponseWriter, req *http.Request) {

}

// Delete Book
func deleteBook(w http.ResponseWriter, req *http.Request) {

}

func main() {
	// Initialize Router
	router := mux.NewRouter()

	// Mock Data

	// Route Handler / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Start Server
	log.Fatal(http.ListenAndServe(":8000", router))
}
