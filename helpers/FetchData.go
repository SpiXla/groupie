package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchData(id string) error {
	pattern := "artists"
	link := "https://groupietrackers.herokuapp.com/api"

	if id != "" {
		pattern += "/" + id
	}
	resp, err := http.Get(link + "/" + pattern)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error fetching data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Log the response body for debugging
	fmt.Println("Response Body:", string(body))

	// Check if the response is an array or an object
	if body[0] == '[' {
		// Unmarshal into a slice of Artist
		var artists []Artist
		err = json.Unmarshal(body, &artists)
		if err != nil {
			return err
		}
		Artists = artists // Assign the fetched artists
	} else {
		// Unmarshal into a single Artist
		var artist Artist
		err = json.Unmarshal(body, &artist)
		if err != nil {
			return err
		}
		Artists = []Artist{artist} // Wrap it in a slice
	}

	return nil
}
