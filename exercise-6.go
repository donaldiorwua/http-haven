package main

import (
	"fmt"
	"net/http"
)

func SecureDashboard (w http.ResponseWriter, r *http.Request)  {
	apikey := r.Header.Get("X-API-Key")
	if apikey != "secret123" {
		http.Error(w, "Acess denied", http.StatusUnauthorized)
	}
	if apikey == "secret123"{
	fmt.Fprint(w, "Welcome Acess granted")
}	}