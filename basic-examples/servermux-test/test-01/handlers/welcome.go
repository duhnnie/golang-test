package handlers

import "net/http"

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to this API"))
}
