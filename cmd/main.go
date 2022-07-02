package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/walid-abubakar/books-api/pkg/handlers"
)


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost)

	log.Println("API is running")
	http.ListenAndServe(":4000", router)
}
