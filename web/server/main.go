package main

import (
	"log"
	"net/http"
)

// Inside ~/mosaic/web/server, run with `go build` then `./server`
func main() {
	handler := http.HandlerFunc(WebServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
