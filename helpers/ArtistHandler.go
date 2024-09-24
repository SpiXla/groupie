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

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}
	id := pathParts[2]
	if r.URL.Path != "/artist/"+id {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	// rje3 aweld
	ad, er := strconv.Atoi(id)
	if er != nil {
		http.Error(w, "5Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	err := FetchData("locations", "/"+id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = FetchData("dates", "/"+id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = FetchData("artists", "/"+id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = FetchData("relation", "/"+id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// go FetchData("dates", id)
	// go FetchData("relation", id)

	Mok.Id = Artists[ad-1].Id
	Mok.Image = Artists[ad-1].Image
	Mok.Name = Artists[ad-1].Name
	Mok.CreationDate = Artists[ad-1].CreationDate
	Mok.FirstAlbum = Artists[ad-1].FirstAlbum
	Mok.Locations = Local.Local
	Mok.ConcertDates = Date.Dates
	Mok.Relations = Rela.DatesLocations
	// for d, t := range Rela.DatesLocations {
	// 	Rela.DatesLocations[d] = t
	// }

	file, err := template.ParseFiles("html/artist.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(Mok)
	err = file.Execute(w, Mok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
