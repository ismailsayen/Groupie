package controllers

import (
	"net/http"

	"groupieTracker/data"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == http.MethodGet {
			var artists []data.Artists
			FetchApi(data.ArtistsUrl, &artists)
			RenderTemplate(w, "./templates/index.html", artists, 200)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		http.Error(w, "Page note Found", http.StatusNotFound)
	}
}
