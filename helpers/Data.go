package helpers

type Artist []struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Card struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}
type Location struct {
	Local []string `json:"locations"`
}

type ConcertDate struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

var (
	Cards   Card
	Artists Artist
	Local   Location
	Date    ConcertDate
	Rela    Relation
)
