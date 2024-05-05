package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/database"
	"main.go/model"
)

// get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database.GetBooks())
}

// get single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	books := database.GetBooks()
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Book{})
}

// create a new book
func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(1000000))
	database.AppendBook(book)
	json.NewEncoder(w).Encode(book)
}

// udpate book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	database.DeleteBook(params["id"])
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(1000000))
	database.AppendBook(book)
	json.NewEncoder(w).Encode(database.GetBooks())
}

// delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	database.DeleteBook(params["id"])
	json.NewEncoder(w).Encode(database.GetBooks())
}
