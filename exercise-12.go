package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func StatusCodeFactory(w http.ResponseWriter, r *http.Request) {
	statuscode := r.URL.Query().Get("code")
	if statuscode == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
	}

	codestr, err := strconv.Atoi(statuscode)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
	}

	if codestr < 100 || codestr > 599 {
		http.Error(w, "code must be a valid HTTP status code (100-599)", http.StatusBadRequest)
	}
	//codename := http.StatusText(codestr)
	 w.WriteHeader(codestr)
	 fmt.Fprintf(w,  "Responding with status [%v]", codestr)

}
