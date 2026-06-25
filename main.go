package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/ping", PingPong)
	http.HandleFunc("/hello", PathValidation)
	http.HandleFunc("/count", TextCounter)
	http.HandleFunc("/calculate", BasicMath)
	http.HandleFunc("/agent", UserAgent)
	http.HandleFunc("/dashboard", SecureDashboard)
	http.HandleFunc("/legacy", SimpleRedirector)
	http.HandleFunc("/v2", v2)
	http.HandleFunc("/method-inspector", MethodInspector)
	http.HandleFunc("/echo", EchoChamber)
	http.HandleFunc("/headers", Headers)
	http.HandleFunc("/form", FormDecoder)
	http.ListenAndServe(":8080", nil)
}
