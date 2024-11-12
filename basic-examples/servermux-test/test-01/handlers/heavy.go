package handlers

import (
	"fmt"
	"net/http"
)

func HeavyHandler(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5_000_000; i++ {
		fmt.Printf("counter is %d\n", i)
	}

	w.Write([]byte("heavy process"))
}
