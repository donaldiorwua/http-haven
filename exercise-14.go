package main

import (
	"html/template"
	"net/http"
)

func TemplateRenderer(w http.ResponseWriter, r *http.Request) {
	const tempStr = `
	</DOCTYPE html>
	<html>
	<head><title>{{.Title}}</title></head>
	<body>
	<h1>{{.Title}}</h1>
	<p>{{.Body}}</p>
	</body>
	</html>`

	tpl := template.Must(template.New("page").Parse(tempStr))

	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")

	type PageData struct {
		Title string
		Body  string
	}
	page := PageData{
		Title: title,
		Body: body,
	}

	if page.Title == "" || page.Body == "" {
		 http.Error(w, "title and body are required", http.StatusBadRequest)
	}

	err := tpl.Execute(w, PageData{Title: title, Body: body})
	if err != nil {
		http.Error(w,  "template execution failed", http.StatusInternalServerError)
	}
}


