package models

type Movie struct {
	ID        string   `json:"id"`
	Isbn      string   `json:"isbn"`
	Title 	  string   `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Movies = []Movie{
	{ID:    "1",
	Isbn:  "99999",
	Title: "movie one",
	Director: &Director{
		Firstname: "henk",
		Lastname:  "radhitya",
	}},
	{ID:    "2",
	Isbn:  "88888",
	Title: "movie two",
	Director: &Director{
		Firstname: "edo",
		Lastname:  "kurniawan",
	}},
}