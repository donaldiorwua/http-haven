package main

import "net/http"

func main() {
	http.HandleFunc("/ping", PingPong)
	http.HandleFunc("/hello", PathValidation)
	http.HandleFunc("/count", TextCounter)
	http.HandleFunc("/calculate", BasicMath)
	http.HandleFunc("/agent", UserAgent)
	http.HandleFunc("/dashboard", SecureDashboard)
	http.HandleFunc("/legacy", SimpleRedirector)
	http.ListenAndServe(":8080", nil)
}
