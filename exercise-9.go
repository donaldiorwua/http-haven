package main

import (
	"fmt"
	"io"
	"net/http"
)

func EchoChamber(w http.ResponseWriter, r *http.Request) {
	
	if r.Method != "POST" {
		http.Error(w, "not allowed method", http.StatusMethodNotAllowed)
	}
	content, err := io.ReadAll(r.Body)
	if err != nil && len(content) == 0 {
		defer r.Body.Close()
		http.Error(w, "body cannot be empty", http.StatusBadRequest)
	}
	fmt.Fprint(w, string(content))
}
