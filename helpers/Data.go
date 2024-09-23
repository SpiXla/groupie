package helpers

type Artist struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
	Locations    string `json:"locations"`
	ConcertDates string `json:"concertDates"`
	Relations    string `json:"relations"`
}
type ArtistResponse struct {
	Artists []Artist `json:"artists"`
}

// type Location struct {
// 	ID       int      `json:"id"`
// 	Location []string `json:"location"`
// 	Dates    Dates
// }

// type ConcertDate struct {
// 	ID    int      `json:"id"`
// 	Dates []string `json:"dates"`
// }

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

var Artists []Artist
