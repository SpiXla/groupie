package main

type Artist struct {
	ID           int           `json:"id"`
	Image        string        `json:"image"`
	Name         string        `json:"name"`
}

type Location struct {
	ID       int      `json:"id"`
	Location []string `json:"location"`
	Dates    Dates
}

type ConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int `json:"id"`
	DatesLocations struct {
		Location Dates
	} `json:"datesLocations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

var Artists []Artist
