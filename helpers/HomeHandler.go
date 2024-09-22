package helpers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Notr Allowed 405", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Not Found 404", http.StatusNotFound)
		return
	}

	err := FetchData("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, err := template.ParseFiles("html/index.html")
	if err != nil {
		http.Error(w, "2Internal Server Error 500", http.StatusInternalServerError)
		return
	}

	err = file.Execute(w, Artists)
	if err != nil {
		http.Error(w, "3Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}
