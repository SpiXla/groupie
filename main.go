package main

import (
	"fmt"
	"net/http"

	"helpers/helpers"
)

func init() {
	if err := helpers.FetchInitialData(); err != nil {
		fmt.Printf("Error during initialization: %v\n", err)
	}
}

func main() {
	if len(helpers.Artists) == 0 {
		if err := helpers.FetchInitialData(); err != nil {
			fmt.Printf("Error during initialization: %v\n", err)
		}
	}
	http.HandleFunc("/", helpers.HomeHandler)
	http.HandleFunc("/artist/", helpers.ArtistHandler)
	fmt.Println("Listening on : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
