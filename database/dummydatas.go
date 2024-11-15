package database

import "github.com/HenkCode/go-CRUD/app/models"

func Dummy() *models.Movie {
	return &models.Movie{
		ID:    "1",
		Isbn:  "99999",
		Title: "movie one",
		Director: &models.Director{
			Firstname: "henk",
			Lastname:  "radhitya",
		},
	}
}