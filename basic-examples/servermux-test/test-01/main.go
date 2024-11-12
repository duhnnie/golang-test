package main

import (
	"example/servermux-test-01/handlers"
	"fmt"
	"net/http"
)

var MainCounter = 0

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler("http://google.com", 307)

	mux.Handle("/mierda", rh)
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/welcome", handlers.WelcomeHandler)
	mux.HandleFunc("/heavy", handlers.HeavyHandler)
	mux.HandleFunc("/stream", handlers.StreamHandler)

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
