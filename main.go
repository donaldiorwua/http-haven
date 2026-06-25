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
	http.HandleFunc("/status", StatusCodeFactory)

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/v1/greet", APISubtreeV2)
	apiMux.HandleFunc("/v1/ping", APISubtreeV1)
	mainMux := http.NewServeMux()
	mainMux.Handle("/api/", http.StripPrefix("/api/", apiMux))
	http.ListenAndServe(":8080", mainMux)

	//http.ListenAndServe(":8080", nil)
}
