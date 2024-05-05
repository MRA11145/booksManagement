package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/controller"
	"main.go/database"
)

func main() {
	
	database.Init()
	r := mux.NewRouter()

	r.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")
	r.HandleFunc("/api/books", controller.CreateBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controller.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}