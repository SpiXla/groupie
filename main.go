package main

import (
	"fmt"
	"net/http"

	"helpers/helpers"
)

func main() {
	http.HandleFunc("/", helpers.HomeHandler)
	http.HandleFunc("/artist/", helpers.ArtistHandler)
	fmt.Println("Listening on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
