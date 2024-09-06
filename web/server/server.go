package main

import (
	"fmt"
	"net/http"
)

func WebServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
