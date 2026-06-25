package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func FormDecoder(w http.ResponseWriter, r *http.Request) {
	var temp *template.Template
	temp = template.Must(template.ParseFiles("index.html"))

	if r.Method != "POST" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	r.ParseForm()

	username := r.FormValue("username")
	language := r.FormValue("language")

	user := struct {
		Username string
		Language string
	}{
		Username: username,
		Language: language,
	}

	if user.Username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
	}
	if user.Language == "" {
		http.Error(w, "language is required", http.StatusBadRequest)
	}

	result, _ := fmt.Fprintf(w, "Hello %v, you are coding in %v!", user.Username, user.Language)

	err := temp.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
