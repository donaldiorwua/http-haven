package main

import (
	"fmt"
	"net/http"
)

func UserAgent(w http.ResponseWriter, r *http.Request)  {
	user := r.Header.Get("User-Agent")
	if user == "" {
	http.Error(w, "not found", http.StatusNotFound)
	}

	fmt.Fprintf(w, "You are visiting us using:[%v]", user)
}