package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HenkCode/go-CRUD/app/models"
	"github.com/gorilla/mux"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Movies)
}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range models.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}


func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range models.Movies {
		if item.ID == params["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Movie is deleted"})
			break
		}
	}
	json.NewEncoder(w).Encode(models.Movies)
	http.Error(w, "Movie not found", http.StatusNotFound)
	
}