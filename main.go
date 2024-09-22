package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Artist struct {
	ID           int           `json:"id"`
	Image        string        `json:"image"`
	Name         string        `json:"name"`
	Members      []string      `json:"members"`
	CreationDate int           `json:"creationDate"`
	FirstAlbum   string        `json:"firstAlbum"`
	Locations    []Location    `json:"locations"`    // Use Location struct
	ConcertDates []ConcertDate `json:"concertDates"` // Use ConcertDate struct
	Relations    []Relation    `json:"relations"`    // Use Relation struct
}

type Location struct {
	ID       int      `json:"id"`
	Location string `json:"location"`
	Dates    Dates
}

type ConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID       int   `json:"id"`
	Location Dates `json:"relatedArtists"`
}
type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

var artists []Artist

func main() {
	http.HandleFunc("/", Homehandler)
	fmt.Println("Listening on : localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err := FetchData("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = file.Execute(w, artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func FetchData(pattern string) error {
	if pattern == "" {
		pattern = "artists"
	}
	link := "https://groupietrackers.herokuapp.com/api"
	resp, err := http.Get(link + "/" + pattern)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, &artists)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return nil
}
