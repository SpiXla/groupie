package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", Homehandler)
	fmt.Println("Listening on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	err := FetchData("")
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	file, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	err = file.Execute(w, Artists)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
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
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &Artists)
	if err != nil {
		return err
	}

	return nil
}
