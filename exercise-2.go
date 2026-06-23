package main

import (
	"fmt"
	"net/http"
)

func PathValidation(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello, %v!", name)
	if name == "" {
		fmt.Fprint(w, "Hello, Guest!")
	}
	
}
