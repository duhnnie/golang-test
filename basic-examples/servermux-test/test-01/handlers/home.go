package handlers

import (
	"example/servermux-test-01/static"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(static.GetIndexPage())
}
