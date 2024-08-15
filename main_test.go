package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker/handlers"
)

func TestAboutHandler(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/about", handlers.AboutHandler)

	writer := httptest.NewRecorder()
	reaquest, _ := http.NewRequest("GET", "/about", nil)

	mux.ServeHTTP(writer, reaquest)

	if writer.Code != 200 {
		t.Errorf("Status Code is %v instead of 200", writer.Code)
		return
	}

	t.Log(writer.Code)
}
