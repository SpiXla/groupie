package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	// if r.URL.Path != "/artist/" {
	// 	http.Error(w, "Not Found 404", http.StatusNotFound)
	// 	return
	// }

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}
	id := pathParts[2]

	err := FetchData(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ar, _ := json.Marshal(Artists)
	var art Artist
	// ar, _ := json.Marshal(art.ConcertDates)
	// fmt.Println(string(ar))
	locations := art.Locations
	concertdates := art.ConcertDates
	relations := art.Relations
	fmt.Println(locations, concertdates, relations)

	

	file, err := template.ParseFiles("html/artist.html")
	if err != nil {
		http.Error(w, "5Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	err = file.Execute(w, Artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
