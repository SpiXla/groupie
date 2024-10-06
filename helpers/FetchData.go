package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchData retrieves data from the API based on the given pattern and id.
func FetchData(pattern string, id string) error {
	link := "http://groupietrackers.herokuapp.com/api/"
	if pattern != "" {
		link += pattern
	}
	if id != "" {
		link += id
	}

	resp, err := http.Get(link)
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

	switch pattern {
	case "artists":
		if id == "" {
			return json.Unmarshal(body, &Artists)
		}
		// Handle individual artist fetching here if needed
	case "locations":
		return json.Unmarshal(body, &Local)
	case "dates":
		return json.Unmarshal(body, &Date)
	case "relation":
		return json.Unmarshal(body, &Rela)
	default:
		return err
	}

	return nil
}

func FetchInitialData() error {
	fetchPatterns := []string{"artists", "locations", "dates", "relation"}
	for _, pattern := range fetchPatterns {
		if err := FetchData(pattern, ""); err != nil {
			return fmt.Errorf("error fetching %s: %w", pattern, err)
		}
	}
	return nil
}