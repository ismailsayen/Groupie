package controllers

import (
	"fmt"
	"net/http"

	"groupieTracker/data"
)

func ArtistDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")
		var artists data.Artists
		var locations data.Locations
		var dates data.Dates
		var relations data.Relation

		FetchApi(data.ArtistsUrl+"/"+id, &artists)
		FetchApi(artists.Locations, &locations)
		FetchApi(locations.Dates, &dates)
		FetchApi(artists.Relations, &relations)

		fmt.Fprint(w, artists)
	}
}
