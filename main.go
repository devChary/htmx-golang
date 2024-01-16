package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Film struct {
	Title       string
	Director    string
	Description string
	ReleaseDate string
}

func main() {
	fmt.Println("Ello Mate!")

	handler := func(w http.ResponseWriter, r *http.Request) {
		template := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "Wolf of Wall Street", Director: "Martin Scorcese", Description: "Based on the true story of Jordan Belfort, from his rise to a wealthy stock-broker living the high life to his fall involving crime, corruption and the federal government.", ReleaseDate: "Mar 16, 2022"},
				{Title: "Inglorious Basterds", Director: "Quentin Tarantino", Description: "In Nazi-occupied France during World War II, a plan to assassinate Nazi leaders by a group of Jewish U.S. soldiers coincides with a theatre owner's vengeful plans for the same.", ReleaseDate: "Jan 2, 2017"},
			},
		}

		template.Execute(w, films)
	}
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
