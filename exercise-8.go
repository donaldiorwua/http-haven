package main

import (
	"fmt"
	"net/http"
)

func MethodInspector(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "You made a GET request.")
	case "POST":
		fmt.Fprint(w, "You made a POST request.")
	default:
		fmt.Fprintf(w, "You made a [%v] request.", r.Method)
	}
}
