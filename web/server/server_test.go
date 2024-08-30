package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETIndex(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	WebServer(response, request)

	t.Run("returns the index", func(t *testing.T) {
		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
