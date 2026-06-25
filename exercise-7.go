package main

import (
	"fmt"
	"net/http"
)

func SimpleRedirector(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/legacy" {
		http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
		return
	}
}

func v2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")

}
