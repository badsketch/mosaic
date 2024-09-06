package main

import (
	"fmt"
	"net/http"
)

const MAX_MEMORY = 1e7

func WebServer(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(MAX_MEMORY)
	fmt.Fprint(w, "20")
}
