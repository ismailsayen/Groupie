package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

type Artists struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type IndexLocation struct {
	Index []Locations `json:"index"`
}

const (
	artistsUrl   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsUrl = "https://groupietrackers.herokuapp.com/api/locations"
	datesUrl     = "https://groupietrackers.herokuapp.com/api/dates"
	relation     = "https://groupietrackers.herokuapp.com/api/relation"
)

func main() {
	var artists []Artists
	var location IndexLocation
	FetchApi(artistsUrl, &artists)
	FetchApi(locationsUrl, &location)
	for _, artist := range artists {
		idLoc, _ := strconv.Atoi(path.Base(artist.Locations))
		if artist.Id == 12 {
			fmt.Printf("ID	: %v\n", artist.Id)
			fmt.Printf("Name	: %v\n", artist.Name)
			fmt.Println("Members	:")
			for _, mbr := range artist.Members {
				fmt.Println("   -", mbr)
			}
			fmt.Println("Locations	:")
			for _, loc := range location.Index {
				if loc.Id == idLoc {
					for _, ele := range loc.Locations {
						fmt.Println("  -", ele)
					}
				}
			}
			fmt.Printf("CreationDate: %v\n", artist.CreationDate)
			fmt.Printf("FirstAlbum: %v\n", artist.FirstAlbum)
		}
	}
}

func FetchApi(Url string, data any) {
	resp, err := http.Get(Url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
