package main

import (
	"fmt"
	"net/http"
)

func Headers(w http.ResponseWriter, r *http.Request) {

	content := r.Header.Get("X-Custom-Token")
	content_type := r.Header.Get("Content-type")

	if content == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
	}
	fmt.Fprintf(w, "Token received: %v", content)
	if content_type == "" {
		fmt.Fprint(w, "Content-Type not provided")
	}
	fmt.Fprintf(w, "Content-Type: %v", content_type)

}
