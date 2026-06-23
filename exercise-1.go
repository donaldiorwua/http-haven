package main

import (
	"fmt"
	"net/http"
)

func PingPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
