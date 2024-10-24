package main

import (
	"net/http"

	"groupieTracker/controllers"
)

func main() {
	http.HandleFunc("/", controllers.HomePage)
	http.HandleFunc("/detail", controllers.ArtistDetails)
	http.ListenAndServe(":8080", nil)
}
