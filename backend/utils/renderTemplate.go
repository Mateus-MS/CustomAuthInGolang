package utils

import (
	"html/template"
	"net/http"
)

func RenderTemplate[T any](w http.ResponseWriter, data T, paths ...string) error {
	template, err := template.ParseFiles(paths...)
	if err != nil {
		return err
	}

	template.Execute(w, data)
	return nil
}
