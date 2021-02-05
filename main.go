package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Book declare structure (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author declare structure
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slice Book struct

var books []Book // --> a variable type slice of type struct book.

// Get All books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))  // Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
		}
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		book.ID = params["id"]  // Mock ID
		books = append(books, book)
		json.NewEncoder(w).Encode(book)
		return
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append(books[:index], books[index+1:]...)
		}
		break
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// init Router
	r := mux.NewRouter()

	// mock data
	books = append(books, Book{ID: "1", Isbn: "546452", Title: "Book Of Wonder", Author: &Author{Firstname: "Keanu", Lastname: "Reeves"}})
	books = append(books, Book{ID: "2", Isbn: "235437", Title: "Book Of Sadness", Author: &Author{Firstname: "James", Lastname: "Kirk"}})
	// Route Handlers/ End points

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
