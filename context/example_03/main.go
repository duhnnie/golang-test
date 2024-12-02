package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func home(w http.ResponseWriter, r *http.Request) {
	select {
	case <-r.Context().Done():
		fmt.Println("Context Done!")
	case <-time.After(time.Second * 5):
		w.Write([]byte("Everything is ok!"))
	}
}

func main() {
	r := chi.NewRouter()

	r.Get("/", home)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Server running at port")
	}
}
