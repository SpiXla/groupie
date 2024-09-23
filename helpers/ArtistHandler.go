package helpers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	// rje3 aweld
	ad, er := strconv.Atoi(id)
	if er != nil {
		http.Error(w, "5Internal Server Error 500", http.StatusInternalServerError)
		return
	}

err :=FetchData("locations", "/"+id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// go FetchData("dates", id)
	// go FetchData("relation", id)
	var Mok Card
	Mok.Art = Artists[ad-1]
	Mok.Locations = Local.Locations
	Mok.ConcertDates = Date.Dates
	for d,t := range Rela.DatesLocations {
		Rela.DatesLocations[d] = t 
	}

	fmt.Println(Local)


	file, err := template.ParseFiles("html/artist.html")
	if err != nil {
		http.Error(w, "5Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	 
	err = file.Execute(w, Mok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
