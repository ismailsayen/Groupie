package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
