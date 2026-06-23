package main

import (
	"fmt"
	"net/http"
)

func SimpleRedirector(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path == "/legacy" {
		http.Redirect(w, r, "http://8080/v2", http.StatusMovedPermanently)
		fmt.Fprint(w, "Welcome to version 2")
	}
}