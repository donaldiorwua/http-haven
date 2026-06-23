package main

import (
	"fmt"
	"io"
	"net/http"
)

func TextCounter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprint(w, "Send a POST request with text to count words")
	}
	if r.Method == "POST" {
		char, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Fprint(w, len(string(char)))
	}
}

