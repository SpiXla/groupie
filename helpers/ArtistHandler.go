package helpers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// ArtistHandler handles requests for artist profiles
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 || pathParts[2] == "" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	id := pathParts[2]
	ad, err := strconv.Atoi(id)
	if err != nil || ad < 1 || ad > len(Artists) {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	if err := fetchArtistData(ad); err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	file, err := template.ParseFiles("html/artist.html")
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	if err := file.Execute(w, Cards); err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}

func fetchArtistData(ad int) error {
	fetchPatterns := []string{"locations", "dates", "artists", "relation"}
	for _, pattern := range fetchPatterns {
		if err := FetchData(pattern, "/"+strconv.Itoa(ad)); err != nil {
			return err
		}
	}

	// Populate Cards with the artist's data
	Cards.Id = Artists[ad-1].Id
	Cards.Image = Artists[ad-1].Image
	Cards.Name = Artists[ad-1].Name
	Cards.Members = Artists[ad-1].Members
	Cards.CreationDate = Artists[ad-1].CreationDate
	Cards.FirstAlbum = Artists[ad-1].FirstAlbum
	Cards.Locations = Local.Local
	Cards.ConcertDates = Date.Dates
	Cards.Relations = Rela.DatesLocations

	return nil
}
