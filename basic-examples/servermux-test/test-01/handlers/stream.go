package handlers

import (
	"fmt"
	"net/http"
	"time"
)

var counter = 0

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	rc := http.NewResponseController(w)
	clientDisconnected := r.Context().Done()

	for {
		select {
		case <-clientDisconnected:
			fmt.Println("Client disconnected")
			return
		case <-t.C:
			// counter++
			fmt.Fprintf(w, "data: Hola #%d\n\n", counter)
			rc.Flush()
		}
	}
}

func increment() {
	t := time.NewTicker(time.Millisecond * 1450)

	for {
		select {
		case <-t.C:
			counter++
		}
	}
}

func init() {
	go increment()
}
