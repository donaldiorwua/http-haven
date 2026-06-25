package main

import (
	"fmt"
	"net/http"
)

func APISubtreeV1(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "pong")
}


func APISubtreeV2(w http.ResponseWriter, r *http.Request)  {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Greetings, Stranger!") 
	}
	fmt.Fprintf(w, "Greetings, [%v]!", name)
}