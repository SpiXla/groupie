package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
		return fmt.Errorf("error fetching data: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if pattern == "artists" {
		// var artists Artist
		err = json.Unmarshal(body, &Artists)
		if err != nil {
			return err
		}
	} else if pattern == "locations" {
		err = json.Unmarshal(body, &Local)
		if err != nil {
			return err
		}
	} else if pattern == "dates" {
		var date []ConcertDate
		err = json.Unmarshal(body, &date)
		if err != nil {
			return err
		}
	} else if pattern == "relation" {
		var relation []Relation
		err = json.Unmarshal(body, &relation)
		if err != nil {
			return err
		}
	}

	// fmt.Println(string(body))

	return nil
}
