package controllers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data any, status int) error {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	w.WriteHeader(status)
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(" Error executingtemplate:", err)
		return err
	}
	return nil
}
