package helpers

type Artist []struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
}


type Card struct {
	Art struct {
		Id           int    `json:"id"`
		Image        string `json:"image"`
		Name         string `json:"name"`
		CreationDate int    `json:"creationDate"`
		FirstAlbum   string `json:"firstAlbum"`
	}
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}
type Location struct {
	Locations []string `json:"locations"`
}

type ConcertDate struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int               `json:"id"`
	DatesLocations map[string]string `json:"datesLocations"`
}

var Artists Artist
var Local Location
var Date ConcertDate
var Rela Relation

// type Relation struct {
// 	ID             int `json:"id"`
// 	DatesLocations struct {
// 		Location Dates
// 	} `json:"datesLocations"`
// }

// type Dates struct {
// 	Id    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }
