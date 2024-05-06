package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"main.go/controller"
	"main.go/database"
)

func main() {
	// Initialize database
	database.Init()

	// Create a new mux router
	r := mux.NewRouter()

	// Serve static files from the "./web/static" directory
	fs := http.FileServer(http.Dir("./web/static/"))
	r.PathPrefix("/web/static/").Handler(http.StripPrefix("/web/static/", fs))

	// Handle requests for the root URL ("/") by serving the index.html file
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/static/index.html")
	})

	// Define routes for the RESTful API
	r.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", controller.GetBook).Methods("GET")
	r.HandleFunc("/api/books", controller.CreateBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", controller.DeleteBook).Methods("DELETE")

	// Start the HTTP server
	log.Println("Listening on port 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal(err)
	}
}
