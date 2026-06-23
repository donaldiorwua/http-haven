package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func BasicMath(w http.ResponseWriter, r *http.Request) {
	operation := r.URL.Query().Get("op")
	numstr1 := r.URL.Query().Get("a")
	numstr2 := r.URL.Query().Get("b")

	num1, err := strconv.Atoi(numstr1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	num2, err := strconv.Atoi(numstr2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var result int
	switch operation {
	case "add":
		result = num1 + num2
		fmt.Fprint(w, result)
		return
	case "multiply":
		result = num1 * num2
		fmt.Fprint(w, result)
		return
	case "divide":
		result = num1 / num2
		fmt.Fprint(w, result)
		return
	case "subtract":
		result = num1 - num2
		fmt.Fprint(w, result)
		return
	default:
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
