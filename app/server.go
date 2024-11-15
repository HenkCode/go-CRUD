package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HenkCode/go-CRUD/app/handlers"
	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter()
	port := ":8080"

	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}